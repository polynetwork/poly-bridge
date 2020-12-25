package models

type Chain struct {
	ChainId uint64 `gorm:"primaryKey;type:bigint(20);not null"`
	Name    string `gorm:"size:64"`
	Height  uint64 `gorm:"type:bigint(20);not null"`
}

type SrcTransaction struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	ChainId uint64 `gorm:"type:bigint(20);not null"`
	State uint64 `gorm:"type:bigint(20);not null"`
	Time uint64 `gorm:"type:bigint(20);not null"`
	Fee uint64 `gorm:"type:bigint(20);not null"`
	Height uint64 `gorm:"type:bigint(20);not null"`
	User string `gorm:"type:varchar(66);not null"`
	DstChainId uint64 `gorm:"type:bigint(20);not null"`
	Contract string `gorm:"type:varchar(66);not null"`
	Key string `gorm:"type:varchar(8192);not null"`
	Param string `gorm:"type:varchar(8192);not null"`
	SrcTransfer     *SrcTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

type SrcTransfer struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	ChainId uint64 `gorm:"type:bigint(20);not null"`
	Time uint64 `gorm:"type:bigint(20);not null"`
	Asset string `gorm:"type:varchar(66);not null"`
	From string `gorm:"type:varchar(66);not null"`
	To string `gorm:"type:varchar(66);not null"`
	Amount uint64 `gorm:"type:bigint(20);not null"`
	DstChainId uint64 `gorm:"type:bigint(20);not null"`
	DstAsset string `gorm:"type:varchar(66);not null"`
	DstUser string `gorm:"type:varchar(66);not null"`
	SrcTransaction     *SrcTransaction `gorm:"foreignKey:Hash;references:Hash"`
}

type PolyTransaction struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	ChainId uint64 `gorm:"type:bigint(20);not null"`
	State uint64 `gorm:"type:bigint(20);not null"`
	Time uint64 `gorm:"type:bigint(20);not null"`
	Fee uint64 `gorm:"type:bigint(20);not null"`
	Height uint64 `gorm:"type:bigint(20);not null"`
	SrcChainId uint64 `gorm:"type:bigint(20);not null"`
	SrcHash  string `gorm:"size:66;not null"`
	DstChainId uint64 `gorm:"type:bigint(20);not null"`
	Key string `gorm:"type:varchar(8192);not null"`
	SrcTransaction *SrcTransaction `gorm:"foreignKey:SrcHash;references:Hash"`
	SrcTransaction0 *SrcTransaction `gorm:"foreignKey:SrcHash;references:Key"`
}

type DstTransaction struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	ChainId uint64 `gorm:"type:bigint(20);not null"`
	State uint64 `gorm:"type:bigint(20);not null"`
	Time uint64 `gorm:"type:bigint(20);not null"`
	Fee uint64 `gorm:"type:bigint(20);not null"`
	Height uint64 `gorm:"type:bigint(20);not null"`
	SrcChainId uint64 `gorm:"type:bigint(20);not null"`
	Contract string `gorm:"type:varchar(66);not null"`
	PolyHash  string `gorm:"size:66;not null"`
	DstTransfer     *DstTransfer `gorm:"foreignKey:Hash;references:Hash"`
}

type DstTransfer struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	ChainId uint64 `gorm:"type:bigint(20);not null"`
	Time uint64 `gorm:"type:bigint(20);not null"`
	Asset string `gorm:"type:varchar(66);not null"`
	From string `gorm:"type:varchar(66);not null"`
	To string `gorm:"type:varchar(66);not null"`
	Amount uint64 `gorm:"type:bigint(20);not null"`
	DstTransaction     *DstTransaction `gorm:"foreignKey:Hash;references:Hash"`
}

type WrapperTransaction struct {
	Hash         string `gorm:"primaryKey;size:66;not null"`
	User         string `gorm:"size:64"`
	SrcChainId   uint64 `gorm:"type:bigint(20);not null"`
	BlockHeight  uint64 `gorm:"type:bigint(20);not null"`
	Time         uint64 `gorm:"type:bigint(20);not null"`
	DstChainId   uint64 `gorm:"type:bigint(20);not null"`
	FeeTokenHash string `gorm:"size:66;not null"`
	FeeToken     *Token `gorm:"foreignKey:FeeTokenHash;references:Hash"`
	FeeAmount    uint64 `gorm:"type:bigint(20);not null"`
	Status   uint64 `gorm:"type:bigint(20);not null"`
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
	Name           string      `gorm:"size:64;not null"`
	Precision     uint64      `gorm:"type:bigint(20);not null"`
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
