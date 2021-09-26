package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DiscoveredApi struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Discovery string `json:"discovery"`
}

func discovery(w http.ResponseWriter, req *http.Request) {
	ingresses, err := app.k8client.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

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

	js, _ := json.Marshal(formatResponse(apis))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
