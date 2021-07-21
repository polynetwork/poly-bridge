package main

import (
	"encoding/json"
	"fmt"
	"github.com/polynetwork/poly-io-test/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"math/big"
	"net/http"
	"poly-bridge/basedef"
	"poly-bridge/common"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
)

type AssetDetail struct {
	BasicName  string
	TokenAsset []*DstChainAsset
	Difference *big.Int
	Precision  uint64
	Price      int64
	Amount_usd string
}
type DstChainAsset struct {
	ChainId     uint64
	Hash        string
	TotalSupply *big.Int
	Balance     *big.Int
	Flow        *big.Int
}

func startCheckAsset(dbCfg *conf.DBConfig) {
	log.Info("q-w-e-r-t start startCheckAsset")
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	resAssetDetails := make([]*AssetDetail, 0)
	extraAssetDetails := make([]*AssetDetail, 0)
	tokenBasics := make([]*models.TokenBasic, 0)
	res := db.
		Where("property = ?", 1).
		Preload("Tokens").
		Find(&tokenBasics)
	if res.Error != nil {
		panic(fmt.Errorf("load chainBasics faild, err: %v", res.Error))
	}
	//log.Info("q-w-e-r-t start to foreach tokenBasics")
	for _, basic := range tokenBasics {
		//log.Info(fmt.Sprintf("	for basicname: %v", basic.Name))
		assetDetail := new(AssetDetail)
		dstChainAssets := make([]*DstChainAsset, 0)
		totalFlow := big.NewInt(0)
		for _, token := range basic.Tokens {
			if notToken(token) {
				continue
			}
			chainAsset := new(DstChainAsset)
			chainAsset.ChainId = token.ChainId
			chainAsset.Hash = token.Hash
			balance, err := common.GetBalance(token.ChainId, token.Hash)
			if err != nil {
				log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, err:%v", token.ChainId, token.Hash, err))
				balance = big.NewInt(0)
				//panic(fmt.Errorf("q-w-e-r-t In CheckAsset Chain: %v,hash: %v , GetBalance faild, err: %v", token.ChainId, token.Hash, res.Error))
			}
			//log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, balance: %v", token.ChainId, token.Hash, balance.String()))
			chainAsset.Balance = balance
			//time sleep
			totalSupply, _ := common.GetTotalSupply(token.ChainId, token.Hash)
			if err != nil {
				totalSupply = big.NewInt(0)
				log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, err:%v ", token.ChainId, token.Hash, err))

				//panic(fmt.Errorf("q-w-e-r-t In CheckAsset Chain: %v,hash: %v , GetTotalSupply faild, err: %v", token.ChainId, token.Hash, res.Error))
			}
			if !inExtraBasic(token.TokenBasicName) && basic.ChainId == token.ChainId {
				totalSupply = big.NewInt(0)
			}
			//specialBasic
			totalSupply = specialBasic(token, totalSupply)
			chainAsset.TotalSupply = totalSupply
			chainAsset.Flow = new(big.Int).Sub(totalSupply, balance)
			//log.Info(fmt.Sprintf("	chainId: %v, Hash: %v, flow: %v", token.ChainId, token.Hash, chainAsset.flow.String()))
			totalFlow = new(big.Int).Add(totalFlow, chainAsset.Flow)
			dstChainAssets = append(dstChainAssets, chainAsset)
		}
		assetDetail.Price = basic.Price
		assetDetail.Precision = basic.Precision
		assetDetail.TokenAsset = dstChainAssets
		log.Info(fmt.Sprintf("	basic: %v,totalFlow: %v", basic.Name, totalFlow.String()))
		assetDetail.Difference = totalFlow
		assetDetail.BasicName = basic.Name
		if inExtraBasic(assetDetail.BasicName) {
			extraAssetDetails = append(extraAssetDetails, assetDetail)
			continue
		}
		if assetDetail.BasicName == "WBTC" {
			chainAsset := new(DstChainAsset)
			chainAsset.ChainId = basedef.O3_CROSSCHAIN_ID
			response, err := http.Get("http://124.156.209.180:9999/balance/0x6c27318a0923369de04df7Edb818744641FD9602/0x7648bDF3B4f26623570bE4DD387Ed034F2E95aad")
			defer response.Body.Close()
			if err != nil || response.StatusCode != 200 {
				log.Error("Get o3 WBTC err:", err)
				continue
			}
			body, _ := ioutil.ReadAll(response.Body)
			o3WBTC := struct {
				Balance *big.Int
			}{}
			json.Unmarshal(body, &o3WBTC)
			fmt.Println(o3WBTC.Balance)
			chainAsset.ChainId = basedef.O3_CROSSCHAIN_ID
			chainAsset.Flow = o3WBTC.Balance
			assetDetail.TokenAsset = append(assetDetail.TokenAsset, chainAsset)
			assetDetail.Difference.Add(assetDetail.Difference, chainAsset.Flow)
		}
		if assetDetail.Difference.Cmp(big.NewInt(0)) == 1 {
			assetDetail.Amount_usd = decimal.NewFromBigInt(assetDetail.Difference, 0).Div(decimal.New(1, int32(assetDetail.Precision))).Mul(decimal.New(assetDetail.Price, -8)).StringFixed(0)
		}

		resAssetDetails = append(resAssetDetails, assetDetail)
	}

	err = sendDing(resAssetDetails)
	if err != nil {
		fmt.Println("------------发送钉钉错误,错误---------")
	}
	fmt.Println("---准确数据---")
	for _, assetDetail := range resAssetDetails {
		fmt.Println(assetDetail.BasicName, assetDetail.Difference, assetDetail.Precision, assetDetail.Price, assetDetail.Amount_usd)
		for _, tokenAsset := range assetDetail.TokenAsset {
			fmt.Printf("%2v %-30v %-30v %-30v\n", tokenAsset.ChainId, tokenAsset.TotalSupply, tokenAsset.Balance, tokenAsset.flow)
		}
	}
	fmt.Println("---BU准确数据---")
	for _, assetDetail := range extraAssetDetails {
		fmt.Println(assetDetail.BasicName, assetDetail.Difference)
		for _, tokenAsset := range assetDetail.TokenAsset {
			fmt.Printf("%2v %-30v %-30v %-30v\n", tokenAsset.ChainId, tokenAsset.TotalSupply, tokenAsset.Balance, tokenAsset.flow)
		}
	}
}
func inExtraBasic(name string) bool {
	extraBasics := []string{"BLES", "GOF", "LEV", "mBTM", "MOZ", "O3", "STN", "USDT", "XMPT"}
	for _, basic := range extraBasics {
		if name == basic {
			return true
		}
	}
	return false
}
func specialBasic(token *models.Token, totalSupply *big.Int) *big.Int {
	presion, _ := new(big.Int).SetString("1000000000000000000", 10)
	if token.TokenBasicName == "YNI" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		return big.NewInt(0)
	}
	if token.TokenBasicName == "YNI" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1000000000000000000", 10)
		return x
	}
	if token.TokenBasicName == "DAO" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1000000000000000000000", 10)
		return x
	}
	if token.TokenBasicName == "DAO" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1000000000000000000000", 10)
		return x
	}
	if token.TokenBasicName == "COPR" && token.ChainId == basedef.BSC_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("274400000", 10)
		return new(big.Int).Mul(x, presion)
	}
	if token.TokenBasicName == "COPR" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("0", 10)
		return x
	}
	if token.TokenBasicName == "DigiCol ERC-721" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		return big.NewInt(0)
	}
	if token.TokenBasicName == "DigiCol ERC-721" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		return big.NewInt(0)
	}
	if token.TokenBasicName == "DMOD" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		return big.NewInt(0)
	}
	if token.TokenBasicName == "DMOD" && token.ChainId == basedef.BSC_CROSSCHAIN_ID {
		return new(big.Int).Mul(big.NewInt(15000000), presion)
	}
	if token.TokenBasicName == "SIL" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1487520675265330391631", 10)
		return x
	}
	if token.TokenBasicName == "SIL" && token.ChainId == basedef.BSC_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("5001", 10)
		return x
	}
	if token.TokenBasicName == "DOGK" && token.ChainId == basedef.BSC_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("0", 10)
		return x
	}
	if token.TokenBasicName == "DOGK" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("285000000000", 10)
		return new(big.Int).Mul(x, presion)
	}

	return totalSupply
}
func notToken(token *models.Token) bool {
	if token.TokenBasicName == "USDT" && token.Precision != 6 {
		return true
	}
	return false
}

