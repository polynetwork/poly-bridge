package models

type Chain struct {
	ChainId uint64 `gorm:"primaryKey;type:bigint(20);not null"`
	Name    string `gorm:"size:64"`
	Height  uint64 `gorm:"type:bigint(20);not null"`
}

type Transaction struct {
	Hash         string `gorm:"primaryKey;size:66;not null"`
	User         string `gorm:"size:64"`
	SrcChainId   uint64 `gorm:"type:bigint(20);not null"`
	BlockHeight  uint64 `gorm:"type:bigint(20);not null"`
	Time         uint64 `gorm:"type:bigint(20);not null"`
	DstChainId   uint64 `gorm:"type:bigint(20);not null"`
	FeeTokenHash string `gorm:"size:66;not null"`
	FeeToken     *Token `gorm:"foreignKey:FeeTokenHash;references:Hash"`
	FeeAmount    uint64 `gorm:"type:bigint(20);not null"`
}

type TokenBasic struct {
	Name     string   `gorm:"primaryKey;size:64;not null"`
	CmcName  string   `gorm:"size:64;not null"`
	CmcPrice uint64   `gorm:"type:bigint(20);not null"`
	CmcInd   uint64   `gorm:"type:bigint(20);not null"`
	BinName  string   `gorm:"size:64;not null"`
	BinPrice uint64   `gorm:"type:bigint(20);not null"`
	BinInd   uint64   `gorm:"type:bigint(20);not null"`
	AvgPrice uint64   `gorm:"type:bigint(20);not null"`
	AvgInd   uint64   `gorm:"type:bigint(20);not null"`
	Time     uint64   `gorm:"type:bigint(20);not null"`
	Tokens   []*Token `gorm:"foreignKey:TokenBasicName;references:Name"`
}

type ChainFee struct {
	ChainId        uint64      `gorm:"primaryKey;type:bigint(20);not null"`
	TokenBasicName string      `gorm:"size:64;not null"`
	TokenBasic     *TokenBasic `gorm:"foreignKey:TokenBasicName;references:Name"`
	MaxFee         uint64      `gorm:"type:bigint(20);not null"`
	MinFee         uint64      `gorm:"type:bigint(20);not null"`
	ProxyFee       uint64      `gorm:"type:bigint(20);not null"`
}

type Token struct {
	Hash           string      `gorm:"primaryKey;size:66;not null"`
	ChainId        uint64      `gorm:"type:bigint(20);not null"`
	Name           string      `gorm:"size:64"`
	TokenBasicName string      `gorm:"size:64;not null"`
	TokenBasic     *TokenBasic `gorm:"foreignKey:TokenBasicName;references:Name"`
	TokenMaps      []*TokenMap `gorm:"foreignKey:SrcTokenHash;references:Hash"`
}

type TokenMap struct {
	SrcTokenHash string `gorm:"primaryKey;size:66;not null"`
	SrcToken     *Token `gorm:"foreignKey:SrcTokenHash;references:Hash"`
	DstTokenHash string `gorm:"primaryKey;size:66;not null"`
	DstToken     *Token `gorm:"foreignKey:DstTokenHash;references:Hash"`
}
