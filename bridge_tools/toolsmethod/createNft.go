package toolsmethod

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"math/big"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
	"strconv"
)

var db *gorm.DB

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NftJson struct {
	Description string       `json:"description"`
	ExternalUrl string       `json:"external_url"` //wang url
	Image       string       `json:"image"`        //imag url
	Name        string       `json:"name"`         //nft name
	Attributes  []*Attribute `json:"attributes"`
}

func Nft(cfg *conf.Config) {
	runflag := os.Getenv("runflag")
	if runflag == "" {
		panic(fmt.Sprintf("runflag is null "))
	}
	Logger := logger.Default
	dbCfg := cfg.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	x, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(fmt.Sprintf("database err", err))
	}
	db = x
	err = db.Debug().AutoMigrate(
		&models.NftUser{},
	)
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate NftUsers err", err))
	}

	nftCfg := cfg.NftConfig

	//create data
	if runflag == "0" {
		createNft()
		nftEffectAmount()
		updateColNftId()
		updateDfNftId()
	} else if runflag == "1" {
		createNft()
	} else if runflag == "2" {
		nftEffectAmount()
	} else if runflag == "3" {
		updateColNftId()
	} else if runflag == "4" {
		updateDfNftId()
	} else if runflag == "5" {
		createipfsjson(nftCfg)
	} else if runflag == "6" {
		signNft(nftCfg)
	}
}

func createNft() {
	logs.Info("--------- start createNft --------------------")
	var counter int
	err := db.Raw("select count(DISTINCT(t.addr)) from (select a.`from` as addr from src_transfers a inner join tokens b on a.chain_id =b.chain_id and a.asset=b.hash inner join src_transactions c on a.tx_hash = c.hash inner join token_basics d on b.token_basic_name = d.name  where a.`from`<> '' and  a.`from` is not null and a.chain_id <> 0 and d.price<>0 and b.precision<>0 and c.time<>0 and a.chain_id<>10 and c.time < 1640966400 group by a.`from`)t").
		Scan(&counter).Error
	if err != nil {
		panic(fmt.Sprint("Scan(&counter).Error:", err))
	}
	for i := 0; i < counter/100+1; i++ {
		users := make([]*models.NftUser, 0)
		//TxAmountUsd,FirstTime,Addrhash
		res := db.Raw("select a.`from` as addr_hash,convert(sum(a.amount*10000/POW(10,b.precision)*d.price/100000000),decimal(37,0)) as tx_amount_usd, min(c.time) as first_time  from src_transfers a inner join tokens b on a.chain_id =b.chain_id and a.asset=b.hash inner join src_transactions c on a.tx_hash = c.hash inner join token_basics d on b.token_basic_name = d.name  where a.`from`<> '' and  a.`from` is not null and a.chain_id <> 0 and d.price<>0 and b.precision<>0 and c.time<>0 and a.chain_id<>10 and c.time < 1640966400 group by a.`from` order by tx_amount_usd desc limit ? , ?", i*100, 100).
			Scan(&users)
		if res.Error != nil {
			panic(fmt.Sprint("Scan(&users).Error:", err))
		}
		if res.RowsAffected == 0 {
			logs.Info("i is %d,break", i)
			break
		}
		for _, user := range users {
			if user.TxAmountUsd.Cmp(big.NewInt(10000)) < 0 {
				continue
			}
			var chainId uint64
			err := db.Raw("SELECT a.chain_id from src_transactions a INNER JOIN src_transfers b on a.hash=b.tx_hash where a.time= ? and b.`from`= ?", user.FirstTime, user.AddrHash).
				First(&chainId).Error
			if err != nil {
				logs.Error("First(&chainId).Error", err)
			}
			if chainId == 0{
				continue
			}
			//ChainId
			user.ColChainId = chainId
			//Address
			user.ColAddress = basedef.Hash2Address(user.ColChainId, user.ColAddress)

			var num uint64
			err = db.Raw("select count(1) from src_transfers where chain_id<>10 and `from`= ?", user.AddrHash).
				First(&num).Error
			if err != nil {
				logs.Error("First(&num).Error", err)
			}
			//Txnum
			user.Txnum = num
			//EffectAmountUsd
			user.EffectAmountUsd = models.NewBigIntFromInt(0)
			err = db.Save(user).Error
			if err != nil {
				logs.Error("db.Save(user) err", err)
			}
		}
	}
	logs.Info("********* end createNft *********")
}

