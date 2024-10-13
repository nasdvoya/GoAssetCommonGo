package configuration

import (
	"assetcommongo/sdk/configuration"
	"encoding/json"
	"fmt"
)

type Tag struct {
	configuration.TagBase
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

type AssetBase struct {
	MqttParameters configuration.MqttParameters `json:"MqttParameters"`
	Tags           []Tag                        `json:"Tags"`
}

func ReturnExampleConfig() string {
	data := `{
		"mqttParameters": {
			"address": "192.168.1.100",
			"port": "1883",
			"isAnonymous": false,
			"username": "user123",
			"password": "pass122"
		},
		"tags": [
			{
				"name": "Temperature",
				"route": 0,
				"publish": 1,
				"mqttTopics": "sensors/temp",
				"mqttRetain": true,
				"foo": 42,
				"bar": "somevalue"
			},
			{
				"name": "Humidity",
				"route": 1,
				"publish": 2,
				"mqttTopics": "sensors/humidity",
				"mqttRetain": false,
				"foo": 43,
				"bar": "othervalue"
			}
		]
	}`
	return data
}

func DeserializeExampleConfig(data string) {

	var assetData AssetBase
	err := json.Unmarshal([]byte(data), &assetData)
	if err != nil {
		fmt.Print("Error deserializing AssetBase JSON: ", err)
		return
	}

	indentedJson, err := json.MarshalIndent(assetData, "", " ")
	if err != nil {
		fmt.Print("Eror generating indented JSON: ", err)
		return
	}
	fmt.Println(string(indentedJson))
}
