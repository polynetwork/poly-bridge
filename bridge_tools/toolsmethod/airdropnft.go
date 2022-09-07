package toolsmethod

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
	"io/ioutil"
	"math/big"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strconv"
	"strings"
)

func initAirDropNft(db *gorm.DB) {
	logs.Info("*** start initAirDropNft ***")
	err := db.Debug().AutoMigrate(
		&models.AirDropNft{},
	)
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate AirDropInfo err:%v", err))
	}

	airDropRanks := make([]*models.AirDropRank, 0)
	db.Table("(?) as b", db.Table("(select @curRank := 0) as r, (?) as t",
		db.Model(&models.AirDropInfo{}).Select("sum(amount) as sum_amount, bind_addr").Group("bind_addr").Order("sum_amount desc, bind_addr")).Select("t.sum_amount as amount,t.bind_addr,@curRank := @curRank + 1 as rank")).
		Where("b.rank <= 1000").
		Find(&airDropRanks)
	airDropNfts := make([]*models.AirDropNft, 0)
	for _, v := range airDropRanks {
		var bindChain uint64
		err = db.Model(&models.AirDropInfo{}).Select("bind_chain_id").Where("bind_addr = ?", v.BindAddr).Limit(1).Scan(&bindChain).
			Error
		if err != nil {
			logs.Error("Select bind_chain_id err,bind_addr:", v.BindAddr, err)
		}
		airDropNft := new(models.AirDropNft)
		airDropNft.BindAddr = v.BindAddr
		airDropNft.Amount = v.Amount
		airDropNft.Rank = v.Rank
		if basedef.IsETHChain(bindChain) {
			bindChain = getAirDropChain()
		}
		airDropNft.BindChainId = bindChain
		if v.Rank <= 100 {
			airDropNft.NftTbId = v.Rank - 1
		}
		airDropNft.NftDfId = v.Rank - 1
		airDropNfts = append(airDropNfts, airDropNft)
	}
	err = db.Save(airDropNfts).Error
	if err != nil {
		panic(fmt.Sprintf("Save airDropNfts err:%v", err))
	}
	logs.Info("*** end initAirDropNft ***")
}

func createipfsjson(nftCfg *conf.NftConfig, db *gorm.DB) {
	logs.Info("--------- start createipfsjson --------------------")
	if nftCfg == nil || nftCfg.TbName == "" || nftCfg.DfName == "" {
		panic(fmt.Sprintf("nftCfg is null"))
	}
	tbDescription := nftCfg.TbDescription
	dfDescription := nftCfg.DfDescription
	externalurl := nftCfg.ExternalUrl
	tbImage := nftCfg.TbImage
	dfImage := nftCfg.DfImage
	tbName := nftCfg.TbName
	txtTbName := strings.ReplaceAll(tbName, " ", "_")
	dfName := nftCfg.DfName
	txtDfName := strings.ReplaceAll(dfName, " ", "_")
	path := "../polynft"
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		logs.Error(err)
	}

	airDropNfts := make([]*models.AirDropNft, 0)
	db.Find(&airDropNfts)
	for _, v := range airDropNfts {
		if v.Rank <= 100 {
			nftJson := new(NftJson)
			nftJson.Description = tbDescription
			nftJson.ExternalUrl = externalurl
			nftJson.Image = tbImage
			nftJson.Name = tbName
			attributes := make([]*Attribute, 0)
			attributes = append(attributes,
				&Attribute{
					"name",
					tbName,
				},
				&Attribute{
					"rank",
					strconv.Itoa(int(v.Rank)),
				},
				&Attribute{
					"amount",
					fmt.Sprintf("$%.2f", float64(v.Amount)/10000.0),
				})
			nftJson.Attributes = attributes
			nftid := strconv.Itoa(int(v.NftTbId))
			data, _ := json.MarshalIndent(nftJson, "", "    ")
			err = ioutil.WriteFile(path+"/"+txtTbName+"_"+nftid, data, 0644)
			if err != nil {
				panic(fmt.Sprintf("WriteFile The bridge POLYNFT Error, addr:%v ,err:%v:", v.BindAddr, err))
			}
		}

		nftJson := new(NftJson)
		nftJson.Description = dfDescription
		nftJson.ExternalUrl = externalurl
		nftJson.Image = dfImage
		nftJson.Name = dfName
		attributes := make([]*Attribute, 0)
		attributes = append(attributes,
			&Attribute{
				"name",
				dfName,
			},
			&Attribute{
				"rank",
				strconv.Itoa(int(v.Rank)),
			},
			&Attribute{
				"amount",
				fmt.Sprintf("$%.2f", float64(v.Amount)/10000.0),
			})
		nftJson.Attributes = attributes
		nftid := strconv.Itoa(int(v.NftDfId))
		data, _ := json.MarshalIndent(nftJson, "", "    ")
		err = ioutil.WriteFile(path+"/"+txtDfName+"_"+nftid, data, 0644)
		if err != nil {
			panic(fmt.Sprintf("WriteFile Data Frigate POLYNFT Error, addr:%v ,err:%v:", v.BindAddr, err))
		}

	}
	logs.Info("********* end createipfsjson *********")
}

