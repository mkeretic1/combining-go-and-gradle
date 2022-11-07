package main

import (
	"fmt"
	"hello-mage/k8s"
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
