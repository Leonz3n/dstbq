package task

type Container struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

type Signature struct {
	Namespace  string       `json:"namespace"` // the K8S namespace where the Job is in
	UUID       string       `json:"uuid"`
	Name       string       `json:"name"`
	Containers []*Container `json:"containers"`
	Retry      int          `json:"retry"`
}