func nftEffectAmount() {
	logs.Info("--------- start effectAmountUsd --------------------")
	/*
		"c38072aa3f8e049de541223a9c9772132bb48634": "SHIB",
		"485cdbff08a4f91a16689e73893a11ae8b76af6d": "FEI",
		"4f99d10e16972ff2fe315eee53a95fc5a5870ce3": "BNB"
	*/
	type BuC struct {
		Addrhash  string
		AmountUsd *models.BigInt
	}

	chainId := 7
	tokenHash := []string{"c38072aa3f8e049de541223a9c9772132bb48634",
		"485cdbff08a4f91a16689e73893a11ae8b76af6d",
		"4f99d10e16972ff2fe315eee53a95fc5a5870ce3"}

	if basedef.ENV == basedef.TESTNET {
		tokenHash = []string{"33b89f811e8986c5b9d32114a21cc1fd5feb6c78"}
	}

	buCs := make([]*BuC, 0)
	err := db.Raw("select a.`to` as addrhash,convert(sum(a.amount/POW(10,14)*d.price/100000000),decimal(37,0)) as amount_usd from dst_transfers a inner join dst_transactions b on a.tx_hash=b.hash INNER JOIN tokens c on a.chain_id=c.chain_id and a.asset=c.hash left join token_basics d on c.token_basic_name=d.name where a.chain_id = ? and a.asset in ? and b.time < 1628744399 group by a.`to` ", chainId, tokenHash).
		Find(&buCs).Error
	if err != nil {
		panic(fmt.Sprint("Find(&addressAndAmounts),Error:", err))
	}
	for _, v := range buCs {
		outAmountUsd := new(models.BigInt)
		err := db.Raw("select convert(sum(a.amount/POW(10,14)*d.price/100000000),decimal(37,0)) as out_amount_usd from src_transfers a INNER JOIN src_transactions b on a.tx_hash=b.hash INNER JOIN tokens c on a.chain_id=c.chain_id and a.asset=c.hash left join token_basics d on c.token_basic_name=d.name where a.chain_id = ? and a.asset in ? and b.time < 1628744399 and `from` = ?", chainId, tokenHash, v.Addrhash).
			First(outAmountUsd).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("First(outAmountUsd).Error", err)
		}
		effectAmountUsd := new(big.Int).Sub(&v.AmountUsd.Int, &outAmountUsd.Int)
		if effectAmountUsd.Cmp(big.NewInt(0)) > 0 {
			//effectAmountUsd
			dfUser:=&models.NftUser{}
			err:=db.Where("addr_hash = ?", v.Addrhash).First(dfUser).
				Error
			if err!=nil{
				continue
			}
			dfUser.EffectAmountUsd=models.NewBigInt(effectAmountUsd)
			dfUser.DfChainId=7
			dfUser.DfAddress= basedef.Hash2Address(7, dfUser.AddrHash)
			err=db.Save(dfUser).Error
			if err != nil {
				logs.Error("Update stop_amount_usd err:%v,addrhash:%v,effectAmountUsd:%v", err, v.Addrhash, effectAmountUsd)
			}
		}
	}
	logs.Info("********* end effectAmountUsd *********")
}

func updateColNftId() {
	logs.Info("--------- start NftColId --------------------")
	chainIds := make([]uint64, 0)
	err := db.Raw("select col_chain_id from nft_users group BY col_chain_id ORDER BY col_chain_id").
		Find(&chainIds).Error
	if err != nil {
		logs.Error("updateColNftId Find(&chainIds) err", err)
	}
	nowNftColId := 0
	for _, v := range chainIds {
		var count int64
		err := db.Model(&models.NftUser{}).Where("col_chain_id = ?", v).Count(&count).Error
		if err != nil {
			panic(fmt.Sprint("Count(&count).Error:", err))
		}
		for i := 0; i < int(count)+1; i++ {
			nftUsers := make([]*models.NftUser, 0)
			db.Model(&models.NftUser{}).Where("col_chain_id = ?", v).Limit(100).Offset(100 * i).Find(&nftUsers)
			for _, nftUser := range nftUsers {
				nftUser.NftColId = nowNftColId
				nowNftColId++
			}
			if len(nftUsers) > 0 {
				err = db.Save(nftUsers).Error
				if err != nil {
					logs.Error("updateNftId Save(nftUsers).Error", err)
				}
			}
		}
	}
	logs.Info("--------- end NftColId --------------------")
}

