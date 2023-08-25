package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	rclient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"127.0.0.1:6379"},
		DB:    0,
	})
	err := rclient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	// use the current context in kube config
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// creates the client set
	k8sclientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	versionInfo, err := k8sclientset.ServerVersion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("k8s server version: %v\n", versionInfo)

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hello",
			Labels: map[string]string{
				"app": "kulery",
			},
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "hello",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "hello",
							Image:   "hello-world:latest",
							Command: []string{},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
	}

	result, err := k8sclientset.BatchV1().Jobs(corev1.NamespaceDefault).Create(context.TODO(), &job, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	// k8sclientset.BatchV1().CronJobs(corev1.NamespaceDefault).Watch(context.TODO(), metav1.ListOptions{})
	// watcher, err := k8sclientset.BatchV1().Jobs(corev1.NamespaceDefault).Watch(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// // 处理 watch 事件
	// for event := range watcher.ResultChan() {
	// 	fmt.Println("===========================================")
	// 	fmt.Printf("Event type: %v\n", event.Type)
	// 	fmt.Printf("Pod name: %v\n", event.Object.(*batchv1.Job).Name)
	// 	fmt.Printf("Pod status succeeded: %v\n", event.Object.(*batchv1.Job).Status.Succeeded)
	// 	fmt.Printf("Pod status active: %v\n", event.Object.(*batchv1.Job).Status.Active)
	// }
}
