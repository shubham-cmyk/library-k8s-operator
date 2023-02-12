package k8s

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/kubernetes"
)

// generateK8sClient create client for kubernetes
// func generateK8sClient() *kubernetes.Clientset {
// 	config, err := generateK8sConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return clientset
// }

func GenerateK8sDynamicClient() dynamic.Interface {
	config, err := generateK8sConfig()
	if err != nil {
		panic(err.Error())
	}
	dynamicClientset, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return dynamicClientset
}

// func generateK8sCustomClient() *custom.Clientset {
// 	config, err := generateK8sConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	customClientset, err := custom.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return customClientset
// }

// generateK8sConfig will load the kube config file
func generateK8sConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here
	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	return kubeConfig.ClientConfig()
}