func updateDfNftId() {
	logs.Info("--------- start NftDfId --------------------")
	chainIds := make([]uint64, 0)
	err := db.Raw("select df_chain_id from nft_users group BY df_chain_id ORDER BY df_chain_id").
		Find(&chainIds).Error
	if err != nil {
		logs.Error("updateDfNftId Find(&chainIds) err", err)
	}
	nowNftDfId := 0
	for _, v := range chainIds {
		var count int64
		err := db.Model(&models.NftUser{}).Where("df_chain_id = ? AND effect_amount_usd > 0", v).Count(&count).Error
		if err != nil {
			logs.Error(fmt.Sprint("chain:%v,Count Error:",v, err))
		}
		if count==0{
			continue
		}
		for i := 0; i < int(count)+1; i++ {
			nftUsers := make([]*models.NftUser, 0)
			db.Model(&models.NftUser{}).Where("df_chain_id = ? AND effect_amount_usd > 0", v).Limit(100).Offset(100 * i).Find(&nftUsers)
			for _, nftUser := range nftUsers {
				nftUser.NftDfId = nowNftDfId
				nowNftDfId++
			}
			if len(nftUsers) > 0 {
				err = db.Save(nftUsers).Error
				if err != nil {
					logs.Error("updateDfNftId Save(nftUsers).Error", err)
				}
			}
		}
	}
	logs.Info("--------- end NftDfId --------------------")

}

func createipfsjson(nftCfg *conf.NftConfig) {
	logs.Info("--------- start createipfsjson --------------------")
	if nftCfg == nil || nftCfg.ColName == "" || nftCfg.DfName == ""{
		panic(fmt.Sprintf("nftCfg is null"))
	}
	description := nftCfg.Description
	externalurl := nftCfg.ExternalUrl
	colImage := nftCfg.ColImage
	dfImage := nftCfg.DfImage
	colName := nftCfg.ColName
	dfName := nftCfg.DfName

	path := "../polynft"
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		logs.Error(err)
	}

	var count int64
	err = db.Model(&models.NftUser{}).Count(&count).Error
	if err != nil {
		panic(fmt.Sprint("Count(&count).Error:", err))
	}
	for i := 0; i < int(count)/100+1; i++ {
		nftUsers := make([]*models.NftUser, 0)
		db.Model(&models.NftUser{}).Limit(100).Offset(100 * i).Find(&nftUsers)
		for _, v := range nftUsers {
			nftJson := new(NftJson)
			nftJson.Description = description
			nftJson.ExternalUrl = externalurl
			nftJson.Image = colImage
			nftJson.Name = colName
			attributes := make([]*Attribute, 0)
			attributes = append(attributes,
				&Attribute{
					"Txnum",
					strconv.Itoa(int(v.Txnum)),
				},
				&Attribute{
					"FirstTime",
					strconv.FormatUint(v.FirstTime, 10),
				},
				&Attribute{
					"TxAmountUsd",
					decimal.NewFromBigInt(&v.TxAmountUsd.Int, -4).StringFixed(2),
				})
			nftJson.Attributes = attributes
			nftid := strconv.Itoa(v.NftColId)
			data, _ := json.Marshal(nftJson)
			err = ioutil.WriteFile(path+"/"+colName+"_"+nftid, data, 0644)
			if err != nil {
				panic(fmt.Sprint("WriteFile POLYNFT Error:", err))
			}
		}
	}

	count=0
	err = db.Model(&models.NftUser{}).Where("effect_amount_usd > 0").Count(&count).Error
	if err != nil {
		panic(fmt.Sprint("Count(&count).Error:", err))
	}
	for i := 0; i < int(count)/100+1; i++ {
		nftUsers := make([]*models.NftUser, 0)
		db.Model(&models.NftUser{}).Where("effect_amount_usd > 0").
			Limit(100).Offset(100 * i).Find(&nftUsers)
		for _, v := range nftUsers {
			nftJson := new(NftJson)
			nftJson.Description = description
			nftJson.ExternalUrl = externalurl
			nftJson.Image = dfImage
			nftJson.Name = dfName
			attributes := make([]*Attribute, 0)
			attributes = append(attributes,
				&Attribute{
					"EffectAmountUsd",
					decimal.NewFromBigInt(&v.EffectAmountUsd.Int, -4).StringFixed(2),
				},
				&Attribute{
					"Time",
					"2021-08-10",
				})
			nftJson.Attributes = attributes
			nftid := strconv.Itoa(v.NftDfId)
			data, _ := json.Marshal(nftJson)
			err = ioutil.WriteFile(path+"/"+dfName+"_"+nftid, data, 0644)
			if err != nil {
				panic(fmt.Sprint("WriteFile POLYNFT Error:", err))
			}
		}
	}
	logs.Info("********* end createipfsjson *********")
}

