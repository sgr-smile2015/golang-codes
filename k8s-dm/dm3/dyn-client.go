package dm3

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func Dm3() {
	fmt.Println("1. Dm2 func.")
	// 1. 加载kube-config配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	checkError("get kube config", err)

	dynClient, err := dynamic.NewForConfig(config)
	checkError("dynamic new config", err)
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	ctx := context.TODO()

	unstructuredList, err := dynClient.Resource(gvr).Namespace("kube-system").List(ctx, metav1.ListOptions{})
	podList := corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredList.UnstructuredContent(), &podList)
	checkError("unstructuredList [-] ", err)
	for _, item := range podList.Items {
		fmt.Printf("namespace: %s, pod name: %s\n", item.Namespace, item.Name)
	}

}
