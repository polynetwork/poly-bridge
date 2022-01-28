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
	"poly-bridge/utils/decimal"
	"poly-bridge/utils/fee"
	"poly-bridge/utils/net"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var LOCAL_IPV4 string

type BotController struct {
	web.Controller
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
			for i, tx := range txs {
				entry, err := getSrcPolyDstRelation(tx)
				if err != nil {
					logs.Error("getSrcPolyDstRelation of hash: %s err: %s", tx.SrcHash, err)
					continue
				}
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
	resp := ""
	if token == conf.GlobalConfig.BotConfig.ApiToken {
		switch status {
		case "skip":
			_, err := cacheRedis.Redis.Set(cacheRedis.MarkTxAsSkipPrefix+tx, "markAsSkipByBot", time.Hour*24*7)
			if err == nil {
				resp = fmt.Sprintf("Success mark %s as skip", tx)
			}
		case "wait":
			_, err = cacheRedis.Redis.Del(cacheRedis.MarkTxAsSkipPrefix + tx)
			if err == nil {
				resp = fmt.Sprintf("Success mark %s as wait", tx)
			}
		default:
			err = fmt.Errorf("Invalid status")
		}
	} else {
		err = fmt.Errorf("Access denied")
	}

	if err != nil {
		resp = fmt.Sprintf("Tx %s Error %s", tx, err.Error())
	}
	logs.Info(resp)
	c.Data["json"] = models.MakeErrorRsp(resp)
	c.ServeJSON()
}

func (c *BotController) MarkUnMarkTxAsPaid() {
	tx := c.Ctx.Input.Query("tx")
	token := c.Ctx.Input.Query("token")
	var err error
	resp := ""
	if token == conf.GlobalConfig.BotConfig.ApiToken {
		exists, _ := cacheRedis.Redis.Exists(cacheRedis.MarkTxAsPaidPrefix + tx)
		if exists {
			_, err = cacheRedis.Redis.Del(cacheRedis.MarkTxAsPaidPrefix + tx)
			if err == nil {
				resp = fmt.Sprintf("Success unmark %s as paid", tx)
			}
		} else {
			_, err := cacheRedis.Redis.Set(cacheRedis.MarkTxAsPaidPrefix+tx, "markAsPaidByBot", time.Hour*12)
			if err == nil {
				resp = fmt.Sprintf("Success mark %s as paid", tx)
			}
		}
	} else {
		err = fmt.Errorf("Access denied")
	}
	if err != nil {
		resp = fmt.Sprintf("Tx %s Error %s", tx, err.Error())
	}
	logs.Info(resp)
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

	srcTransactions := make([]*models.SrcTransaction, 0)
	err = db.Table("src_transactions").Where("hash in ?", hashes).Find(&srcTransactions).Error
	if err != nil {
		return
	}
	o3Hashes := []string{}
	for _, tx := range srcTransactions {
		if tx.ChainId == basedef.O3_CROSSCHAIN_ID {
			o3Hashes = append(o3Hashes, tx.Hash)
		}
	}

	o3SrcHash2DstChainId := make(map[string]uint64, 0)
	if len(o3Hashes) > 0 {
		o3SrcHash2DstChainId, err = getSwapSrcTransactions(o3Hashes)
		o3srcs := []string{}
		for hash, _ := range o3SrcHash2DstChainId {
			o3srcs = append(o3srcs, hash)
		}

		o3txs := make([]*models.WrapperTransactionWithToken, 0)
		err = db.Table("wrapper_transactions").Where("hash in ?", o3srcs).Preload("FeeToken").Preload("FeeToken.TokenBasic").Find(&o3txs).Error
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
		chainId := tx.DstChainId
		if chainId == basedef.O3_CROSSCHAIN_ID {
			chainId = o3SrcHash2DstChainId[tx.Hash]
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
		payFee := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.Int64FromFigure(int(tx.FeeToken.Precision))))
		payFee = new(big.Float).Quo(payFee, new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		x = new(big.Int).Mul(&chainFee.MinFee.Int, big.NewInt(chainFee.TokenBasic.Price))
		minFee := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt64(basedef.PRICE_PRECISION))
		minFee = new(big.Float).Quo(minFee, new(big.Float).SetInt64(basedef.FEE_PRECISION))
		minFee = new(big.Float).Quo(minFee, new(big.Float).SetInt64(basedef.Int64FromFigure(int(chainFee.TokenBasic.Precision))))

		// get optimistic L1 fee on ethereum
		if chainId == basedef.OPTIMISTIC_CROSSCHAIN_ID {
			ethChainFee, ok := chain2Fees[basedef.ETHEREUM_CROSSCHAIN_ID]
			if ok {
				l1MinFee, _, err := fee.GetL1Fee(ethChainFee, chainId)
				if err == nil {
					minFee = new(big.Float).Add(minFee, l1MinFee)
				}
			}
		}

		res := models.CheckFeeResult{}
		if payFee.Cmp(minFee) >= 0 {
			res.Pass = true
		}
		res.Paid, _ = payFee.Float64()
		res.Min, _ = minFee.Float64()
		fees[tx.Hash] = res
	}

	return
}

