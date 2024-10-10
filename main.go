package main

import (
	"GoAssetCommonGo/assetcommon"
	"encoding/json"
	"fmt"
)

func main() {
	data := `{
		"mqttParameters": {
			"address": "192.168.1.100",
			"port": "1883",
			"isAnonymous": false,
			"username": "user123",
			"password": "pass123"
		},
		"tags": [
			{
				"name": "Temperature",
				"route": 0,
				"publish": 1,
				"mqttTopics": "sensors/temp",
				"mqttRetain": true
			},
			{
				"name": "Humidity",
				"route": 1,
				"publish": 2,
				"mqttTopics": "sensors/humidity",
				"mqttRetain": false
			}
		]
	}`

	var asset assetcommon.AssetBase
	err := json.Unmarshal([]byte(data), &asset)

	if err != nil {
		fmt.Print("Error deserializing JSON: ", err)
	}

	fmt.Printf("Deserialized Asset with Tags: %+v\n", data)

}
