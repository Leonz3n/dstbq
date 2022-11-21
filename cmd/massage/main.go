package main

import (
	"context"
	redisbackend "github.com/Leonz3n/k8s-job-massage/internal/backends/redis"
	redisbroker "github.com/Leonz3n/k8s-job-massage/internal/brokers/redis"
	"github.com/Leonz3n/k8s-job-massage/internal/server"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
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

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "desktop-docker"))
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	broker := redisbroker.NewBrokerGR(logger.Sugar(), rclient)
	backend := redisbackend.NewBackendGR(logger.Sugar(), rclient)

	svr := server.NewServer(&server.Config{
		Concurrency: 3,
		Logger:      logger.Sugar(),
		Broker:      broker,
		Backend:     backend,
		KClientSet:  clientset,
	})
	svr.Start()
}
