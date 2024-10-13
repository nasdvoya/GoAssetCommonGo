package main

import (
	"assetcommongo/examples/configuration"      // Import example configuration
	sdkConfig "assetcommongo/sdk/configuration" // Import the SDK configuration package
)

func main() {
	// Testing: Get and deserialize the example configuration
	data := configuration.ReturnExampleConfig()
	asset := configuration.DeserializeExampleConfig(data)

	// Call the Commands function from SDK and pass ExampleAsset
	sdkConfig.Commands(asset)
}
