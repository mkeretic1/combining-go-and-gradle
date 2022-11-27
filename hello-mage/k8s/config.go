package k8s

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Cert files locations
var CaCertFile = "ca.pem"
var ClientCertFile = "client-cert.pem"
var ClientKeyFile = "client-key.pem"

// Cluster url
var Host = "https://172.17.0.1:32815"

func GetClientSet() (*kubernetes.Clientset, error) {
	ca, caErr := ioutil.ReadFile(CaCertFile)
	clientCrt, clientCrtErr := ioutil.ReadFile(ClientCertFile)
	clientKey, clientKeyErr := ioutil.ReadFile(ClientKeyFile)

	if clientCrtErr != nil || clientKeyErr != nil || caErr != nil {
		fmt.Println(clientCrtErr, clientKeyErr, caErr, ca)
	}

	config := &rest.Config{
		Host: Host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   ca,
			CertData: clientCrt,
			KeyData:  clientKey,
		},
	}

	return kubernetes.NewForConfig(config)
}

func decodeString(val string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(val)
}
