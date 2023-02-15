package dm1

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func Dm1() {
	fmt.Println("1. Dm1 func.")
	// 1. 加载kube-config配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/sgr/.kube/config")
	if err != nil {
		log.Fatalf("get kube config error, %s\n", err.Error())
	}
	fmt.Println(config.APIPath)
	// 2. 配置api路径
	config.APIPath = "api"
	// 3. 配置组名称
	config.GroupVersion = &corev1.SchemeGroupVersion

	config.NegotiatedSerializer = scheme.Codecs
	// 4. 实例化config对象
	client, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalf("get rest error, %s\n", err.Error())
	}

	result := corev1.PodList{}
	ctx := context.TODO()

	err = client.
		Get().
		Namespace("kube-system").
		Resource("pods").
		//VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	if err != nil {
		log.Println("get pod error, ", err)
	}

	for _, item := range result.Items {
		//log.Println(item.Namespace, "=> ", item.Name)
		log.Printf("namespace: %v, name: %v\n", item.Namespace, item.Name)
	}
}
