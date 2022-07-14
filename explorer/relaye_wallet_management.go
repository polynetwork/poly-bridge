package explorer

import (
	"encoding/json"
	"fmt"
	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"math"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"sort"
	"strconv"
	"strings"
	"time"
)

type RelayerWalletManagementController struct {
	web.Controller
}

const (
	TransactionDetailsPrefix    = "TransactionDetailsPrefix_"
	AssetCirculateDetailsPrefix = "AssetCirculateDetailsPrefix_"
)

var AssetCirculateDetailDescription = []string{
	"l2 wallet refill relayer account",
	"oneinch swap + cbridge cross-chain + oninch swap",
	"oneinch swap + cbridge cross-chain",
	"cbridge cross-chain + oninch swap",
	"auto extract fee from wrapper and circluate to BSC",
	"l2 wallet refill relayer account with distribute",
	"contract deployment",
	"auto extract fee from wrapper",
	"auto swap asset for avoiding risk",
}
var specialRefillDetailsId = map[uint]bool{
	0: true,
	1: true,
	3: true,
	5: true,
	7: true,
	8: true,
}

var specialTransactionInfoId = map[uint]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
	8: true,
}

var TransactionDetailDescription = []string{
	"basic transfer",
	"batch transfer",
	"cbridge cross chain transaction",
	"oneinch swap transaction",
	"poly crosschain transaction",
	"netswap transaction",
	"extractfee from wrapper",
	"approve",
	"cbridge cross chain target Chain funding transaction",
	"contract deployment",
	"wrapper set fee collector",
}

var testChainId = []uint64{
	17, 22, 6, 21, 24, 23, 19,
}

type RelayerAccountStatus struct {
	ChainId   uint64
	ChainName string
	Address   string
	Balance   float64
	Threshold float64
	Status    string
	Time      int64
}

type AssetCirculateDetail struct {
	FromChainId   uint64
	FromChainName string
	FromAddress   string
	ToChainId     uint64
	ToChainName   string
	ToAddress     string
	FromTokenName string
	FromTokenHash string
	ToTokenName   string
	ToTokenHash   string
	PaymentAmount string
	GasUsed       string
	IncomeAmount  string
	CirculateType uint //0 l2 refill relayer 1:oneinch swap + cbridge cross-chain + oninch swap  2: cbridge cross-chain + oninch swap 3: oneinch swap + cbridge cross-chain 4: auto extract fee from wrapper and circluate to BSC
	Description   string
	Status        string // finished or error message
	Time          int64
}

type TransactionInfo struct {
	TxHash          string
	FromChainId     uint64
	FromChainName   string
	FromAddress     string
	ToChainId       uint64
	ToChainName     string
	ToAddress       string
	FromTokenName   string
	FromTokenHash   string
	ToTokenName     string
	ToTokenHash     string
	PaymentAmount   string
	GasUsed         string
	IncomeAmount    string
	TransactionType uint //0 basic transfer 1 batch transfer 2: cbridge cross chain transaction 3: oneinch swap transaction 4 poly crosschain transaction 5: netswap transaction 6 extractfee from wrapper 7 approve
	Description     string
	Status          string // finished or error message
	Time            int64
}

func GetAssetCirculateDetailSpecificType(chainName string) (details []*AssetCirculateDetail) {
	if exist, _ := cacheRedis.Redis.Exists(AssetCirculateDetailsPrefix + chainName); !exist {
		return nil
	}
	if data, err := cacheRedis.Redis.Smember(AssetCirculateDetailsPrefix + chainName); err == nil {
		for _, dataStr := range data {
			var detail *AssetCirculateDetail
			if err := json.Unmarshal([]byte(dataStr), &detail); err != nil {
				log.Error("%s asset circulate detail data Unmarshal error: ", chainName, err)
				return nil
			}
			if ok, _ := specialRefillDetailsId[detail.CirculateType]; ok {
				details = append(details, detail)
			}
		}
	}
	return
}

func GetTransactionInfoSpecificType(chainName string) (infoArr []*TransactionInfo) {
	if exist, _ := cacheRedis.Redis.Exists(TransactionDetailsPrefix + chainName); !exist {
		return nil
	}
	if data, err := cacheRedis.Redis.Smember(TransactionDetailsPrefix + chainName); err == nil {
		for _, dataStr := range data {
			var info *TransactionInfo
			if err := json.Unmarshal([]byte(dataStr), &info); err != nil {
				log.Error("%s transaction info data Unmarshal error: ", chainName, err)
				return nil
			}
			if ok, _ := specialTransactionInfoId[info.TransactionType]; ok {
				infoArr = append(infoArr, info)
			}
		}
	}
	return
}

