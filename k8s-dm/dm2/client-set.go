package dm2

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func Dm2() {
	fmt.Println("1. Dm2 func.")
	// 1. 加载kube-config配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/sgr/.kube/config")
	if err != nil {
		log.Fatalf("get kube config error, %s\n", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("--", err.Error())
	}

	ctx := context.TODO()
	podList, err := clientset.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal("list pod : ", err.Error())
	}

	for _, item := range podList.Items {
		fmt.Printf("namespace: %s, pod name: %s\n", item.Namespace, item.Name)
	}

}
