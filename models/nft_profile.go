package models

type NFTProfile struct {
	Id             int64  `gorm:"primaryKey;autoIncrement"`
	TokenBasicName string `gorm:"uniqueIndex:idx_name_token;size:64;not null"`
	NftTokenId     string `gorm:"uniqueIndex:idx_name_token;type:varchar(64);not null"`
	Name           string `gorm:"size:64;not null"`
	Url            string `gorm:"type:varchar(256)"`
	Image          string `gorm:"type:varchar(256);not null"`
	Description    string `gorm:"type:varchar(256)"`
	Text           string `gorm:"type:text"`
}