func signNft(nftCfg *conf.NftConfig, db *gorm.DB) {
	logs.Info("--------- start signNft --------------------")
	if nftCfg == nil || nftCfg.TbName == "" || nftCfg.DfName == "" {
		panic(fmt.Sprintf("nftCfg is null"))
	}
	if nftCfg.Pwd == "" {
		panic(fmt.Sprintf("nftCfgPwd is null"))
	}
	tbName := nftCfg.TbName
	txtTbName := strings.ReplaceAll(tbName, " ", "_")
	dfName := nftCfg.DfName
	txtDfName := strings.ReplaceAll(dfName, " ", "_")
	ipfsurl := nftCfg.IpfsUrl

	privateKeyBytes := hexutil.MustDecode(nftCfg.Pwd)
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("crypto.ToECDSA(privateKeyBytes)  Error:", err))
	}

	airDropNfts := make([]*models.AirDropNft, 0)
	db.Find(&airDropNfts)

	for _, v := range airDropNfts {
		if v.BindChainId != getAirDropChain() {
			continue
		}
		if v.Rank > 1000 {
			continue
		}

		if v.Rank <= 100 {
			//tokenId
			tbTokenId := big.NewInt(v.NftTbId)
			//user addr
			tbAccount := common.HexToAddress(v.BindAddr)
			//ipfs uri
			uri := ipfsurl + txtTbName + "_" + strconv.Itoa(int(v.NftTbId))
			hash := crypto.Keccak256Hash(
				common.BigToHash(tbTokenId).Bytes(),
				tbAccount[:],
				[]byte(uri),
			)
			// normally we sign prefixed hash
			// as in solidity with `ECDSA.toEthSignedMessageHash`
			prefixedHash := crypto.Keccak256Hash(
				[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
				hash.Bytes(),
			)
			// sign hash to validate later in Solidity
			sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
			if err != nil {
				panic(fmt.Sprint("crypto.Sign Error:", err))
			}
			v.NftTbSig = fmt.Sprintf("%x", sig)
		}

		dfTokenId := big.NewInt(v.NftDfId)
		dfAccount := common.HexToAddress(v.BindAddr)
		//ipfs uri
		uri := ipfsurl + txtDfName + "_" + strconv.Itoa(int(v.NftDfId))
		hash := crypto.Keccak256Hash(
			common.BigToHash(dfTokenId).Bytes(),
			dfAccount[:],
			[]byte(uri),
		)
		// normally we sign prefixed hash
		// as in solidity with `ECDSA.toEthSignedMessageHash`

		prefixedHash := crypto.Keccak256Hash(
			[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
			hash.Bytes(),
		)

		// sign hash to validate later in Solidity
		sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
		if err != nil {
			panic(fmt.Sprint("crypto.Sign Error:", err))
		}
		v.NftDfSig = fmt.Sprintf("%x", sig)

		err = db.Save(v).Error
		if err != nil {
			logs.Error("save sign nftUser err", err)
		}
	}

	logs.Info("********* end signNft *********")
}

func signOtherNft(nftCfg *conf.NftConfig, db *gorm.DB) {
	logs.Info("--------- start signOtherNft --------------------")
	if nftCfg == nil || nftCfg.TbName == "" || nftCfg.DfName == "" {
		panic(fmt.Sprintf("nftCfg is null"))
	}
	if nftCfg.Pwd == "" {
		panic(fmt.Sprintf("nftCfgPwd is null"))
	}
	tbName := nftCfg.TbName
	txtTbName := strings.ReplaceAll(tbName, " ", "_")
	dfName := nftCfg.DfName
	txtDfName := strings.ReplaceAll(dfName, " ", "_")
	ipfsurl := nftCfg.IpfsUrl

	privateKeyBytes := hexutil.MustDecode(nftCfg.Pwd)
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("crypto.ToECDSA(privateKeyBytes)  Error:", err))
	}

	airDropNfts := make([]*models.AirDropNft, 0)
	db.Where("nft_df_sig = ''").Or("nft_df_sig is null").Find(&airDropNfts)

	for _, v := range airDropNfts {
		if v.BindChainId != getAirDropChain() {
			continue
		}

		if v.Rank > 1000 {
			continue
		}
		if v.Rank <= 100 {
			//tokenId
			tbTokenId := big.NewInt(v.NftTbId)
			//user addr
			tbAccount := common.HexToAddress(v.BindAddr)
			//ipfs uri
			uri := ipfsurl + txtTbName + "_" + strconv.Itoa(int(v.NftTbId))
			hash := crypto.Keccak256Hash(
				common.BigToHash(tbTokenId).Bytes(),
				tbAccount[:],
				[]byte(uri),
			)
			// normally we sign prefixed hash
			// as in solidity with `ECDSA.toEthSignedMessageHash`
			prefixedHash := crypto.Keccak256Hash(
				[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
				hash.Bytes(),
			)
			// sign hash to validate later in Solidity
			sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
			if err != nil {
				panic(fmt.Sprint("crypto.Sign Error:", err))
			}
			v.NftTbSig = fmt.Sprintf("%x", sig)
		}

		dfTokenId := big.NewInt(v.NftDfId)
		dfAccount := common.HexToAddress(v.BindAddr)
		//ipfs uri
		uri := ipfsurl + txtDfName + "_" + strconv.Itoa(int(v.NftDfId))
		hash := crypto.Keccak256Hash(
			common.BigToHash(dfTokenId).Bytes(),
			dfAccount[:],
			[]byte(uri),
		)
		// normally we sign prefixed hash
		// as in solidity with `ECDSA.toEthSignedMessageHash`

		prefixedHash := crypto.Keccak256Hash(
			[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
			hash.Bytes(),
		)

		// sign hash to validate later in Solidity
		sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
		if err != nil {
			panic(fmt.Sprint("crypto.Sign Error:", err))
		}
		v.NftDfSig = fmt.Sprintf("%x", sig)

		err = db.Save(v).Error
		if err != nil {
			logs.Error("save sign nftUser err", err)
		}
	}

	logs.Info("********* end signOtherNft *********")
}

func getAirDropChain() uint64 {
	airdropChain := basedef.ETHEREUM_CROSSCHAIN_ID
	if basedef.ENV == basedef.TESTNET {
		airdropChain = basedef.RINKEBY_CROSSCHAIN_ID
	}
	return airdropChain
}