func PrintTransactionInfoAndCirculateDetailCirculateOnly(details []*AssetCirculateDetail, infoArr []*TransactionInfo, classifyRes map[int][]int, chainId []uint64) (htmlStr string) {
	chainTables := make([]string, 0)
	reportStr := PrintTransactionReportTableByChain(details, infoArr, chainId)
	chainTables = append(chainTables, reportStr)

	var detail *AssetCirculateDetail
	var info *TransactionInfo
	var detailTotalExpenses, detailTotalIncome, infoTotalExpenses, infoTotalIncome string
	var detailKeys []int
	for k := range classifyRes {
		detailKeys = append(detailKeys, k)
	}
	sort.SliceStable(detailKeys, func(i, j int) bool {
		return detailKeys[i] > detailKeys[j]
	})
	for count, detailKey := range detailKeys {
		//todo title
		detail = details[detailKey]
		detailRows := make([]string, 0)
		detail = details[detailKey]
		txRows := make([]string, 0)
		gasUsed := 0.0
		for _, infoKey := range classifyRes[detailKey] {
			info = infoArr[infoKey]
			if info.GasUsed == "" {
				info.GasUsed = "0"
			}
			gas, err := strconv.ParseFloat(info.GasUsed, 64)
			if err != nil || info.FromChainName != detail.FromChainName {
				log.Error("tx info gasUsed err, %v", "err", err)
			} else {
				gasUsed += gas
			}
			fmt.Println("info", detailKey, infoKey)
			infoTotalExpenses = feeStrPrintHandle(info.PaymentAmount, info.FromTokenName)
			infoTotalIncome = feeStrPrintHandle(info.IncomeAmount, info.ToTokenName)
			txRows = append(txRows, fmt.Sprintf(
				fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>\n", 10)),
				info.TxHash,
				info.FromChainName,
				info.ToChainName,
				info.ToAddress,
				infoTotalExpenses,
				info.GasUsed,
				infoTotalIncome,
				TransactionDetailDescription[info.TransactionType],
				info.Status,
				time.Unix(info.Time, 0).Format("2006-01-02 15:04:05"),
			))
			detailTotalIncome = feeStrPrintHandle(info.IncomeAmount, info.ToTokenName)
		}
		if detail.FromChainId != detail.FromChainId {
			detailTotalExpenses = ""
		} else {
			detailTotalExpenses = feeStrPrintHandle(detail.PaymentAmount, detail.FromTokenName)
		}
		detailRows = append(detailRows, fmt.Sprintf(
			fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>\n", 8)),
			detail.FromAddress,
			detail.ToChainName,
			detailTotalExpenses,
			fmt.Sprintf("%f", gasUsed),
			detailTotalIncome,
			AssetCirculateDetailDescription[detail.CirculateType],
			detail.Status,
			time.Unix(detail.Time, 0).Format("2006-01-02 15:04:05"),
		))

		detailRow := fmt.Sprintf(
			`<tr><td colspan="8">
					<table border="4" style="width:100%%">
						<tr>
	                 <th>Hash</th>
                    <th>From</th>
                    <th>To</th>
      <th>To Address</th>
                    <th>Expenses</th>
                    <th>GasUsed</th>
                    <th>Income</th>
                    <th>Descriotion</th>
                    <th>status</th>
                    <th>Time</th>
						</tr>
						%s
					</table></td><tr>`,
			strings.Join(txRows, "\n"))
		detailRows = append(detailRows, detailRow)
		table := fmt.Sprintf(
			`<h3> %s </h3>
					<table style="width:100%%" border = "2">
						<tr>
            <th>l2 wallet address</th>
            <th>Target Chain</th>
            <th>Total Expenses</th>
       <th>Gas Used</th>
            <th>Total Income</th>
            <th>Descriotion</th>
    <th>Status</th>
            <th>Time</th>
						</tr>
						%s
					</table>`,
			"No."+fmt.Sprintf("%d", count+1)+" l2 wallet transaction record on chain "+detail.FromChainName, strings.Join(detailRows, "\n"))
		chainTables = append(chainTables, table)
	}

	htmlStr = fmt.Sprintf(`<html><body>
				<h1><center>Relayer Wallet Management Transaction Record</center></h1>
				%s
				</body></html>`,
		strings.Join(chainTables, "\n"))
	return
}

