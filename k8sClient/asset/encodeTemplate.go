package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func SaveYaml(mapObject map[string]interface{}) error {

	// Marshal The Object To yaml
	yamlBytes, err := yaml.Marshal(mapObject)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Write the YAML to a file
	if err := ioutil.WriteFile("output.yaml", yamlBytes, 0644); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

func SaveJson(mapObject map[string]interface{}) error {

	// Marshal The Object To json
	jsonBytes, err := json.MarshalIndent(mapObject, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Write the json to a file
	if err := ioutil.WriteFile("output.json", jsonBytes, 0644); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = DecodeJsontoObject(jsonBytes)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil

}
