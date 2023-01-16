package bridgedao

import (
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
	"testing"
)

func TestBridgeDao_WrapperTransactionCheckFee(t *testing.T) {
	cfg := conf.NewConfig("../../config.json")
	dao := NewBridgeDao(cfg.DBConfig, false)
	var wrapperTransactions []*models.WrapperTransaction
	dao.db.Model(&models.WrapperTransaction{}).Limit(1).Where("time = ?", 1654755372).Find(&wrapperTransactions)
	var srcTransactions []*models.SrcTransaction
	dao.db.Model(&models.SrcTransaction{}).Limit(1).Where("time = ?", 1654755372).Find(&srcTransactions)
	//dao.WrapperTransactionCheckFee()
	jsona, _ := json.MarshalIndent(wrapperTransactions, "", "	")
	fmt.Println(string(jsona))
	jsona, _ = json.MarshalIndent(srcTransactions, "", "	")
	fmt.Println(string(jsona))
	//fmt.Println(wrapperTransactions[0])
	//fmt.Println(srcTransactions[0])
	fmt.Println(dao.WrapperTransactionCheckFee(wrapperTransactions, srcTransactions))
	jsona, _ = json.MarshalIndent(wrapperTransactions, "", "	")
	fmt.Println(string(jsona))
	err := dao.UpdateEvents(wrapperTransactions, srcTransactions, nil, nil, nil, nil)
	fmt.Println("err", err)
}

func TestFloat64(t *testing.T) {
	//compare current result with gasPay in db
	gasPay := big.NewFloat(123456789.123456789)
	fmt.Println(gasPay.Float64())
	PaidGasFloat64, _ := gasPay.Float64()
	PaidGas := decimal.NewFromFloat(PaidGasFloat64).Mul(decimal.NewFromInt(100))
	fmt.Println(PaidGas)
	PaidGasDbBigInt := models.NewBigInt(PaidGas.BigInt())
	fmt.Println(PaidGasDbBigInt)
	PaidGasDecimal := decimal.NewFromBigInt(&PaidGasDbBigInt.Int, 0).Div(decimal.NewFromInt(100))
	fmt.Println(PaidGasDecimal)
	PaidGasFloat64New, _ := PaidGasDecimal.Float64()
	PaidGasDb := big.NewFloat(PaidGasFloat64New)
	fmt.Println(PaidGasDb, gasPay)
	fmt.Println(gasPay.Cmp(PaidGasDb))

}

func TestPointer(t *testing.T) {
	var curSrcTransaction *models.SrcTransaction
	if curSrcTransaction == nil {
		fmt.Println("aaa")
	}
}
