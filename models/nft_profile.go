package models

type NFTProfile struct {
	TokenBasicName string `gorm:"primaryKey;size:64;not null"`
	NftTokenId     string `gorm:"primaryKey;type:varchar(64);not null"`
	Name           string `gorm:"size:64;not null"`
	Url            string `gorm:"size:64;not null"`
	Image          string `gorm:"size:64;not null"`
	Description    string `gorm:"type:varchar(256)"`
	Text           string `gorm:"type:text"`
}