func signNft(nftCfg *conf.NftConfig) {
	logs.Info("--------- start signNft --------------------")
	if nftCfg == nil || nftCfg.ColName == "" || nftCfg.DfName == ""{
		panic(fmt.Sprintf("nftCfg is null"))
	}
	if nftCfg.Pwd == "" {
		panic(fmt.Sprintf("nftCfgPwd is null"))
	}
	colName := nftCfg.ColName
	dfName := nftCfg.DfName
	ipfsurl := nftCfg.IpfsUrl

	privateKeyBytes := hexutil.MustDecode(nftCfg.Pwd)
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("crypto.ToECDSA(privateKeyBytes)  Error:", err))
	}

	var count int64
	err = db.Model(&models.NftUser{}).Count(&count).Error
	if err != nil {
		panic(fmt.Sprint("Count(&count).Error:", err))
	}
	for i := 0; i < int(count)/100+1; i++ {
		nftUsers := make([]*models.NftUser, 0)
		db.Model(&models.NftUser{}).Limit(100).Offset(100 * i).Find(&nftUsers)
		for j, v := range nftUsers {
			//tokenId
			colTokenId := big.NewInt(int64(v.NftColId))
			dfTokenId := big.NewInt(int64(v.NftDfId))
			//user addr
			colAccount := common.HexToAddress(v.ColAddress)
			dfAccount := common.HexToAddress(v.DfAddress)
			//ipfs uri
			uri := ipfsurl + colName + "_" + strconv.Itoa(v.NftColId)
			hash := crypto.Keccak256Hash(
				common.BigToHash(colTokenId).Bytes(),
				colAccount[:],
				[]byte(uri),
			)
			// normally we sign prefixed hash
			// as in solidity with `ECDSA.toEthSignedMessageHash`

			// expect
			prefixedHash := crypto.Keccak256Hash(
				[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
				hash.Bytes(),
			)

			// sign hash to validate later in Solidity
			sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
			if err != nil {
				panic(fmt.Sprint("crypto.Sign Error:", err))
			}
			v.NftColsig = fmt.Sprintf("%x", sig)

			if v.EffectAmountUsd.Cmp(big.NewInt(0))>0{
				uri := ipfsurl + dfName + "_" + strconv.Itoa(v.NftDfId)
				hash := crypto.Keccak256Hash(
					common.BigToHash(dfTokenId).Bytes(),
					dfAccount[:],
					[]byte(uri),
				)
				// normally we sign prefixed hash
				// as in solidity with `ECDSA.toEthSignedMessageHash`

				// expect
				prefixedHash := crypto.Keccak256Hash(
					[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
					hash.Bytes(),
				)

				// sign hash to validate later in Solidity
				sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
				if err != nil {
					panic(fmt.Sprint("crypto.Sign Error:", err))
				}
				v.NftDfsig = fmt.Sprintf("%x", sig)
			}

			err = db.Save(v).Error
			if err != nil {
				logs.Error("save sign nftUser err", err)
			}

			if j == 0 {
				logs.Info("address: %v sig: %x  hash: %x  preFixedHash: %x  signer: %x  receiver: %x  tokenId: %d  uri %s",
					v.AddrHash, sig, hash, prefixedHash, crypto.PubkeyToAddress(privateKey.PublicKey), colAccount, colTokenId, uri)
			}

		}
	}
	logs.Info("********* end signNft *********")
}