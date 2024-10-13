package configuration

import (
	"flag"
	"fmt"
)

type Asset interface {
	Commands(command string)
	AssetHelp()
}

type AssetBase struct {
	MqttParameters MqttParameters `json:"mqttParameters"`
	Tags           *[]interface{} `json:"tags"`
}

type MqttParameters struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	IsAnonymous bool   `json: "isAnonymous"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type TagBase struct {
	Name       string      `json:"name"`
	Route      RouteType   `json:"route"`
	Publish    PublishType `json:"publish"`
	MqttTopic  string      `json:"mqttTopics"`
	MqttRetain bool        `json:"mqttRetain"`
}

type PublishType int

const (
	Update PublishType = iota
	Cyclic
	Once
)

type RouteType int

const (
	ToSubscribers RouteType = iota
	ToOther
)

func Commands(a Asset) {
	helpFlag := flag.Bool("help", false, "Shows help information")
	runFlag := flag.Bool("run", false, "Runs the asset")
	defaultConfigFlag := flag.Bool("defaultConfig", false, "Shows default configuration")
	schemaFlag := flag.Bool("schema", false, "Shows the schema")

	flag.Parse()

	if *helpFlag {
		a.AssetHelp()
	}
	if *runFlag {
		fmt.Println("Running the asset... (implement logic here)")
	}
	if *defaultConfigFlag {
		fmt.Println("Default configuration:")
	}
	if *schemaFlag {
		fmt.Println("Schema generation not yet implemented.")
	}
}
