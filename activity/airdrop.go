package activity

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/models"
	"strconv"
	"time"
)

func (this *ActivityStats) StartAirDrop() {
	timeNow := time.Now().Unix()
	if timeNow < this.cfg.AirDropStartTime {
		time.Sleep(time.Second * time.Duration(this.cfg.AirDropStartTime-timeNow))
	}
	if timeNow > this.cfg.AirDropEndTime {
		logs.Info("AirDrop arrive endTime!!!!")
		return
	}
	logs.Info("start AirDropInfoStats")
	go this.run(this.cfg.AirDropInfoInterval, this.AirDropInfoStats)
}

func (this *ActivityStats) StartTokenPrice() {
	timeNow := time.Now().Unix()
	if timeNow < this.cfg.TokenPriceStartTime {
		time.Sleep(time.Second * time.Duration(this.cfg.TokenPriceStartTime-timeNow))
	}
	if timeNow > this.cfg.TokenPriceEndTime {
		logs.Info("TokenPrice arrive endTime!!!!")
		return
	}
	logs.Info("start TokenPriceAvgStats")
	go this.run(this.cfg.TokenPriceAvgInterval, this.TokenPriceAvgStats)
}

func (this *ActivityStats) AirDropInfoStats() (err error) {
	count, err := this.dao.GetAirDropInfoCount()
	if err != nil {
		logs.Error("AirDropInfoStats GetAirDropInfoCount err:", err)
		return err
	}
	if count == 0 {
		err = this.initialization()
		if err != nil {
			return
		}
	}
	lastSrcTx, err := this.dao.GetLastSrcTx()
	if err != nil || lastSrcTx == nil {
		return fmt.Errorf("GetLastSrcTx err: %v", err)
	}
	maxSrcTxId, err := this.dao.GetMaxSrcIdInAirDrop()
	if err != nil {
		return err
	}
	gapId := lastSrcTx.Id - maxSrcTxId
	if gapId <= 0 {
		return
	}
	txCount, err := this.dao.GetSrcCountWithIdAndChains(maxSrcTxId, lastSrcTx.Id, this.GetAirDropChain())
	if err != nil {
		return err
	}
	if txCount == 0 {
		return
	}
	for i := int64(0); i <= txCount-1/10; i++ {
		srcTxs, err := this.dao.GetSrcTxsWithIdAndChains(maxSrcTxId, lastSrcTx.Id, this.GetAirDropChain(), 10, int(i*10))
		if err != nil {
			return err
		}
		txHashes := make([]string, 0)
		for _, srcTx := range srcTxs {
			txHashes = append(txHashes, srcTx.Hash)
		}
		wrapperTxs, err := this.dao.GetWrapperTxsWithHashes(txHashes)
		if err != nil {
			return err
		}
		for _, srcTx := range srcTxs {
			airDropInfo := this.fillAirDropInfo(srcTx)
			for _, wrapperTx := range wrapperTxs {
				if srcTx.Hash == wrapperTx.Hash {
					token, err := this.dao.GetToken(wrapperTx.SrcChainId, wrapperTx.FeeTokenHash)
					if err == nil && token != nil {
						price := token.TokenBasic.Price
						tokenPrice, _ := this.dao.GetTokenPriceAvg(token.TokenBasicName)
						if tokenPrice != nil {
							price = tokenPrice.PriceAvg
						}
						x := new(big.Int).Mul(&wrapperTx.FeeAmount.Int, big.NewInt(price))
						y := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(token.Precision))))
						y = new(big.Float).Mul(y, new(big.Float).SetInt64(100))
						y = new(big.Float).Quo(y, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
						z := fmt.Sprintf("%.0f", y)
						amount, err := strconv.Atoi(z)
						if err == nil {
							airDropInfo.Amount += int64(amount)
						}
					}
					break
				}
			}
			dataAirDrop, _ := this.dao.GetAirDropByUser(airDropInfo.User)
			if dataAirDrop != nil {
				if !basedef.IsETHChain(dataAirDrop.BindChainId) && basedef.IsETHChain(airDropInfo.BindChainId) {
					dataAirDrop.BindChainId = airDropInfo.BindChainId
					dataAirDrop.BindAddr = airDropInfo.BindAddr
				}
				dataAirDrop.Amount += airDropInfo.Amount
				err = this.dao.UpdateAirDrop(dataAirDrop)
				if err != nil {
					logs.Error("UpdateAirDrop err", err)
					return err
				}
			} else {
				err = this.dao.SaveAirDrop(airDropInfo)
				if err != nil {
					logs.Error("SaveAirDrop err", err)
					return err
				}
			}
		}
	}
	return nil
}

