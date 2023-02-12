package asset

import (
	"encoding/json"
	"fmt"
	types "k8sclient/api"
)

func DecodeJsontoObject(jsonData []byte) error {

	var myObject types.RedisReplication

	// Decode the JSON data into the struct
	if err := json.Unmarshal(jsonData, &myObject); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// fmt.Println(myObject)
	return nil

}
