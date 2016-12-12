package charta
import (
	"encoding/json"
	"net/http"
	"k8s.io/client-go/kubernetes"
	"github.com/gorilla/mux"
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

	nsi := make([]string,len(*internalCV))
	for i := 0; i < len(*internalCV) ; i++ {
		nsi[i] = (*internalCV)[i].Name
	}

	js, err := json.Marshal(nsi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetNamespaceView(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	for i := 0; i < len(*internalCV) ; i++ {
		if (*internalCV)[i].Name == params["name"] {
			js, err := json.Marshal((*internalCV)[i])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}

}