func getSwapSrcTransactions(o3Hashs []string) (map[string]uint64, error) {
	o3SrcTransaction := make([]*models.SrcTransaction, 0)
	err := db.Table("src_transactions").
		Where("src_transactions.hash in ?", o3Hashs).Find(&o3SrcTransaction).Error
	if err != nil {
		return nil, err
	}

	srcPolyDstRelation := make([]*models.SrcPolyDstRelation, 0)
	err = db.Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash").
		Where("dst_transactions.hash in ?", o3Hashs).
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Find(&srcPolyDstRelation).Error
	if err != nil {
		return nil, err
	}
	checkHashes := make(map[string]uint64, 0)
	for _, relation := range srcPolyDstRelation {
		for _, src := range o3SrcTransaction {
			if relation.DstHash == src.Hash {
				checkHashes[relation.SrcHash] = src.DstChainId
			}
		}
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
			c.Data["json"] = c.makeBottxsRsp(pageSize, pageNo,
				(count+pageSize-1)/pageSize, count, txs, fees)
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = err.Error()
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.ServeJSON()
}

func (c *BotController) getTxs(pageSize, pageNo, from int, skip []uint64) ([]*models.TxHashChainIdPair, int, error) {
	//skips := append(skip, basedef.STATE_FINISHED, basedef.STATE_SKIP)
	tt := time.Now().Unix()
	end := tt - conf.GlobalConfig.EventEffectConfig.HowOld
	if from == 0 {
		from = 3
	}
	endBsc := tt - conf.GlobalConfig.EventEffectConfig.HowOld2

	txs := make([]*models.TxHashChainIdPair, 0)
	var count int64

	var polyProxies []string
	for k, _ := range conf.PolyProxy {
		polyProxies = append(polyProxies, k)
	}
	query := db.Debug().Table("src_transactions").
		Select("src_transactions.hash as src_hash, src_transactions.chain_id as src_chain_id, src_transactions.dst_chain_id as dst_chain_id, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, wrapper_transactions.id as wrapper_id").
		Where("UPPER(src_transactions.contract) in ?", polyProxies).
		Where("src_transactions.time > ?", tt-24*60*60*int64(from)).
		Where("(src_transactions.time < ?) OR (src_transactions.time < ? and ((src_transactions.chain_id = ? and src_transactions.dst_chain_id = ?) or (src_transactions.chain_id = ? and src_transactions.dst_chain_id = ?)))", end, endBsc, basedef.BSC_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.HECO_CROSSCHAIN_ID, basedef.BSC_CROSSCHAIN_ID).
		Where("((select count(*) from poly_transactions where src_transactions.hash = poly_transactions.src_hash) = 0 OR (select count(*) from dst_transactions where poly_transactions.hash=dst_transactions.poly_hash) = 0)").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("left join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash")

	err := query.Limit(pageSize).Offset(pageSize * pageNo).Order("src_transactions.time desc").Find(&txs).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(txs); {
		hash := txs[i].SrcHash
		if (txs[i].SrcChainId == basedef.NEO_CROSSCHAIN_ID ||
			txs[i].DstChainId == basedef.NEO_CROSSCHAIN_ID ||
			txs[i].SrcChainId == basedef.NEO3_CROSSCHAIN_ID ||
			txs[i].DstChainId == basedef.NEO3_CROSSCHAIN_ID) && txs[i].WrapperId == 0 {
			count--
			txs = append(txs[:i], txs[i+1:]...)
			logs.Info("skip %s, because it is a NEO/NEO3 tx with no wrapper_transactions", hash)
			continue
		}

		exists, _ := cacheRedis.Redis.Exists(cacheRedis.MarkTxAsSkipPrefix + hash)
		if exists {
			count--
			txs = append(txs[:i], txs[i+1:]...)
			logs.Info("%s has been marked as a skip", hash)
		} else {
			i++
		}
	}
	return txs, int(count), nil
}

func (c *BotController) makeBottxsRsp(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*models.TxHashChainIdPair, fees map[string]models.CheckFeeResult) map[string]interface{} {
	rsp := map[string]interface{}{}
	rsp["PageSize"] = pageSize
	rsp["PageNo"] = pageNo
	rsp["TotalPage"] = totalPage
	rsp["TotalCount"] = totalCount
	txs := make([]models.BotTx, len(transactions))
	for i, tx := range transactions {
		srcPolyDstRelation, err := getSrcPolyDstRelation(tx)
		if err != nil {
			logs.Error("getSrcPolyDstRelation of hash: %s err: %s", tx.SrcHash, err)
			continue
		}
		txs[i] = models.ParseBotTx(srcPolyDstRelation, fees)
	}
	rsp["Transactions"] = txs
	return rsp
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

func getSrcPolyDstRelation(tx *models.TxHashChainIdPair) (*models.SrcPolyDstRelation, error) {
	hash := tx.SrcHash
	if tx.SrcChainId == basedef.O3_CROSSCHAIN_ID {
		originTx := new(models.TxHashChainIdPair)
		err := db.Debug().Table("src_transactions").
			Select("src_transactions.hash as src_hash, src_transactions.chain_id as src_chain_id, poly_transactions.hash as poly_hash").
			Where("dst_transactions.hash = ?", tx.SrcHash).
			Joins("LEFT JOIN poly_transactions on src_transactions.hash=poly_transactions.src_hash").
			Joins("LEFT JOIN dst_transactions on dst_transactions.poly_hash=poly_transactions.hash").
			Order("src_transactions.time desc").Find(&originTx).Error
		if err == nil {
			hash = originTx.SrcHash
		}
	}
	srcPolyDstRelation := new(models.SrcPolyDstRelation)
	err := db.Debug().Table("src_transactions").
		Select("src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash, wrapper_transactions.fee_token_hash as fee_token_hash").
		Where("src_transactions.hash = ?", hash).
		Joins("left join src_transfers on src_transactions.hash = src_transfers.tx_hash").
		Joins("left join poly_transactions on src_transactions.hash = poly_transactions.src_hash").
		Joins("left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash").
		Joins("left join wrapper_transactions on src_transactions.hash = wrapper_transactions.hash").
		Preload("WrapperTransaction").
		Preload("SrcTransaction").
		Preload("SrcTransaction.SrcTransfer").
		Preload("PolyTransaction").
		Preload("DstTransaction").
		Preload("DstTransaction.DstTransfer").
		Preload("Token").
		Preload("Token.TokenBasic").
		Preload("FeeToken").
		Order("src_transactions.time desc").
		Find(&srcPolyDstRelation).Error
	return srcPolyDstRelation, err
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
		srcPolyDstRelation, err := getSrcPolyDstRelation(tx)
		if err != nil {
			logs.Error("getSrcPolyDstRelation of hash: %s err: %s", tx.SrcHash, err)
			continue
		}
		entry := models.ParseBotTx(srcPolyDstRelation, fees)
		if existed, err := cacheRedis.Redis.Exists(cacheRedis.StuckTxAlarmHasSendPrefix + strings.ToLower(entry.Hash)); err == nil && existed {
			logs.Info("stuck TX alarm has been sent: %s", tx.SrcHash)
			continue
		}

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

		baseUrl := conf.GlobalConfig.BotConfig.BaseUrl
		apiToken := conf.GlobalConfig.BotConfig.ApiToken
		btns := []map[string]string{
			map[string]string{
				"title":     "List All",
				"actionURL": baseUrl + conf.GlobalConfig.BotConfig.DetailUrl,
			},
			map[string]string{
				"title":     "Mark As Skipped",
				"actionURL": fmt.Sprintf("%stoken=%s&tx=%s&status=skip", baseUrl+conf.GlobalConfig.BotConfig.FinishUrl, apiToken, entry.Hash),
			},
			map[string]string{
				"title":     "Mark As Waiting",
				"actionURL": fmt.Sprintf("%stoken=%s&tx=%s&status=wait", baseUrl+conf.GlobalConfig.BotConfig.FinishUrl, apiToken, entry.Hash),
			},
			map[string]string{
				"title":     "Mark/Unmark As Paid",
				"actionURL": fmt.Sprintf("%stoken=%s&tx=%s", baseUrl+conf.GlobalConfig.BotConfig.MarkAsPaidUrl, apiToken, entry.Hash),
			},
			map[string]string{
				"title":     "Open",
				"actionURL": baseUrl + conf.GlobalConfig.BotConfig.TxUrl + entry.Hash,
			},
		}

		err = c.PostDingCard(title, body, btns)
		if err != nil {
			logs.Error("send tx stuck ding alarm error. hash: %s, err:", tx.SrcHash, err)
		} else {
			if _, err := cacheRedis.Redis.Set(cacheRedis.StuckTxAlarmHasSendPrefix+strings.ToLower(entry.Hash), "done", time.Hour*24*time.Duration(conf.GlobalConfig.BotConfig.CheckFrom)); err != nil {
				logs.Error("mark tx stuck alarm hash been sent error. hash: %s err: %s", entry.Hash, err)
			}
		}
	}

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

func (c *BotController) ListLargeTxPage() {
	apiToken := c.Ctx.Input.Query("token")
	var err error
	largeTxs := make([]*basedef.LargeTx, 0)
	if apiToken == conf.GlobalConfig.BotConfig.ApiToken {
		ltxs, err := cacheRedis.Redis.LRange(cacheRedis.LargeTxList, -100, -1)
		if err == nil && len(ltxs) != 0 {
			srcPolyDstRelations := make([]*models.SrcPolyDstRelation, 0)
			if err = db.Debug().Table("src_transactions").
				Select("src_transactions.hash as src_hash, src_transactions.chain_id as chain_id").
				Where("src_transactions.hash in ?", ltxs).
				Preload("SrcTransaction").
				Preload("SrcTransaction.SrcTransfer").
				Preload("SrcTransaction.SrcTransfer.Token").
				Preload("SrcTransaction.SrcTransfer.Token.TokenBasic").
				Preload("SrcTransaction.SrcSwap").
				Preload("Token").
				Preload("Token.TokenBasic").
				Order("src_transactions.time desc").
				Find(&srcPolyDstRelations).Error; err != nil {
				logs.Error("query SrcPolyDstRelation err: %s", err)
			} else {
				for _, v := range srcPolyDstRelations {
					srcChainName := strconv.FormatUint(v.ChainId, 10)
					dstChainName := strconv.FormatUint(v.SrcTransaction.DstChainId, 10)
					srcChain := new(models.Chain)
					dstChain := new(models.Chain)
					err = db.Where("chain_id = ?", v.ChainId).First(srcChain).Error
					if err == nil {
						srcChainName = srcChain.Name
					}
					err = db.Where("chain_id = ?", v.SrcTransaction.DstChainId).First(dstChain).Error
					if err == nil {
						dstChainName = dstChain.Name
					}
					if v.SrcTransaction.SrcSwap != nil && v.SrcTransaction.SrcSwap.DstChainId != 0 {
						dstChain := new(models.Chain)
						dstChainName = strconv.FormatUint(v.SrcTransaction.SrcSwap.DstChainId, 10)
						err = db.Where("chain_id = ?", v.SrcTransaction.SrcSwap.DstChainId).First(dstChain).Error
						if err == nil {
							dstChainName = dstChain.Name
						}
					}

					txType := "SWAP"
					if v.SrcTransaction.SrcSwap != nil {
						switch v.SrcTransaction.SrcSwap.Type {
						case basedef.SWAP_SWAP:
							txType = "SWAP"
						case basedef.SWAP_ROLLBACK:
							txType = "ROLLBACK"
						case basedef.SWAP_ADDLIQUIDITY:
							txType = "ADDLIQUIDITY"
						case basedef.SWAP_REMOVELIQUIDITY:
							txType = "REMOVELIQUIDITY"
						}
					}

					var amount, usdAmount decimal.Decimal
					var assetName string
					if v.SrcTransaction.SrcTransfer != nil &&
						v.SrcTransaction.SrcTransfer.Token != nil &&
						v.SrcTransaction.SrcTransfer.Token.TokenBasic != nil {
						assetName = v.SrcTransaction.SrcTransfer.Token.Name
						amount = decimal.NewFromBigInt(&v.SrcTransaction.SrcTransfer.Amount.Int, 0).
							Div(decimal.NewFromInt(basedef.Int64FromFigure(int(v.SrcTransaction.SrcTransfer.Token.Precision))))
						usdAmount = decimal.NewFromBigInt(&v.SrcTransaction.SrcTransfer.Amount.Int, 0).
							Div(decimal.NewFromInt(basedef.Int64FromFigure(int(v.SrcTransaction.SrcTransfer.Token.Precision)))).
							Mul(decimal.NewFromInt(v.SrcTransaction.SrcTransfer.Token.TokenBasic.Price)).
							Div(decimal.NewFromInt(100000000))
					}

					intUsdAmount := usdAmount.IntPart() / 10000

					largeTx := &basedef.LargeTx{
						Asset:     assetName,
						From:      srcChainName,
						To:        dstChainName,
						Type:      txType,
						Amount:    amount.String(),
						USDAmount: strconv.FormatInt(intUsdAmount, 10) + "w",
						Hash:      v.SrcHash,
						User:      v.SrcTransaction.User,
						Time:      time.Unix(int64(v.SrcTransaction.Time), 0).Format("2006-01-02 15:04:05"),
					}
					largeTxs = append(largeTxs, largeTx)
				}
			}
		}

		rows := make([]string, len(largeTxs))
		for i, tx := range largeTxs {
			rows[i] = fmt.Sprintf(
				fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>", 9)),
				tx.Asset,
				tx.Type,
				tx.From,
				tx.To,
				tx.Amount,
				tx.USDAmount,
				tx.Time,
				tx.Hash,
				tx.User,
			)
		}
		rb := []byte(
			fmt.Sprintf(
				`<html><body><h1>Poly large transactions</h1>
					<div>the last %d transactions </div>
						<table style="width:100%%">
						<tr>
							<th>Asset</th>
							<th>Type</th>
							<th>From</th>
							<th>To</th>
							<th>Amount</th>
							<th>USD</th>
							<th>Time</th>
							<th>Hash</th>
							<th>User</th>
						</tr>
						%s
						</table>
				</body></html>`,
				len(largeTxs), strings.Join(rows, "\n"),
			),
		)
		if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
			c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		}
		c.Ctx.Output.Body(rb)
		return
	} else {
		err = fmt.Errorf("access denied")
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
}

func (c *BotController) ListNodeStatusPage() {
	apiToken := c.Ctx.Input.Query("token")
	if apiToken == conf.GlobalConfig.BotConfig.ApiToken {
		nodeStatusesMap := make(map[string][]basedef.NodeStatus, 0)
		chainNames := make([]string, 0)
		for _, cfg := range conf.GlobalConfig.ChainNodes {
			if dataStr, err := cacheRedis.Redis.Get(cacheRedis.NodeStatusPrefix + cfg.ChainName); err == nil {
				var nodeStatuses []basedef.NodeStatus
				if err := json.Unmarshal([]byte(dataStr), &nodeStatuses); err != nil {
					logs.Error("chain %s node status data Unmarshal error: ", cfg.ChainName, err)
					continue
				}
				chainNames = append(chainNames, cfg.ChainName)
				nodeStatusesMap[cfg.ChainName] = nodeStatuses
			}
		}

		sort.Strings(chainNames)
		tables := make([]string, 0)
		for _, chainName := range chainNames {
			nodeStatuses := nodeStatusesMap[chainName]
			rows := make([]string, len(nodeStatuses))
			for i, status := range nodeStatuses {
				rows[i] = fmt.Sprintf(
					fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>\n", 4)),
					status.Url,
					strconv.FormatUint(status.Height, 10),
					status.Status,
					time.Unix(status.Time, 0).Format("2006-01-02 15:04:05"),
				)
			}
			table := fmt.Sprintf(
				`<h2> %s </h2>
					<table style="width:100%%">
						<tr>
							<th>Url</th>
							<th>Height</th>
							<th>Status</th>
							<th>Time</th>
						</tr>
						%s
					</table>`,
				chainName, strings.Join(rows, "\n"))
			tables = append(tables, table)
		}

		htmlBytes := []byte(fmt.Sprintf(`<html><body>
				<h1><center>Chain node status</center></h1>
				%s
				</body></html>`,
			strings.Join(tables, "\n")))
		if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
			c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		}
		c.Ctx.Output.Body(htmlBytes)
		return
	} else {
		err := fmt.Errorf("access denied")
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
}

func (c *BotController) IgnoreNodeStatusAlarm() {
	node := c.Ctx.Input.Query("node")
	day := c.Ctx.Input.Query("day")
	token := c.Ctx.Input.Query("token")
	var err error
	resp := ""
	if token == conf.GlobalConfig.BotConfig.ApiToken {
		dayNum, err := strconv.Atoi(day)
		if err == nil && dayNum >= 0 {
			if dayNum == 0 {
				_, err = cacheRedis.Redis.Del(cacheRedis.IgnoreNodeStatusAlarmPrefix + node)
				if err == nil {
					resp = fmt.Sprintf("success cancel ignore alarm")
				}
			} else {
				_, err := cacheRedis.Redis.Set(cacheRedis.IgnoreNodeStatusAlarmPrefix+node, "ignore", time.Hour*time.Duration(24*dayNum))
				if err == nil {
					resp = fmt.Sprintf("success ignore alarm for %d days", dayNum)
				}
			}
		} else {
			err = fmt.Errorf("invalid parameter day：%s, err: %s", day, err)
		}
	} else {
		err = fmt.Errorf("Access denied")
	}
	if err != nil {
		resp = fmt.Sprintf("Error %s", err.Error())
	}
	logs.Info(resp)
	c.Data["json"] = models.MakeErrorRsp(resp)
	c.ServeJSON()
}

func (c *BotController) ListRelayerAccountStatus() {
	apiToken := c.Ctx.Input.Query("token")
	if apiToken == conf.GlobalConfig.BotConfig.ApiToken {
		accountStatusesMap := make(map[string][]basedef.RelayerAccountStatus, 0)
		chainNames := make([]string, 0)
		for _, cfg := range conf.GlobalConfig.ChainNodes {
			if dataStr, err := cacheRedis.Redis.Get(cacheRedis.RelayerAccountStatusPrefix + cfg.ChainName); err == nil {
				var accountStatuses []basedef.RelayerAccountStatus
				if err := json.Unmarshal([]byte(dataStr), &accountStatuses); err != nil {
					logs.Error("%s relayer account status data Unmarshal error: ", cfg.ChainName, err)
					continue
				}
				chainNames = append(chainNames, cfg.ChainName)
				accountStatusesMap[cfg.ChainName] = accountStatuses
			}
		}
		sort.Strings(chainNames)
		tables := make([]string, 0)
		for _, chainName := range chainNames {
			accountStatuses := accountStatusesMap[chainName]
			rows := make([]string, len(accountStatuses))
			for i, status := range accountStatuses {
				rows[i] = fmt.Sprintf(
					fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>\n", 5)),
					status.Address,
					strconv.FormatFloat(status.Balance, 'f', 6, 64),
					strconv.FormatFloat(status.Threshold, 'f', 6, 64),
					status.Status,
					time.Unix(status.Time, 0).Format("2006-01-02 15:04:05"),
				)
			}
			table := fmt.Sprintf(
				`<h2> %s </h2>
					<table style="width:100%%">
						<tr>
							<th>Address</th>
							<th>Balance</th>
							<th>Threshold</th>
							<th>Status</th>
							<th>Time</th>
						</tr>
						%s
					</table>`,
				chainName, strings.Join(rows, "\n"))
			tables = append(tables, table)
		}

		htmlBytes := []byte(fmt.Sprintf(`<html><body>
				<h1><center>Relayer Account Status</center></h1>
				%s
				</body></html>`,
			strings.Join(tables, "\n")))
		if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
			c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		}
		c.Ctx.Output.Body(htmlBytes)
		return
	} else {
		err := fmt.Errorf("access denied")
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
}
