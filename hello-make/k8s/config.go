package k8s

import (
	"encoding/base64"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetClientSet() (*kubernetes.Clientset, error) {
	ca, caErr := decodeString("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJlRENDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdGMyVnkKZG1WeUxXTmhRREUyTmpjNE1UQTJPRFV3SGhjTk1qSXhNVEEzTURnME5EUTFXaGNOTXpJeE1UQTBNRGcwTkRRMQpXakFqTVNFd0h3WURWUVFEREJock0zTXRjMlZ5ZG1WeUxXTmhRREUyTmpjNE1UQTJPRFV3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFSOTJzTHRMSEJObzAxQzJKZjkxNUNnQjdhSkl5R3dYNFllSTJ4YWF5YUgKVXVqc21sbGg1aFZVUjZzaXI0cnNPM3BCbHAyb3h4TTFZRW9ad3QvYTRPQ0FvMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVXE5RDdYV3VCeGhTcUJvK1RaSVB2Cm9EeVlHWnd3Q2dZSUtvWkl6ajBFQXdJRFNRQXdSZ0loQUpoMk40ck9kUkdpUURmZTgxMVJRcVc5N1ZIU2pzK2gKS3VEaFNXUWRGZVplQWlFQTVLdHFtQU9wZFg3UW4wZFBuUXYrUWJVR1BpQTIwMnRVV0luWFNXZjRMbW89Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K")
	clientCrt, clientCrtErr := decodeString("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJrVENDQVRlZ0F3SUJBZ0lJUDdDYnp2OTFXcjB3Q2dZSUtvWkl6ajBFQXdJd0l6RWhNQjhHQTFVRUF3d1kKYXpOekxXTnNhV1Z1ZEMxallVQXhOalkzT0RFd05qZzFNQjRYRFRJeU1URXdOekE0TkRRME5Wb1hEVEl6TVRFdwpOekE0TkRRME5Wb3dNREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhGVEFUQmdOVkJBTVRESE41CmMzUmxiVHBoWkcxcGJqQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlBd0VIQTBJQUJOSmRLOWMzZGY2RWZCbmsKL2M5WENEMHNSZ2Z4UnFTUTNLRjZGN0cxZ0MrVm9seXoxV1drcEhnSEE0WGg5ckZadGsxNjYrR2NOUDUxVjRmYgphS0dGOGcralNEQkdNQTRHQTFVZER3RUIvd1FFQXdJRm9EQVRCZ05WSFNVRUREQUtCZ2dyQmdFRkJRY0RBakFmCkJnTlZIU01FR0RBV2dCUmt5SjdZWlJNWlFsTDdZRFJrNVJKWXdCWFhpREFLQmdncWhrak9QUVFEQWdOSUFEQkYKQWlCNUkwWFhmdUR2RTkrMG95RHJDSFIvMlAreS9RUzlwZFRpVlU3Zm5sR3lzd0loQUtsTW05VHhlb1RwTG5XSAoxNWtmb0I0UTBDVll0MEZmTlhOd05NKzBsYTJ1Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJkekNDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdFkyeHAKWlc1MExXTmhRREUyTmpjNE1UQTJPRFV3SGhjTk1qSXhNVEEzTURnME5EUTFXaGNOTXpJeE1UQTBNRGcwTkRRMQpXakFqTVNFd0h3WURWUVFEREJock0zTXRZMnhwWlc1MExXTmhRREUyTmpjNE1UQTJPRFV3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFSTzVsdlFPWm1id1gwK0J2ZmlBVFNzcVVDUFdTV2NWclRsdllxWVFzeG8KUXl0Y0tpbXAzaEU0cGNrQ3MvWXVBemp4R3RDbkZBbzladU91dFl6T2k1ek9vMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVVpNaWUyR1VUR1VKUysyQTBaT1VTCldNQVYxNGd3Q2dZSUtvWkl6ajBFQXdJRFNBQXdSUUlnSWxCTEc4TTJzUzdrSlZjbWlzR2E0RkFXYUIyUUFNclMKRVhLTVYvVUdQbmdDSVFEekdXZmRsOVQycnJhTWJqYmY1NlN6ZWRFMHNzajd4WEw3d0RCQXZuSFQ3dz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K")
	clientKey, clientKeyErr := decodeString("LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUVWazViVWdEcklVWWViUGREUFNKQ0tvVVU1OUErZGFtT2MxMldQRmRpTVFvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFMGwwcjF6ZDEvb1I4R2VUOXoxY0lQU3hHQi9GR3BKRGNvWG9Yc2JXQUw1V2lYTFBWWmFTawplQWNEaGVIMnNWbTJUWHJyNFp3MC9uVlhoOXRvb1lYeUR3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo=")

	if clientCrtErr != nil || clientKeyErr != nil || caErr != nil {
		fmt.Println(clientCrtErr, clientKeyErr, caErr)
	}

	config := &rest.Config{
		Host: "https://localhost:39205",
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
