package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d ingress in the cluster\n", len(ingresses.Items))

	for ingress := range ingresses.Items {
		fmt.Println(ingresses.Items[ingress].Spec.Rules[0].Host)
	}
}
