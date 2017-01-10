package charta

import "time"

type InternalCV struct {
	Name string         `json: "name"`
	Pods []*InternalPod `json: "pods"`
}

type InternalPod struct {
	Name         string            `json: "name"`
	Env          map[string]string `json: "vars"`
	CreationTS   time.Time
	ReplicaCount *int32 `json: "replicacount"`
	Active       bool   `json: "active"`
}
