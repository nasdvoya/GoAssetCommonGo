package assetcommon

type MqttParameters struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	IsAnonymous bool   `json: "isAnonymous"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Tag struct {
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

type AssetBase struct {
	MqttParameters MqttParameters `json:"mqttParameters"`
	Tags           *[]Tag         `json:"tags"`
}
