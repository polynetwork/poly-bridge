package bridgeeffect

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
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
	"time"
)

type AssetDetail struct {
	BasicName  string
	TokenAsset []*DstChainAsset
	Difference *big.Int
	Precision  uint64
	Price      int64
	Amount_usd string
	Reason     string
}
type DstChainAsset struct {
	ChainId     uint64
	Hash        string
	TotalSupply *big.Int
	Balance     *big.Int
	Flow        *big.Int
}

func StartCheckAsset(dbCfg *conf.DBConfig,ipCfg *conf.IPPortConfig) error {
	logs.Info("StartCheckAsset,start startCheckAsset")
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		return err
	}
	resAssetDetails := make([]*AssetDetail, 0)
	extraAssetDetails := make([]*AssetDetail, 0)
	tokenBasics := make([]*models.TokenBasic, 0)
	res := db.
		Where("property = ?", 1).
		Preload("Tokens").
		Find(&tokenBasics)
	if res.Error != nil {
		return err
	}
	for _, basic := range tokenBasics {
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
				assetDetail.Reason = err.Error()
				logs.Info("chainId: %v, Hash: %v, err:%v", token.ChainId, token.Hash, err)
				balance = big.NewInt(0)
			}
			chainAsset.Balance = balance
			time.Sleep(time.Second)
			totalSupply, _ := common.GetTotalSupply(token.ChainId, token.Hash)
			if err != nil {
				assetDetail.Reason = err.Error()
				totalSupply = big.NewInt(0)
				logs.Info("chainId: %v, Hash: %v, err:%v ", token.ChainId, token.Hash, err)
			}
			if !inExtraBasic(token.TokenBasicName) && basic.ChainId == token.ChainId {
				totalSupply = big.NewInt(0)
			}
			//specialBasic
			totalSupply = specialBasic(token, totalSupply)
			chainAsset.TotalSupply = totalSupply
			chainAsset.Flow = new(big.Int).Sub(totalSupply, balance)
			totalFlow = new(big.Int).Add(totalFlow, chainAsset.Flow)
			dstChainAssets = append(dstChainAssets, chainAsset)
		}
		assetDetail.Price = basic.Price
		assetDetail.Precision = basic.Precision
		assetDetail.TokenAsset = dstChainAssets
		assetDetail.Difference = totalFlow
		assetDetail.BasicName = basic.Name
		if inExtraBasic(assetDetail.BasicName) {
			extraAssetDetails = append(extraAssetDetails, assetDetail)
			continue
		}
		if assetDetail.BasicName == "WBTC" {
			chainAsset := new(DstChainAsset)
			chainAsset.ChainId = basedef.O3_CROSSCHAIN_ID
			response, err := http.Get(ipCfg.WBTCIP)
			defer response.Body.Close()
			if err != nil || response.StatusCode != 200 {
				logs.Error("Get o3 WBTC err:", err)
				continue
			}
			body, _ := ioutil.ReadAll(response.Body)
			o3WBTC := struct {
				Balance *big.Int
			}{}
			json.Unmarshal(body, &o3WBTC)
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
	err = sendDing(resAssetDetails,ipCfg.DingIP)
	if err != nil {
		logs.Error("------------sendDingDINg err---------")
	}
	logs.Info("rightdata___")
	for _, assetDetail := range resAssetDetails {
		logs.Info(assetDetail.BasicName, assetDetail.Difference, assetDetail.Precision, assetDetail.Price, assetDetail.Amount_usd)
		for _, tokenAsset := range assetDetail.TokenAsset {
			logs.Info(fmt.Sprintf("%2v %-30v %-30v %-30v %-30v\n", tokenAsset.ChainId, tokenAsset.Hash, tokenAsset.TotalSupply, tokenAsset.Balance, tokenAsset.Flow))
		}
	}
	fmt.Println("wrongdata___")
	for _, assetDetail := range extraAssetDetails {
		if assetDetail.BasicName == "USDT" {
			chainAsset := new(DstChainAsset)
			chainAsset.ChainId = basedef.O3_CROSSCHAIN_ID
			response, err := http.Get(ipCfg.USDTIP)
			defer response.Body.Close()
			if err != nil || response.StatusCode != 200 {
				logs.Error("Get o3 USDT err:", err)
				continue
			}
			body, _ := ioutil.ReadAll(response.Body)
			o3USDT := struct {
				Balance *big.Int
			}{}
			json.Unmarshal(body, &o3USDT)
			chainAsset.ChainId = basedef.O3_CROSSCHAIN_ID
			chainAsset.Balance = o3USDT.Balance
			assetDetail.TokenAsset = append(assetDetail.TokenAsset, chainAsset)
		}
		for _, tokenAsset := range assetDetail.TokenAsset {
			logs.Info(fmt.Sprintf("%2v %-30v %-30v %-30v %-30v\n", tokenAsset.ChainId, tokenAsset.Hash, tokenAsset.TotalSupply, tokenAsset.Balance, tokenAsset.Flow))
		}
	}
	return nil
}
func inExtraBasic(name string) bool {
	extraBasics := []string{"BLES", "GOF", "LEV", "mBTM", "MOZ", "O3", "USDT", "STN", "XMPT"}
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
		x, _ := new(big.Int).SetString("1", 10)
		return new(big.Int).Mul(x, presion)
	}
	if token.TokenBasicName == "DAO" && token.ChainId == basedef.ETHEREUM_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1000", 10)
		return new(big.Int).Mul(x, presion)
	}
	if token.TokenBasicName == "DAO" && token.ChainId == basedef.HECO_CROSSCHAIN_ID {
		x, _ := new(big.Int).SetString("1000", 10)
		return new(big.Int).Mul(x, presion)
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
		return new(big.Int).Mul(x, presion)
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

func sendDing(assetDetail []*AssetDetail,dingUrl string) error {
	ss := "[poly_NB]\n"
	flag := false
	for _, assetDetail := range assetDetail {
		if assetDetail.Difference.Cmp(big.NewInt(0)) == 1 {
			usd, _ := decimal.NewFromString(assetDetail.Amount_usd)
			if usd.Cmp(decimal.NewFromInt32(10000)) == 1 {
				flag = true
				ss += fmt.Sprintf("【%v】totalflow:%v $%v\n", assetDetail.BasicName, decimal.NewFromBigInt(assetDetail.Difference, 0).Div(decimal.New(1, int32(assetDetail.Precision))).StringFixed(2), assetDetail.Amount_usd)
				if assetDetail.Reason != "" {
					ss += "err Reason:" + assetDetail.Reason + "\n"
				}
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
	if flag {
		err := common.PostDingtext(ss,dingUrl)
		return err
	}
	return nil
}
