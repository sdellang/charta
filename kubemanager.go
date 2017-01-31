package charta

import (
	"fmt"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

//LoadNameSpace load namespaces from Kubernetes to build initial cluster view
func LoadNameSpace(clientset *kubernetes.Clientset) (*[]InternalCV, error) {

	nsi, err := clientset.Core().Namespaces().List(v1.ListOptions{})

	nsl := make([]InternalCV, len(nsi.Items))

	for i := 0; i < len(nsi.Items); i++ {
		nsl[i] = InternalCV{Name: nsi.Items[i].GetName()}
	}
	if err != nil {
		return nil, err
	}

	return &nsl, nil
}

//LoadPodConfig load POD configuration to build initial cluster view
func LoadPodConfig(clientset *kubernetes.Clientset, internalNs *[]InternalCV) (*[]InternalCV, error) {

	for i := 0; i < len(*internalNs); i++ {
		dc, err := clientset.Core().ReplicationControllers((*internalNs)[i].Name).List(v1.ListOptions{})

		if err != nil {
			return nil, err
		}

		if dc.Items != nil && len(dc.Items) != 0 {

			(*internalNs)[i].Pods = make([]*InternalPod, len(dc.Items))

			for j := 0; j < len(dc.Items); j++ {
				(*internalNs)[i].Pods[j] = convertPod(dc.Items[j])
			}
		}

		selectActive(&(*internalNs)[i])
	}

	return internalNs, nil
}

//BuildClusterView build initial cluster view
func BuildClusterView(clientset *kubernetes.Clientset) (*[]InternalCV, error) {

	fmt.Printf("Building Internal Cluster View")
	tm, err := LoadNameSpace(clientset)
	if err != nil {
		return nil, err
	}

	tm, err = LoadPodConfig(clientset, tm)
	if err != nil {
		return nil, err
	}

	return tm, nil
}

func convertPod(rc v1.ReplicationController) *InternalPod {
	ipm := make(map[string]string)
	var at bool

	co := rc.Spec.Template.Spec.Containers[len(rc.Spec.Template.Spec.Containers)-1]
	env := co.Env

	for j := 0; j < len(env); j++ {
		ipm[env[j].Name] = env[j].Value
	}

	at = (*rc.Spec.Replicas > 0)

	return &InternalPod{Name: rc.Name, Env: ipm, CreationTS: rc.CreationTimestamp.Time, ReplicaCount: rc.Spec.Replicas, Active: at}
}

func selectActive(in *InternalCV) {
	//select the active rc configurations based on last config and active replicas count
	fmt.Printf("Selecting Active ReplicationController configurations")
	ae := make(map[string]*InternalPod)

	for _, el := range in.Pods {

		nn := el.Name[:strings.Index(el.Name, "-")]
		ap, ex := ae[nn]

		if ex == true {

			if el.CreationTS.After(ap.CreationTS) {
				ae[nn] = el
			}

		} else {
			ae[nn] = el
		}
	}

	for _, value := range ae {
		value.Active = true
	}

}
