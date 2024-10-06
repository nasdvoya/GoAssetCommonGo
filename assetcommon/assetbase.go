package assetcommon

type TagBase interface {
	GetName() string
}

type AssetBase struct {
	MqttParameters string
}

type AssetBaseWithTags[TTag TagBase] struct {
	AssetBase
	Tags []TTag
}

func NewAssetBaseWithTags[TTag TagBase](mqttBase string) *AssetBaseWithTags[TTag] {
	return &AssetBaseWithTags[TTag]{
		AssetBase: AssetBase{MqttParameters: mqttBase},
		Tags:      []TTag{},
	}
}
