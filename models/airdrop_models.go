package models

import (
	"fmt"
	"poly-bridge/basedef"
)

type AirDropInfo struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	User        string `gorm:"uniqueIndex;type:varchar(66);not null"`
	ChainID     uint64 `gorm:"type:bigint(20);not null"`
	IsEth       bool   `gorm:"type:int(8);not null"`
	BindAddr    string `gorm:"type:varchar(66);not null"`
	BindChainId uint64 `gorm:"type:bigint(20);not null"`
	Amount      int64  `gorm:"type:bigint(20);not null"`
	SrcTxId     int64  `gorm:"index:airdropinfos_srctxid;type:bigint(20);not null"`
	IsClaim     bool   `gorm:"type:int(8);not null"`
}

type AirDropRank struct {
	User        string
	ChainID     uint64
	BindAddr    string
	BindChainId uint64
	Amount      int64
	IsClaim     bool
	Rank        int64
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

func MakeAirDropRsp(addressReq AirDropReq, airDropRanks []*AirDropRank) *AirDropRsp {
	airDropRsp := new(AirDropRsp)
	for _, v := range addressReq.Users {
		user := &AirDropRspData{
			v.ChainId,
			basedef.Hash2Address(v.ChainId, v.Address),
			"",
			0,
			"0",
		}
		for _, airDropRank := range airDropRanks {
			if v.Address == airDropRank.User {
				if basedef.IsETHChain(airDropRank.BindChainId) {
					user.AirDropAddr = basedef.Hash2Address(airDropRank.BindChainId, airDropRank.BindAddr)
				}
				user.Rank = airDropRank.Rank
				user.Amount = fmt.Sprintf("%v", float64(airDropRank.Amount)/100.0)
				break
			}
		}
		airDropRsp.Users = append(airDropRsp.Users, user)
	}
	return airDropRsp
}
