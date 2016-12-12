package charta

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

func LoadNameSpace(clientset *kubernetes.Clientset) (*[]InternalCV, error) {

	nsi, err := clientset.Core().Namespaces().List(v1.ListOptions{})


	nsl := make([]InternalCV, len(nsi.Items))

	for i := 0; i < len(nsi.Items) ; i++ {
		nsl[i] = InternalCV{Name: nsi.Items[i].GetName()}
	}
	if err != nil {
		return nil, err
	}

	return &nsl,nil
}

func LoadPodConfig(clientset *kubernetes.Clientset, internalNs *[]InternalCV) (*[]InternalCV, error) {

	for i := 0; i < len(*internalNs); i++ {
		dc, err := clientset.Core().Pods((*internalNs)[i].Name).List(v1.ListOptions{})

		if err != nil {
			return nil, err
		}
		if dc.Items != nil && len(dc.Items) != 0 {
			(*internalNs)[i].Pods = dc.Items
		}
	}

	return internalNs,nil
}

func BuildClusterView(clientset *kubernetes.Clientset) (*[]InternalCV, error) {

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