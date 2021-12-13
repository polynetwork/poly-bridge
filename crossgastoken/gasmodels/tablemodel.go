package gasmodels

import bridge "poly-bridge/models"

type CrossGasTx struct {
	Id             int64          `gorm:"primaryKey;autoIncrement"`
	SrcHash        string         `gorm:"uniqueIndex;size:66;not null"`
	DstHash        string         `gorm:"size:66;not null"`
	SrcChainId     uint64         `gorm:"type:bigint(20);not null"`
	DstChainId     uint64         `gorm:"type:bigint(20);not null"`
	SrcAsset       string         `gorm:"type:varchar(120);not null"`
	DstAsset       string         `gorm:"type:varchar(120);not null"`
	SrcUser        string         `gorm:"type:varchar(66);not null"`
	SrcToAddress   string         `gorm:"type:varchar(66);not null"`
	DstUser        string         `gorm:"type:varchar(66);not null"`
	DstFromAddress string         `gorm:"type:varchar(66);not null"`
	Amount         *bridge.BigInt `gorm:"type:varchar(80);not null"`
	PolyFeeAmount  *bridge.BigInt `gorm:"type:varchar(80);not null"`
	TxFeeAmount    *bridge.BigInt `gorm:"type:varchar(80);not null"`
	SrcAssetPrice  *bridge.BigInt `gorm:"type:varchar(80);not null"`
	DstAssetPrice  *bridge.BigInt `gorm:"type:varchar(80);not null"`
	SrcToken       *bridge.Token  `gorm:"foreignKey:SrcAsset,SrcChainId;references:Hash,ChainId"`
	DstToken       *bridge.Token  `gorm:"foreignKey:DstAsset,DstChainId;references:Hash,ChainId"`
	SrcTime        uint64         `gorm:"index;type:bigint(20);not null"`
	DstTime        uint64         `gorm:"index;type:bigint(20);not null"`
	Status         uint64         `gorm:"type:bigint(20);not null"`
}

type CrossGasMap struct {
	Id           int64     `gorm:"primaryKey;autoIncrement"`
	SrcChainId   uint64    `gorm:"uniqueIndex:idx_native_map;type:bigint(20);not null"`
	SrcTokenHash string    `gorm:"uniqueIndex:idx_native_map;size:66;not null"`
	DstChainId   uint64    `gorm:"uniqueIndex:idx_native_map;type:bigint(20);not null"`
	DstTokenHash string    `gorm:"uniqueIndex:idx_native_map;size:66;not null"`
	SrcToken     *GasToken `gorm:"foreignKey:SrcTokenHash,SrcChainId;references:Hash,ChainId"`
	DstToken     *GasToken `gorm:"foreignKey:DstTokenHash,DstChainId;references:Hash,ChainId"`
	Property     int64     `gorm:"type:bigint(20);not null"`
}

type GasToken struct {
	Id              int64          `gorm:"primaryKey;autoIncrement"`
	Hash            string         `gorm:"uniqueIndex:idx_token;size:66;not null"`
	ChainId         uint64         `gorm:"uniqueIndex:idx_token;type:bigint(20);not null"`
	Name            string         `gorm:"size:64;not null"`
	Precision       uint64         `gorm:"type:bigint(20);not null"`
	Property        int64          `gorm:"type:bigint(20);not null"`
	AvailableAmount *bridge.BigInt `gorm:"type:varchar(64)"`
	TokenMaps       []*CrossGasMap `gorm:"foreignKey:SrcTokenHash,SrcChainId;references:Hash,ChainId"`
}

type GasChainResp struct {
	SrcChainId   uint64
	SrcTokenHash string
	SrcTokenName string
	DstChainId   uint64
	DstTokenHash string
	DstTokenName string
}

type GasFeeReq struct {
	SrcChainId   uint64
	SrcTokenHash string
	DstChainId   uint64
	DstTokenHash string
}