func ClassifyTransactionInfoByCirculateDetailByChain(details []*AssetCirculateDetail, infoArr []*TransactionInfo, chainIds []uint64) (classifyRes map[uint64]map[int][]int) {
	classifyRes = make(map[uint64]map[int][]int)
	var timeNow, timeNext int64
	var cur = 0
	sort.SliceStable(details, func(i, j int) bool {
		return details[i].Time < details[j].Time
	})
	sort.SliceStable(infoArr, func(i, j int) bool {
		return infoArr[i].Time < infoArr[j].Time
	})

	for _, chainId := range chainIds {
		classifyRes[chainId] = make(map[int][]int)
	}
	for id, detail := range details {
		if detail == nil {
			continue
		}
		timeNow = detail.Time
		if id < len(details)-1 {
			timeNext = details[id+1].Time
		} else {
			timeNext = math.MaxInt64
		}
		if infoArr[cur] == nil || infoArr[cur].Time > timeNext {
			continue
		}
		var infoTmp, infoSpecial []int
		var cbridgeFundingChainId uint64
		for i := cur; i < len(infoArr); i++ {
			if infoArr[i] == nil {
				continue
			}
			if infoArr[i].Time >= timeNow && infoArr[i].Time < timeNext {
				// for the cbridge crosschain from erc20 => erc20 on target chain
				if infoArr[i].FromChainId != detail.FromChainId {
					infoSpecial = append(infoSpecial, i)
					cbridgeFundingChainId = infoArr[i].FromChainId
				} else {
					infoTmp = append(infoTmp, i)
				}
				cur = i + 1
			}
		}
		if infoSpecial != nil {
			classifyRes[cbridgeFundingChainId][id] = infoSpecial
		}
		if infoTmp != nil {
			classifyRes[detail.FromChainId][id] = infoTmp
		}
	}
	return classifyRes
}

func PrintTransactionReportTableByChain(details []*AssetCirculateDetail, infoArr []*TransactionInfo, chainIds []uint64) (htmlStr string) {
	classifyRes := ClassifyTransactionInfoByCirculateDetailByChain(details, infoArr, chainIds)
	var detail *AssetCirculateDetail
	var info *TransactionInfo
	detailRows := make([]string, 0)
	var totalIncomeFromWrapper, totalIncomeFromOtherChains, totalExpensesSupportingOtherChains, totalAmountRelayers, totalGasConsumption, totalCbridgeConsumption float64
	var relayerRefillCount int = 0
	for _, chainId := range chainIds {
		relayerRefillCount = 0
		totalGasConsumption = 0.0
		totalIncomeFromWrapper = 0.0
		totalIncomeFromOtherChains = 0.0
		totalExpensesSupportingOtherChains = 0.0
		totalAmountRelayers = 0.0
		totalCbridgeConsumption = 0.0
		m := classifyRes[chainId]
		if len(m) == 0 {
			continue
		}
		var detailKeys []int
		for k := range m {
			detailKeys = append(detailKeys, k)
		}
		sort.SliceStable(detailKeys, func(i, j int) bool {
			return detailKeys[i] > detailKeys[j]
		})
		for _, detailKey := range detailKeys {
			detail = details[detailKey]
			for _, infoKey := range m[detailKey] {
				info = infoArr[infoKey]
				totalGasConsumption = addStrToFloat(totalGasConsumption, info.GasUsed)
				if info.TransactionType == 6 {
					totalIncomeFromWrapper = addStrToFloat(totalIncomeFromWrapper, info.IncomeAmount)
				}
				if info.TransactionType == 8 {
					totalIncomeFromOtherChains = addStrToFloat(totalIncomeFromOtherChains, info.IncomeAmount)
				}
				if info.TransactionType == 2 {
					totalCbridgeConsumption = addTwoStrSubtractToFloat(totalCbridgeConsumption, info.PaymentAmount, infoArr[infoKey+1].IncomeAmount)
				}
			}
			if detail.CirculateType == 0 || detail.CirculateType == 5 {
				relayerRefillCount++
				totalAmountRelayers = addStrToFloat(totalAmountRelayers, detail.PaymentAmount)
			}
			if detail.CirculateType == 1 && detail.FromChainId == chainId {
				totalExpensesSupportingOtherChains = addStrToFloat(totalExpensesSupportingOtherChains, detail.PaymentAmount)
			}
		}
		detailRows = append(detailRows, fmt.Sprintf(
			fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td>%s</td>\n", 9)),
			details[detailKeys[0]].FromChainName,
			details[detailKeys[0]].FromAddress,
			fmt.Sprintf("%f", totalIncomeFromWrapper)+" "+details[detailKeys[0]].FromTokenName,
			fmt.Sprintf("%f", totalIncomeFromOtherChains)+" USDT",
			fmt.Sprintf("%f", totalAmountRelayers)+" "+details[detailKeys[0]].FromTokenName,
			fmt.Sprintf("%d", relayerRefillCount)+" times",
			fmt.Sprintf("%f", totalGasConsumption)+" "+details[detailKeys[0]].FromTokenName,
			fmt.Sprintf("%f", totalCbridgeConsumption)+" USDT",
			fmt.Sprintf("%f", totalExpensesSupportingOtherChains)+" "+details[detailKeys[0]].FromTokenName,
		))
	}
	return fmt.Sprintf(
		`<h3> %s </h3>
					<table style="width:100%%" border = "2">
						<tr>
            <th>Chains</th>
                   <th>l2 wallet address</th>
            <th>Total Income From Wrapper</th>
     <th>Total Income From Other Chains</th>
            <th>Total Amount For Relayers</th>
    <th>Total Times Refill Relayers</th>
       <th>Total Gas Consumption</th>
    <th>Total Cbridge Consumption</th>
         <th>Total Amount For Helping Other Chains</th>
						</tr>
						%s
					</table>`,
		"Transaction Report", strings.Join(detailRows, "\n"))
}

