package main

import (
	"assetcommongo/examples/configuration"
	"encoding/json"
	"fmt"
)

func main() {
	// TODO: Testing like this for now, need unit tests
	data := configuration.ReturnExampleConfig()

	var assetData configuration.AssetBase
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
