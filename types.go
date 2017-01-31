package charta

import "time"

//InternalCV internal structure for namespaces
type InternalCV struct {
	Name string         `json:"name"`
	Pods []*InternalPod `json:"pods"`
}

//InternalPod internal structure for Pods and configuration given rc
type InternalPod struct {
	Name         string            `json:"name"`
	Env          map[string]string `json:"vars"`
	CreationTS   time.Time
	ReplicaCount *int32 `json:"replicacount"`
	Active       bool   `json:"active"`
}

//DTOs

//NamespeceDTO DTO for namespaces list
type NamespeceDTO struct {
	Data []string `json:"data"`
}

//Pod DTO
type PodDTO struct {
	Name string            `json:"name"`
	Env  map[string]string `json:"vars"`
}

//Pod List DTO
type PodListDTO struct {
	Data []PodDTO `json:"data"`
}
