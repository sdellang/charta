package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdellang/charta"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
	port := flag.String("port", "8080", "HTTP port")

	flag.Parse()

	fmt.Printf("Configuring kubernetes client...\n KubeConfig: %s \n", *kubeconfig)

	config, err := buildConfig(*kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)

	icv, err := charta.BuildClusterView(clientset)

	charta.SetKubeClient(clientset)
	charta.SetInternalCV(icv)

	if err != nil {
		panic(err.Error())
	}

	//router configuration
	router := mux.NewRouter()

	router.HandleFunc("/api/cluster", charta.GetClusterView).Methods("GET")
	router.HandleFunc("/api/namespaces", charta.GetNamespaces).Methods("GET")
	router.HandleFunc("/api/namespaces/{name}", charta.GetNamespaceView).Methods("GET")
	//serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/dist/")))
	fmt.Printf("Starting...")

	log.Fatal(http.ListenAndServe(":"+*port, router))

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
