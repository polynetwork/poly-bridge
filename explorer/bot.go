/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package explorer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"poly-bridge/models"
	"poly-bridge/utils/net"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// Deduplicate alarms
var ALARMS = map[string]struct{}{}
var LOCAL_IPV4 string

type BotController struct {
	web.Controller
	Conf *conf.Config
}

func init() {
	LOCAL_IPV4 = getLocalIp()
	logs.Info("localIPV4：%s", LOCAL_IPV4)
}

func getLocalIp() string {
	ips, err := net.GetLocalIPv4s()
	if err != nil {
		logs.Error("get local IP error: %s", err)
		panic("get local IP error: " + err.Error())
	}
	if len(ips) == 0 {
		logs.Error("local IPv4s reading error")
		panic("local IPv4s reading error")
	}

	return ips[0]
}

func (c *BotController) BotPage() {
	var err error
	pageNo, _ := strconv.Atoi(c.Ctx.Input.Query("page_no"))
	pageSize, _ := strconv.Atoi(c.Ctx.Input.Query("page_size"))
	from, _ := strconv.Atoi(c.Ctx.Input.Query("from"))
	if pageSize == 0 {
		pageSize = 10
	}

	txs, count, err := c.getTxs(pageSize, pageNo, from, nil)
	if err == nil {
		// Check fee
		hashes := make([]string, len(txs))
		for i, tx := range txs {
			hashes[i] = tx.SrcHash
		}
		fees, checkFeeError := c.checkFees(hashes)
		if checkFeeError != nil {
			err = checkFeeError
		} else {

			rows := make([]string, len(txs))
			for i, entry := range txs {
				tx := models.ParseBotTx(entry, fees)
				rows[i] = fmt.Sprintf(
					fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%v</td>", 13)),
					tx.Asset,
					tx.Amount,
					tx.SrcChainName,
					tx.DstChainName,
					tx.Hash,
					tx.FeeToken,
					tx.FeePaid,
					tx.FeeMin,
					tx.FeePass,
					tx.Status,
					tx.Time,
					tx.Duration,
					tx.PolyHash,
				)
			}
			pages := count / pageSize
			if count%pageSize != 0 {
				pages++
			}

			rb := []byte(
				fmt.Sprintf(
					`<html><body><h1>Poly transaction status</h1>
					<div>total %d transactions (page %d/%d current page size %d)</div>
						<table style="width:100%%">
						<tr>
							<th>Asset</th>
							<th>Amount</th>
							<th>From</th>
							<th>To</th>
							<th>Hash</th>
							<th>FeeToken</th>
							<th>FeePaid</th>
							<th>FeeMin</th>
							<th>FeePass</th>
							<th>Status</th>
							<th>Time</th>
							<th>Duration</th>
							<th>PolyHash</th>
						</tr>
						%s
						</table>
				</body></html>`,
					count, pageNo, pages, len(txs), strings.Join(rows, "\n"),
				),
			)
			if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
				c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
			}
			c.Ctx.Output.Body(rb)
			return
		}
	}
	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()

}

func (c *BotController) FinishTx() {
	tx := c.Ctx.Input.Query("tx")
	token := c.Ctx.Input.Query("token")
	status := c.Ctx.Input.Query("status")
	var err error
	if status != "skip" && status != "wait" {
		err = fmt.Errorf("Invalid status")
	} else if token == conf.GlobalConfig.BotConfig.ApiToken {
		target := basedef.STATE_WAIT
		if status == "skip" {
			target = basedef.STATE_SKIP
		}
		err = db.Table("wrapper_transactions").Where("hash in ?", []string{tx}).Update("status", target).Error
	} else {
		err = fmt.Errorf("Access denied")
	}
	resp := fmt.Sprintf("Success: %s as %s", tx, status)
	if err != nil {
		resp = fmt.Sprintf("Tx %s Error %s", tx, err.Error())
	}
	c.Data["json"] = models.MakeErrorRsp(resp)
	c.ServeJSON()
}

