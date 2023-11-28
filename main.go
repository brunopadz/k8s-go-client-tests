package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
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
	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	fmt.Println(len(pods.Items))

	for _, pod := range pods.Items {
		fmt.Println(pod.Status)
		//if pod.Status == "Failed" {
		//	clientset.CoreV1().Pods("").Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		//}
	}
}
