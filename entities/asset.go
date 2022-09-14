package entities

type Asset struct {
	AssetId          int64  `json:"asset_id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	AssetDescription string `json:"asset_description"`
}

type ResponseAsset struct {
	Status   bool
	Messagae string
	Data     Asset
}