func (c *BotController) CheckFees() {
	hashes := []string{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hashes)
	if err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	result, err := c.checkFees(hashes)
	if err == nil {
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
}

func (c *BotController) checkFees(hashes []string) (fees map[string]models.CheckFeeResult, err error) {
	wrapperTransactionWithTokens := make([]*models.WrapperTransactionWithToken, 0)
	err = db.Table("wrapper_transactions").Where("hash in ?", hashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&wrapperTransactionWithTokens).Error
	if err != nil {
		return
	}
	o3Hashes := []string{}
	for _, tx := range wrapperTransactionWithTokens {
		if tx.DstChainId == basedef.O3_CROSSCHAIN_ID {
			o3Hashes = append(o3Hashes, tx.Hash)
		}
	}
	if len(o3Hashes) > 0 {
		srcHashes, err := getSwapSrcTransactions(o3Hashes)
		o3srcs := []string{}
		for _, v := range srcHashes {
			o3srcs = append(o3srcs, v)
		}

		o3txs := make([]*models.WrapperTransactionWithToken, 0)
		err = db.Table("wrapper_transactions").Where("hash in ?", hashes).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&o3txs).Error
		if err != nil {
			return nil, err
		}
		wrapperTransactionWithTokens = append(wrapperTransactionWithTokens, o3txs...)
	}

	chainFees := make([]*models.ChainFee, 0)
	db.Preload("TokenBasic").Find(&chainFees)
	chain2Fees := make(map[uint64]*models.ChainFee, 0)
	for _, chainFee := range chainFees {
		chain2Fees[chainFee.ChainId] = chainFee
	}

	fees = make(map[string]models.CheckFeeResult, 0)
	for _, tx := range wrapperTransactionWithTokens {
		if tx.DstChainId == basedef.O3_CROSSCHAIN_ID {
			continue
		}
		chainId := tx.DstChainId
		if chainId == basedef.O3_CROSSCHAIN_ID {
			chainId = tx.SrcChainId
		}

		chainFee, ok := chain2Fees[chainId]
		if !ok {
			logs.Error("Failed to find chain fee for %d", tx.DstChainId)
			continue
		}

		if tx.FeeAmount == nil || tx.FeeToken == nil || tx.FeeToken.TokenBasic == nil {
			continue
		}

		x := new(big.Int).Mul(&tx.FeeAmount.Int, big.NewInt(tx.FeeToken.TokenBasic.Price))
		feePay := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(tx.FeeToken.Precision))))
		feePay = new(big.Float).Quo(feePay, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		feeMin := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		feeMin = new(big.Float).Quo(feeMin, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))
		res := models.CheckFeeResult{}
		if feePay.Cmp(feeMin) >= 0 {
			res.Pass = true
		}
		res.Paid, _ = feePay.Float64()
		res.Min, _ = feeMin.Float64()
		fees[tx.Hash] = res
	}

	return
}

func getSwapSrcTransactions(o3Hashs []string) (map[string]string, error) {
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	res := db.Table("dst_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").
		Where("dst_transactions.hash in ?", o3Hashs).
		Joins("inner join poly_transactions on dst_transactions.poly_hash = poly_transactions.hash").
		Joins("inner join src_transactions on poly_transactions.src_hash = src_transactions.hash").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		return nil, res.Error
	}
	checkHashes := make(map[string]string, 0)
	for _, srcPolyDstRelation := range srcPolyDstRelations {
		checkHashes[srcPolyDstRelation.DstHash] = srcPolyDstRelation.SrcHash
	}
	return checkHashes, nil
}

func (c *BotController) GetTxs() {
	var err error
	pageNo, _ := strconv.Atoi(c.Ctx.Input.Query("page_no"))
	pageSize, _ := strconv.Atoi(c.Ctx.Input.Query("page_size"))
	from, _ := strconv.Atoi(c.Ctx.Input.Query("from"))
	if pageSize == 0 {
		pageSize = 10
	}

	txs, count, err := c.getTxs(pageSize, pageNo, from, nil)
	if err == nil {
		// Check fee
		hashes := make([]string, len(txs))
		for i, tx := range txs {
			hashes[i] = tx.SrcHash
		}
		fees, checkFeeError := c.checkFees(hashes)
		if checkFeeError != nil {
			err = checkFeeError
		} else {
			c.Data["json"] = models.MakeBottxsRsp(pageSize, pageNo,
				(count+pageSize-1)/pageSize, count, txs, fees)
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
}

func (c *BotController) getTxs(pageSize, pageNo, from int, skip []uint64) ([]*models.SrcPolyDstRelation, int, error) {
	skips := append(skip, basedef.STATE_FINISHED, basedef.STATE_SKIP)
	srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
	tt := time.Now().Unix()
	end := tt - conf.GlobalConfig.EventEffectConfig.HowOld
	if from == 0 {
		from = 3
	}
	endBsc := tt - conf.GlobalConfig.EventEffectConfig.HowOld2
	query := db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("wrapper_transactions.status NOT IN ?", skips). // Where("dst_transactions.hash is null").Where("src_transactions.standard = ?", 0).
		Where("src_transactions.time > ?", tt-24*60*60*int64(from)).
		Where("(wrapper_transactions.time < ?) OR (wrapper_transactions.time < ? AND ((wrapper_transactions.src_chain_id = ? and wrapper_transactions.dst_chain_id = ?) OR (wrapper_transactions.src_chain_id = ? and wrapper_transactions.dst_chain_id = ?)))", end, endBsc, basedef.BSC_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID).
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("inner join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken")
	res := query.
		Limit(pageSize).Offset(pageSize * pageNo).
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelations)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	var transactionNum int64
	err := query.Count(&transactionNum).Error
	if err != nil {
		return nil, 0, err
	}
	return srcPolyDstRelations, int(transactionNum), nil
}

func (c *BotController) CheckTxs() {
	err := c.checkTxs()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "Success"
	}
	c.ServeJSON()
}

