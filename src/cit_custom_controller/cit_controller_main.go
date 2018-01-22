/*
Copyright 2017

*/
package main

import (
	"cit_custom_controller/crd_controller"
	"flag"
	"github.com/golang/glog"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"sync"

	// import additional authentication providers - openid etc
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// GetClientConfig returns rest config, if path not specified assume in cluster config
func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func main() {

	glog.V(4).Infof("Starting Cit Custom Controller")
	kubeconf := flag.String("kubeconf", "", "Path to a kube config. Only required if out-of-cluster.")
	flag.Parse()

	config, err := GetClientConfig(*kubeconf)
	if err != nil {
		panic(err.Error())
	}

	cs, err := apiextcs.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//Create mutex to sync operation between the two CRD threads
	var crdmutex = &sync.Mutex{}

	plCrdDef := crd_controller.GetPLCrdDef()

	//crd_controller.NewCitCustomResourceDefinition to create PL CRD
	err = crd_controller.NewCitCustomResourceDefinition(cs, &plCrdDef)
	if err != nil {
		panic(err)
	}

	// Create a queue
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "citPLcontroller")

	plindexer, plinformer := crd_controller.NewPLIndexerInformer(config, queue, crdmutex)

	controller := crd_controller.NewCitPLController(queue, plindexer, plinformer)
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	glCrdDef := crd_controller.GetGLCrdDef()

	// note: if the CRD exist our CreateCRD function is set to exit without an error
	err = crd_controller.NewCitCustomResourceDefinition(cs, &glCrdDef)
	if err != nil {
		panic(err)
	}

	// Create a queue
	glQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "citGLcontroller")

	glindexer, glinformer := crd_controller.NewGLIndexerInformer(config, glQueue, crdmutex)

	geolocationController := crd_controller.NewCitGLController(glQueue, glindexer, glinformer)
	stopGl := make(chan struct{})
	defer close(stopGl)
	go geolocationController.Run(1, stopGl)

	glog.V(4).Infof("Waiting for updates on  Cit Custom Resource Definitions")

	// Wait forever
	select {}
}
