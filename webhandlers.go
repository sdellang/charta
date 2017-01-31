package charta

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
)

var clientset *kubernetes.Clientset

var internalCV *[]InternalCV

func GetInternalCV() *[]InternalCV {
	return internalCV
}

func SetInternalCV(icv *[]InternalCV) {
	internalCV = icv
}

func GetKubeClient() *kubernetes.Clientset {
	return clientset
}

func SetKubeClient(client *kubernetes.Clientset) {
	clientset = client
}

func GetClusterView(w http.ResponseWriter, req *http.Request) {

	js, err := json.Marshal(internalCV)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetNamespaces(w http.ResponseWriter, req *http.Request) {

	nsi := make([]string, len(*internalCV))
	for i := 0; i < len(*internalCV); i++ {
		nsi[i] = (*internalCV)[i].Name
	}

	ndto := NamespeceDTO{nsi}

	js, err := json.Marshal(ndto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetNamespaceView(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	var pl = PodListDTO{Data: nil}

	for i := 0; i < len(*internalCV); i++ {
		if (*internalCV)[i].Name == params["name"] {
			for _, pod := range (*internalCV)[i].Pods {
				if pod.Active == true {
					pd := PodDTO{Name: pod.Name, Env: pod.Env}
					pl.Data = append(pl.Data, pd)
				}
			}
			js, err := json.Marshal(pl)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}

}