func sendDing(assetDetail []*AssetDetail) error {
	ss := "[poly_NB]\n"
	for _, assetDetail := range assetDetail {
		if assetDetail.Difference.Cmp(big.NewInt(0)) == 1 {
			usd, _ := decimal.NewFromString(assetDetail.Amount_usd)
			if usd.Cmp(decimal.NewFromInt32(10000)) == 1 {
				ss += fmt.Sprintf("【%v】totalflow:%v $%v\n", assetDetail.BasicName, decimal.NewFromBigInt(assetDetail.Difference, 0).Div(decimal.New(1, int32(assetDetail.Precision))).StringFixed(2), assetDetail.Amount_usd)
				for _, x := range assetDetail.TokenAsset {
					ss += "ChainId: " + fmt.Sprintf("%v", x.ChainId) + "\n"
					ss += "Hash: " + fmt.Sprintf("%v", x.Hash) + "\n"
					ss += "TotalSupply: " + decimal.NewFromBigInt(x.TotalSupply, 0).Div(decimal.New(1, int32(assetDetail.Precision))).StringFixed(2) + " "
					ss += "Balance: " + decimal.NewFromBigInt(x.Balance, 0).Div(decimal.New(1, int32(assetDetail.Precision))).StringFixed(2) + " "
					ss += "Flow: " + decimal.NewFromBigInt(x.Flow, 0).Div(decimal.New(1, int32(assetDetail.Precision))).StringFixed(2) + "\n"
				}
			}
		}
	}
	err := common.PostDingtext(ss)
	return err
}
