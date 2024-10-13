package main

import (
	"assetcommongo/examples/configuration" // Import example configuration
)

func main() {
	// Testing: Get and deserialize the example configuration
	data := configuration.ReturnExampleConfig()
	asset := configuration.DeserializeExampleConfig(data)

	// Call the Commands function from SDK and pass ExampleAsset
	asset.Commands("--run")
}
