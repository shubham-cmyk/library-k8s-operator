package main

import (
	// custom "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	"context"
	"encoding/json"
	"fmt"
	"log"

	myapi "k8sclient/api"
	myk8s "k8sclient/k8s"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {

	// config, _ := generateK8sConfig()

	// myClient, _ := rest.HTTPClientFor(config)

	//	groupVersion := schema.GroupVersion{Group: "redis.redis.opstreelabs.in", Version: "v1beta1"}
	// myscheme := scheme.Scheme.Name()
	//myURl := generateK8sClient().RESTClient().Delete().URL()

	//apiPath := rest.DefaultVersionedAPIPath("", groupVersion)

	//request := rest.NewRequestWithClient(myURl, apiPath, rest.ClientContentConfig{GroupVersion: groupVersion}, myClient)

	// fmt.Println(output)
	// err := generateK8sClient().RESTClient().Get().
	// 	Name(replicationName).Namespace(replicationNamespace).
	// 	Resource("pod").Verb("GET").SpecificallyVersionedParams("",runtime.NewParameterCodec(),)
	// 	Do(context.TODO()).Error()

	//err := request.Name(replicationName).Namespace(replicationNamespace).Verb("GET").Resource("RedisReplication").Do(context.TODO()).Error()

	// group := generateK8sClient().RESTClient().APIVersion().Group
	// version := generateK8sClient().RESTClient().APIVersion().Version

	// fmt.Println(group)
	// fmt.Println(version)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// resp, err := generateK8sClient().Discovery().ServerResourcesForGroupVersion("redis.redis.opstreelabs.in/v1beta1")
	// if err != nil {
	// 	log.Fatalf("Failed to get server resources: %v", err)
	// }

	// for _, resource := range resp.APIResources {
	// 	if resource.Name == "redis-replication" && resource.Kind == "RedisReplication" {
	// 		resp, err := generateK8sClient().RESTClient().Get().
	// 			Namespace("default").Name("redis-replication").
	// 			Resource("redisreplication").VersionedParams(&resource.).
	// 			Do(context.TODO()).Raw()

	// 		if err != nil {
	// 			log.Fatalf("Failed to get custom resource: %v", err)
	// 		}

	// 		fmt.Println(resp)
	// 		break
	// 	}
	// }

	// yamlBytes, err := yaml.Marshal(mapObject)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Write the YAML to a file
	// if err := ioutil.WriteFile("output.yaml", yamlBytes, 0644); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	err := DecodeUsingDynamicClient()
	if err != nil {
		panic(err)
	}
}

func DecodeUsingDynamicClient() error {

	replicationName := "redis-replication"
	replicationNamespace := "default"

	customObject, err := myk8s.GenerateK8sDynamicClient().Resource(schema.GroupVersionResource{
		Group:    "redis.redis.opstreelabs.in",
		Version:  "v1beta1",
		Resource: "redisreplications",
	}).Namespace(replicationNamespace).Get(context.TODO(), replicationName, v1.GetOptions{})

	if err != nil {
		log.Fatalln(err, "Failed to Execute Get command", "replication name", replicationName, "namespace", replicationNamespace)
		return err
	} else {
		log.Default().Println("Successfully Execute the Get command")

	}

	//	mapObject := customObject.UnstructuredContent()
	myjson, err := customObject.MarshalJSON()
	if err != nil {
		fmt.Printf("Failed to Decode customObject.MarshalJSON")
		return err
	}
	// Save the Object To Yaml File
	// err = myasset.SaveYaml(mapObject)
	// if err != nil {
	// 	return err
	// }

	// Save the Object To Json File
	// err = myasset.SaveJson(mapObject)
	// if err != nil {
	// 	return err
	// }

	var replicationInstance myapi.RedisReplication
	// err = mapstructure.Decode(myjson, &replicationInstance)
	// Decode the JSON data into the struct
	if err := json.Unmarshal(myjson, &replicationInstance); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// if err != nil {
	// 	fmt.Printf("Failed to Decode")
	// 	return err
	// }

	fmt.Println(replicationInstance)
	return nil
}
