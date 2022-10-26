package iface

type Backend interface {
	GetTask(taskUUID string) error
}