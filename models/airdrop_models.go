package models

import (
	"fmt"
	"poly-bridge/basedef"
	"poly-bridge/common"
	"poly-bridge/conf"
	"strconv"
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
				user.Amount = fmt.Sprintf("%.2f", float64(airDropRank.Amount)/10000.0)
				break
			}
		}
		airDropRsp.Users = append(airDropRsp.Users, user)
	}
	return airDropRsp
}

type AirDropClaimReq struct {
	AirDropAddrs []string
}

type AirDropClaimNft struct {
	AirDropAddr     string
	NftTbId         int64
	NftDfId         int64
	NftTbSig        string
	NftDfSig        string
	NftTbContract   string
	NftDfContract   string
	NftTbIpfsUri    string
	NftDfIpfsUri    string
	NftTbOpenseaUrl string
	NftDfOpenseaUrl string
	IsClaimTb       bool
	IsClaimDf       bool
}

type AirDropClaimRsp struct {
	AirDropClaimNft []*AirDropClaimNft
}

func MakeAirDropClaimRsp(airDropNfts []*AirDropNft) (*AirDropClaimRsp, map[int]bool) {
	claimFlag := make(map[int]bool, 0)
	airDropClaimRsp := new(AirDropClaimRsp)
	for i, v := range airDropNfts {
		airDropClaimNft := new(AirDropClaimNft)
		airDropClaimNft.AirDropAddr = basedef.Hash2Address(v.BindChainId, v.BindAddr)
		if v.Rank <= 100 {
			airDropClaimNft.NftTbId = v.NftTbId
			airDropClaimNft.IsClaimTb = v.IsClaimTb
			airDropClaimNft.NftTbContract = conf.GlobalConfig.NftConfig.TbContract
			if !v.IsClaimTb {
				_, err := common.GetNftOwner(basedef.ETHEREUM_CROSSCHAIN_ID, airDropClaimNft.NftTbContract, int(airDropClaimNft.NftTbId))
				if err != nil {
					airDropClaimNft.NftTbSig = v.NftTbSig
					airDropClaimNft.NftTbIpfsUri = conf.GlobalConfig.NftConfig.IpfsUrl + conf.GlobalConfig.NftConfig.TbName + "_" + strconv.Itoa(int(v.NftTbId))
				} else {
					v.IsClaimTb = true
					airDropClaimNft.IsClaimTb = true
					claimFlag[i] = true
				}
			}
		}
		if v.IsClaimTb {
			airDropClaimNft.NftTbOpenseaUrl = conf.GlobalConfig.NftConfig.OpenseaUrl + conf.GlobalConfig.NftConfig.TbContract + "/" + strconv.Itoa(int(airDropClaimNft.NftTbId))
		}
		airDropClaimNft.NftDfId = v.NftDfId
		airDropClaimNft.IsClaimDf = v.IsClaimDf
		airDropClaimNft.NftDfContract = conf.GlobalConfig.NftConfig.DfContract
		if !v.IsClaimDf {
			_, err := common.GetNftOwner(basedef.ETHEREUM_CROSSCHAIN_ID, airDropClaimNft.NftDfContract, int(airDropClaimNft.NftDfId))
			if err != nil {
				airDropClaimNft.NftDfSig = v.NftDfSig
				airDropClaimNft.NftDfIpfsUri = conf.GlobalConfig.NftConfig.IpfsUrl + conf.GlobalConfig.NftConfig.DfName + "_" + strconv.Itoa(int(v.NftDfId))
			} else {
				v.IsClaimDf = true
				airDropClaimNft.IsClaimDf = true
				claimFlag[i] = true
			}
		}
		if v.IsClaimDf {
			airDropClaimNft.NftDfOpenseaUrl = conf.GlobalConfig.NftConfig.OpenseaUrl + conf.GlobalConfig.NftConfig.DfContract + "/" + strconv.Itoa(int(airDropClaimNft.NftDfId))
		}
		airDropClaimRsp.AirDropClaimNft = append(airDropClaimRsp.AirDropClaimNft, airDropClaimNft)
	}
	return airDropClaimRsp, claimFlag
}
