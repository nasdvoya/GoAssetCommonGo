package main

import (
	"assetcommongo/examples/configuration"
)

func main() {
	// TODO: Testing like this for now, need unit tests
	data := configuration.ReturnExampleConfig()
	asset := configuration.DeserializeExampleConfig(data)

	configuration.ShowHelp(asset)
}
