package charta
import (
"k8s.io/client-go/pkg/api/v1"
)

type InternalCV struct {
	Environment string "json: environment"
	Name string "json: name"
	Pods []v1.Pod "json: pods"
}