func ClassifyTransactionInfoByCirculateDetailByRefillTurn(details []*AssetCirculateDetail, infoArr []*TransactionInfo) (classifyRes map[int][]int) {
	classifyRes = make(map[int][]int)
	var timeNow, timeNext int64
	var cur = 0
	sort.SliceStable(details, func(i, j int) bool {
		return details[i].Time < details[j].Time
	})
	sort.SliceStable(infoArr, func(i, j int) bool {
		return infoArr[i].Time < infoArr[j].Time
	})
	for id, detail := range details {
		if detail == nil {
			continue
		}
		timeNow = detail.Time
		if id < len(details)-1 {
			timeNext = details[id+1].Time
		} else {
			timeNext = math.MaxInt64
		}
		if infoArr[cur] == nil || infoArr[cur].Time > timeNext {
			continue
		}
		var infoTmp []int
		for i := cur; i < len(infoArr); i++ {
			if infoArr[i] == nil {
				continue
			}
			if infoArr[i].Time >= timeNow && infoArr[i].Time < timeNext {
				infoTmp = append(infoTmp, i)
				cur = i + 1
			}
		}
		if infoTmp != nil {
			classifyRes[id] = infoTmp
		}
	}
	return classifyRes
}

func (c *RelayerWalletManagementController) ListRelayerRefillTransactionRecord() {
	apiToken := c.Ctx.Input.Query("token")
	if apiToken == conf.GlobalConfig.BotConfig.ApiToken {
		var details []*AssetCirculateDetail
		var infoArr []*TransactionInfo
		for _, nodes := range conf.GlobalConfig.ChainNodes {
			details = append(details, GetAssetCirculateDetailSpecificType(basedef.GetChainName(nodes.ChainId))...)
			infoArr = append(infoArr, GetTransactionInfoSpecificType(basedef.GetChainName(nodes.ChainId))...)
		}
		classifyRes := ClassifyTransactionInfoByCirculateDetailByRefillTurn(details, infoArr)
		htmlBytes := []byte(PrintTransactionInfoAndCirculateDetailCirculateOnly(details, infoArr, classifyRes, testChainId))
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

func feeStrPrintHandle(amt, tokenName string) string {
	if amt == "" {
		return ""
	} else {
		return amt + " " + tokenName
	}
}

func addStrToFloat(f float64, str string) float64 {
	tmp, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Error("tx info f err, %v", "err", err)
		return f
	} else {
		return f + tmp
	}
}

func addTwoStrSubtractToFloat(f float64, str1, str2 string) float64 {
	tmp1, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		log.Error("str1 info f err, %v", "err", err)
		return f
	}
	tmp2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		log.Error("str2 info f err, %v", "err", err)
		return f
	}
	return f + tmp1 - tmp2
}
