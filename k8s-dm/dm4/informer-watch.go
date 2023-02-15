package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	fmt.Println("start Dm4 ...")
	// 1. get kube-config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	// 2. use clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	factory := informers.NewSharedInformerFactory(clientSet, time.Second*60)

	informer := factory.Core().V1().Pods().Informer()
	//nodeInfomer := factory.Node().V1().RuntimeClasses().Informer()

	_, _ = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			newObj := obj.(metav1.Object)
			log.Printf("[+]New pod added: %s\n", newObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(metav1.Object)
			newPod := newObj.(metav1.Object)
			log.Printf("[!]Update pod status to store: %s -> %s \n", oldPod.GetName(), newPod.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			delPod := obj.(metav1.Object)
			log.Printf("[-]Delete pod: %s\n", delPod.GetName())
		},
	})

	stopCh := make(chan struct{})
	informer.Run(stopCh)
}
