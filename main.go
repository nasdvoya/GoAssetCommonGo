package main

import (
	"GoAssetCommonGo/assetcommon"
	"fmt"
)

type ExampleTag struct {
	Name string
}

func (t ExampleTag) GetName() string {
	return t.Name
}

func main() {
	asset := assetcommon.NewAssetBaseWithTags[ExampleTag]("mqtt://example.com")

	asset.Tags = append(asset.Tags, ExampleTag{Name: "Tag1"})
	asset.Tags = append(asset.Tags, ExampleTag{Name: "Tag2"})

	fmt.Println("MqttParameters:", asset.AssetBase.MqttParameters)
	for _, tag := range asset.Tags {
		fmt.Println("Tag:", tag.GetName())
	}
}
