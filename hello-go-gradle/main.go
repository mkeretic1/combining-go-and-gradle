package main

import (
	"fmt"
	"github.com/mkeretic/hello-go-gradle/k8s"
)

func main() {
	clientSet, err := k8s.GetClientSet()
	if err != nil {
		k8s.HandleError(err)
	}

	namespaces, err := k8s.ListNamespaces(clientSet)
	if err != nil {
		k8s.HandleError(err)
	}

	fmt.Println("Namespaces: ", namespaces)
}
