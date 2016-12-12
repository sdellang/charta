package main
import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/sdellang/charta"
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

	icv, err :=charta.BuildClusterView(clientset)

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

	fmt.Printf("Starting...")

	log.Fatal(http.ListenAndServe(":" + *port, router))

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}