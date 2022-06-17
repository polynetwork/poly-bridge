package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
	"testing"
)

func TestFeeController_NewCheckFee(t *testing.T) {

	//cfg := conf.NewConfig("../../config.json")

	conf.NewConfig("../config.json")
	Init()

	var mapCheckFeesReq map[string]*models.CheckFeeRequest = make(map[string]*models.CheckFeeRequest)
	mapCheckFeesReq["426e25c36f36b42377c04d75ca58f6a9422d1e9eb7999aca102b759ddf362ab6"] = &models.CheckFeeRequest{
		ChainId:                     2,
		TxId:                        "0000000000000000000000000000000000000000000000000000000000002fb5",
		PolyHash:                    "426e25c36f36b42377c04d75ca58f6a9422d1e9eb7999aca102b759ddf362ab6",
		Paid:                        0,
		PaidGas:                     0,
		Min:                         0,
		Status:                      0,
		SrcTransaction:              nil,
		WrapperTransactionWithToken: nil,
	}
	srcHashs := make([]string, 0)
	for k, v := range mapCheckFeesReq {
		if v.ChainId == basedef.SWITCHEO_CROSSCHAIN_ID || v.ChainId == basedef.ZILLIQA_CROSSCHAIN_ID {
			//switcheo || zilliqa
			v.Status = SKIP
			logs.Info("check fee poly_hash %s SKIP,is switcheo or zilliqa", k)
			continue
		}
		srcTransaction, err := checkFeeSrcTransaction(v.ChainId, v.TxId)
		if err != nil {
			//has not listen src_transaction
			v.Status = MISSING
			logs.Info("check fee poly_hash %s MISSING,hasn't src_Transaction %s", k, err)
			continue
		}
		if len(conf.PolyProxy) > 0 {
			if _, in := conf.PolyProxy[strings.ToUpper(srcTransaction.Contract)]; !in {
				//is not poly proxy
				v.Status = SKIP
				logs.Info("check fee poly_hash %s SKIP,is not poly proxy", k)
				continue
			}
		}
		v.SrcTransaction = srcTransaction
		srcHashs = append(srcHashs, srcTransaction.Hash)
	}
	checkFeewrapperTransaction(srcHashs, mapCheckFeesReq)
	jsona, _ := json.MarshalIndent(mapCheckFeesReq, "", "	")
	fmt.Println(string(jsona))
	for k, v := range mapCheckFeesReq {
		//check db
		if v.WrapperTransactionWithToken.IsPaid == true {
			logs.Info("check fee poly_hash %s marked as paid in db", k)
			v.Status = PAID
			continue
		}
	}
}