func (this *ActivityStats) initialization() error {
	srcTxs, err := this.dao.GetSrcTxsWithTimeAndChains(this.cfg.AirDropStartTime, this.GetAirDropChain(), 1)
	if err != nil {
		return err
	}
	if len(srcTxs) == 0 || srcTxs[0] == nil {
		logs.Info("AirDropInfoStats first GetSrcTxsWithTimeAndChains len(srcTxs) is zero")
		return fmt.Errorf("AirDropInfoStats first GetSrcTxsWithTimeAndChains len(srcTxs) is zero")
	}
	airDropInfo := this.fillAirDropInfo(srcTxs[0])
	err = this.dao.SaveAirDrop(airDropInfo)
	return err
}

func (this *ActivityStats) fillAirDropInfo(srcTx *models.SrcTransaction) *models.AirDropInfo {
	airDropInfo := &models.AirDropInfo{
		User:        srcTx.User,
		ChainID:     srcTx.ChainId,
		IsEth:       basedef.IsETHChain(srcTx.ChainId),
		BindAddr:    srcTx.User,
		BindChainId: srcTx.ChainId,
		Amount:      0,
		SrcTxId:     srcTx.Id,
	}
	if !airDropInfo.IsEth {
		if basedef.IsETHChain(srcTx.DstChainId) && !basedef.IsETHChain(airDropInfo.BindChainId) {
			airDropInfo.BindChainId = srcTx.DstChainId
			airDropInfo.BindAddr = srcTx.SrcTransfer.DstUser
		}
	}
	chainFees, _ := this.dao.GetChainFeeTokens([]uint64{srcTx.ChainId})
	if len(chainFees) > 0 {
		if chainFees[0] != nil {
			feePresion := models.FeePrecison(chainFees[0].ChainId)
			feeAmount := srcTx.Fee
			price := chainFees[0].TokenBasic.Price

			tokenPrice, _ := this.dao.GetTokenPriceAvg(chainFees[0].TokenBasicName)
			if tokenPrice != nil && tokenPrice.PriceAvg > 0 {
				price = tokenPrice.PriceAvg
			}

			x := new(big.Int).Mul(&feeAmount.Int, big.NewInt(price))
			y := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(feePresion)))
			y = new(big.Float).Mul(y, new(big.Float).SetInt64(100))
			y = new(big.Float).Quo(y, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
			z := fmt.Sprintf("%.0f", y)
			amount, err := strconv.Atoi(z)
			if err == nil {
				airDropInfo.Amount += int64(amount)
			}
		}
	}
	return airDropInfo
}

func (this *ActivityStats) TokenPriceAvgStats() (err error) {
	currentTime := time.Now()
	nowDayStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	timeNow := currentTime.Unix()

	tokenPriceAvgs, err := this.dao.GetTokenPriceAvgs()
	if err != nil || tokenPriceAvgs == nil {
		return fmt.Errorf("TokenPriceAvgStats GetTokenPriceAvgs err")
	}
	tokenBasics, err := this.dao.GetTokenBasics()
	if err != nil || tokenBasics == nil {
		return fmt.Errorf("TokenPriceAvgStats GetTokenBasics err")
	}
	priceMap := make(map[string]int64, 0)
	for _, v := range tokenBasics {
		priceMap[v.Name] = v.Price
	}
	tokenHavePriceAvgMap := make(map[string]bool, 0)
	for _, v := range tokenPriceAvgs {
		tokenHavePriceAvgMap[v.Name] = true
	}

	for _, v := range tokenPriceAvgs {
		if price, ok := priceMap[v.Name]; ok {
			v.UpdateTime = timeNow
			v.PriceTotal += price
			v.PriceNumber++
			if timeNow > v.PriceTime+86400 {
				v.PriceAvg = v.PriceTotal / v.PriceNumber
				v.PriceTime = nowDayStartTime
			}
		}
	}
	for _, v := range tokenBasics {
		if _, ok := tokenHavePriceAvgMap[v.Name]; !ok {
			tokenPriceAvgs = append(tokenPriceAvgs, &models.TokenPriceAvg{
				Name:        v.Name,
				PriceAvg:    v.Price,
				UpdateTime:  timeNow,
				PriceTotal:  v.Price,
				PriceNumber: 1,
				PriceTime:   nowDayStartTime,
			})
		}
	}
	err = this.dao.SaveTokenPriceAvgs(tokenPriceAvgs)
	return
}

func (this *ActivityStats) GetAirDropChain() []uint64 {
	chain := make([]uint64, 0)
	for _, v := range basedef.ETH_CHAINS {
		chain = append(chain, v)
	}
	chain = append(chain, basedef.STARCOIN_CROSSCHAIN_ID, basedef.ONT_CROSSCHAIN_ID, basedef.NEO_CROSSCHAIN_ID, basedef.NEO3_CROSSCHAIN_ID)
	return chain
}
func (this *ActivityStats) IsAirDropChain(chain uint64) bool {
	for _, v := range this.GetAirDropChain() {
		if chain == v {
			return true
		}
	}
	return false
}