func (c *BotController) RunChecks() {
	if conf.GlobalConfig.BotConfig == nil || conf.GlobalConfig.BotConfig.DingUrl == "" {
		panic("Invalid ding url")
	}
	interval := conf.GlobalConfig.BotConfig.Interval
	if interval == 0 {
		interval = 60 * 5
	}
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for _ = range ticker.C {
		var isCheckBot bool
		botIp, err := cacheRedis.Redis.Get(cacheRedis.TxCheckBot)
		if err != nil {
			//lock
			lock, err := cacheRedis.Redis.Lock(cacheRedis.TxCheckBot, LOCAL_IPV4, 2*time.Second*time.Duration(interval))
			if err != nil {
				return
			}
			if lock {
				isCheckBot = true
			}
		} else if botIp == LOCAL_IPV4 {
			_, err := cacheRedis.Redis.Expire(cacheRedis.TxCheckBot, 2*time.Second*time.Duration(interval))
			if err != nil {
				return
			}
			isCheckBot = true
		}

		if isCheckBot {
			err := c.checkTxs()
			if err != nil {
				logs.Error("check txs error %s", err)
			}
		}
	}
}

func (c *BotController) checkTxs() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "CoGroup panic captured: %s", debug.Stack())
		}
	}()

	from := conf.GlobalConfig.BotConfig.CheckFrom
	pageSize := 20
	pageNo := 0
	txs, _, err := c.getTxs(pageSize, pageNo, int(from), []uint64{basedef.STATE_WAIT})
	if err != nil {
		return err
	}
	hashes := make([]string, len(txs))
	for i, tx := range txs {
		hashes[i] = tx.SrcHash
	}
	fees, err := c.checkFees(hashes)
	if err != nil {
		return err
	}
	for _, tx := range txs {
		_, ok := ALARMS[tx.SrcHash]
		if ok {
			continue
		}
		ALARMS[tx.SrcHash] = struct{}{}
		entry := models.ParseBotTx(tx, fees)
		title := fmt.Sprintf("Asset %s(%s->%s): %s", entry.Asset, entry.SrcChainName, entry.DstChainName, entry.Status)
		body := fmt.Sprintf(
			"## %s\n- Amount %v\n- Time %v\n- Duration %v\n- Fee %v(%v min:%v)\n- Hash %v\n- Poly %v\n",
			title,
			entry.Amount,
			entry.Time,
			entry.Duration,
			entry.FeePass,
			entry.FeePaid,
			entry.FeeMin,
			entry.Hash,
			entry.PolyHash,
		)

		btns := []map[string]string{
			map[string]string{
				"title":     "ListAll",
				"actionURL": conf.GlobalConfig.BotConfig.DetailUrl,
			},
			map[string]string{
				"title":     "MarkAsSkipped",
				"actionURL": fmt.Sprintf("%stoken=%s&tx=%s&status=skip", conf.GlobalConfig.BotConfig.FinishUrl, conf.GlobalConfig.BotConfig.ApiToken, entry.Hash),
			},
			map[string]string{
				"title":     "MarkAsWaiting",
				"actionURL": fmt.Sprintf("%stoken=%s&tx=%s&status=wait", conf.GlobalConfig.BotConfig.FinishUrl, conf.GlobalConfig.BotConfig.ApiToken, entry.Hash),
			},
			map[string]string{
				"title":     "Open",
				"actionURL": "https://explorer.poly.network/tx/" + entry.Hash,
			},
		}

		err = c.PostDingCard(title, body, btns)
		if err != nil {
			logs.Error("Post dingtalk error %s", err)
		}

		if basedef.ENV == basedef.TESTNET {
			break
		}
	}

	/*
		title := fmt.Sprintf("### Total %d, page %d/%d page size %d", count, pageNo, pages, len(txs))
		list := make([]string, len(txs))
		for i, tx := range txs {
			pass := "Lack"
			fee, ok := fees[tx.SrcHash]
			if ok && fee.Pass {
				pass = "Pass"
			}
			tsp := time.Unix(int64(tx.WrapperTransaction.Time), 0).Format(time.RFC3339)
			list[i] = fmt.Sprintf("- %s %s fee_paid(%s) %v fee_min %v", tsp, tx.SrcHash, pass, fee.Paid, fee.Min)
		}
		body := strings.Join(list, "\n")
		return c.PostDing(title, body)
	*/
	return nil
}

func (c *BotController) PostDingCard(title, body string, btns interface{}) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "actionCard"
	card := map[string]interface{}{}
	card["title"] = title
	card["text"] = body
	card["hideAvatar"] = 0
	card["btns"] = btns
	payload["actionCard"] = card
	return c.postDing(payload)
}

func (c *BotController) PostDingMarkDown(title, body string) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "markdown"
	payload["markdown"] = map[string]string{
		"title": title,
		"text":  fmt.Sprintf("%s\n%s", title, body),
	}
	return c.postDing(payload)
}

func (c *BotController) postDing(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", conf.GlobalConfig.BotConfig.DingUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	logs.Info("PostDing response Body:", string(respBody))
	return nil
}
