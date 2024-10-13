package tag

import "assetcommongo/asset"

type Tag struct {
	asset.TagBase
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}
