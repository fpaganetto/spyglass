package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type DiscoveredApi struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Discovery string `json:"discovery"`
}

var clientset *kubernetes.Clientset

func main() {
	localEnvironment := flag.Bool("local", false, "(optional) run from local environment")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	clientset = initK8Client(*localEnvironment, *kubeconfig)

	http.HandleFunc("/", discovery)
	http.ListenAndServe(":8090", nil)
}

func discovery(w http.ResponseWriter, req *http.Request) {
	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d ingress in the cluster\n", len(ingresses.Items))

	var apis []DiscoveredApi
	for _, ingress := range ingresses.Items {
		if name, exists := ingress.Annotations["spyglass/name"]; exists {
			api := DiscoveredApi{name, ingress.Spec.Rules[0].Host, ingress.Annotations["spyglass/discovery"]}
			fmt.Printf("%+v\n", api)
			apis = append(apis, api)
		}
	}

	js, err := json.Marshal(formatResponse(apis))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func initK8Client(localEnvironment bool, kubeConfigPath string) *kubernetes.Clientset {
	var config *rest.Config
	var err error

	if localEnvironment {
		fmt.Println("Starting using local kubeconfig configuration")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	} else {
		fmt.Println("Starting using cluster configuration")
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func formatResponse(apis []DiscoveredApi) interface{} {
	objectStyle := true //XXX from flag, global variable?
	if objectStyle {
		type DiscoveredApiObjectDto struct {
			Url       string `json:"url"`
			Discovery string `json:"discovery"`
		}
		toReturn := make(map[string]DiscoveredApiObjectDto)
		for _, api := range apis {
			toReturn[api.Name] = DiscoveredApiObjectDto{api.Url, api.Discovery}
		}
		return toReturn
	} else {
		return apis
	}
}
