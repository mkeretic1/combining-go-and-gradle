package k8s

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNamespaces(clientset kubernetes.Interface) ([]string, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaceList []string
	for _, namespace := range namespaces.Items {
		namespaceList = append(namespaceList, namespace.Name)
	}

	if namespaceList == nil || len(namespaceList) == 0 {
		return []string{}, nil
	}
	return namespaceList, nil
}
