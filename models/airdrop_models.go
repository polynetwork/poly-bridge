package models

import (
	"math/big"
	"poly-bridge/basedef"
)

type AirDropInfo struct {
	id           int64   `gorm:"primaryKey;autoIncrement"`
	User         string  `gorm:"uniqueIndex;type:varchar(66);not null"`
	ChainID      uint64  `gorm:"type:bigint(20);not null"`
	IsEth        bool    `gorm:"type:int(8);not null"`
	OntAddr      string  `gorm:"type:varchar(66)"`
	NeoAddr      string  `gorm:"type:varchar(66)"`
	Neo3Addr     string  `gorm:"type:varchar(66)"`
	StarCoinAddr string  `gorm:"type:varchar(66)"`
	Amount       *BigInt `gorm:"type:varchar(80);not null"`
	WrapperTxId  int64   `gorm:"type:bigint(20);not null"`
	SrcTxId      int64   `gorm:"type:bigint(20);not null"`
	Rank         int64   `gorm:"type:bigint(20);not null"`
	IsClaim      bool    `gorm:"type:int(8);not null"`
}

type TokenPriceAvg struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"uniqueIndex;size:64;not null"`
	PriceAvg    int64  `gorm:"type:bigint(20);not null"`
	UpdateTime  int64  `gorm:"type:bigint(20);not null"`
	PriceTotal  int64  `gorm:"type:bigint(20);not null"`
	PriceNumber int64  `gorm:"type:bigint(20);not null"`
	PriceTime   int64  `gorm:"type:bigint(20);not null"`
}

type AirDropReqData struct {
	ChainId uint64
	Address string
}

type AirDropReq struct {
	Users []AirDropReqData
}

type AirDropRspData struct {
	ChainId     uint64
	Address     string
	AirDropAddr string
	Rank        int64
	Amount      string
}

type AirDropRsp struct {
	Users []*AirDropRspData
}

func MakeAirDropRsp(addressReq AirDropReq, airDropInfos []*AirDropInfo) *AirDropRsp {
	airDropRsp := new(AirDropRsp)
	for _, v := range addressReq.Users {
		for _, airDropInfo := range airDropInfos {
			if v.Address == airDropInfo.User || v.Address == airDropInfo.OntAddr || v.Address == airDropInfo.NeoAddr ||
				v.Address == airDropInfo.Neo3Addr || v.Address == airDropInfo.StarCoinAddr {
				user := &AirDropRspData{
					v.ChainId,
					basedef.Hash2Address(v.ChainId, v.Address),
					basedef.Hash2Address(airDropInfo.ChainID, airDropInfo.User),
					airDropInfo.Rank,
					new(big.Int).Div(&airDropInfo.Amount.Int, big.NewInt(100)).String(),
				}
				if !basedef.IsETHChain(airDropInfo.ChainID) {
					user.AirDropAddr = ""
				}
				airDropRsp.Users = append(airDropRsp.Users, user)
				break
			}
		}
		user := &AirDropRspData{
			v.ChainId,
			basedef.Hash2Address(v.ChainId, v.Address),
			"",
			0,
			"0",
		}
		airDropRsp.Users = append(airDropRsp.Users, user)
	}
	return airDropRsp
}
