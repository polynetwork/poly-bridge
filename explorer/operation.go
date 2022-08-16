package explorer

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"poly-bridge/models"
	"strings"
)

type OperationController struct {
	web.Controller
}

const (
	AIRDROP_AMOUNT_PRECISION = 10000
	AIRDROP                  = "airdrop"
)

func (c *BotController) GetOperationData() {
	token := c.Ctx.Input.Query("token")
	if len(token) == 0 || token != conf.GlobalConfig.OperationConfig.ApiToken {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		return
	}

	method := c.Ctx.Input.Query("method")
	var err error
	var bodyData []byte
	switch method {
	case AIRDROP:
		bodyData, err = c.getAirDropData()
	default:
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("data is null!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		return
	}
	if len(bodyData) == 0 || err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("data is null!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		return
	}
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}
	c.Ctx.Output.Body(bodyData)
	return
}

type methodValue struct {
	Method string
	Value  string
}

func (c *BotController) getAirDropData() ([]byte, error) {
	methodValues := make([]*methodValue, 0)
	var count_all int64
	err := db.Table("(?) as z", db.Model(&models.AirDropInfo{}).Select("bind_addr").Group("bind_addr")).
		Count(&count_all).
		Error
	if err != nil {
		return nil, err
	}
	methodValues = append(methodValues, &methodValue{
		"total", fmt.Sprint(count_all),
	})
	count_gt_1, err := getCountWithAmount(1 * AIRDROP_AMOUNT_PRECISION)
	if err != nil {
		return nil, err
	}
	methodValues = append(methodValues, &methodValue{
		"gt$1", fmt.Sprint(count_gt_1),
	})
	count_gt_10, err := getCountWithAmount(10 * AIRDROP_AMOUNT_PRECISION)
	if err != nil {
		return nil, err
	}
	methodValues = append(methodValues, &methodValue{
		"gt$10", fmt.Sprint(count_gt_10),
	})
	count_gt_100, err := getCountWithAmount(100 * AIRDROP_AMOUNT_PRECISION)
	if err != nil {
		return nil, err
	}
	methodValues = append(methodValues, &methodValue{
		"gt$100", fmt.Sprint(count_gt_100),
	})
	type airDropSumAmount struct {
		BindChainId uint64
		BindAddr    string
		Amount      int64
	}
	airDropSumAmounts := make([]*airDropSumAmount, 0)
	err = db.Table("(?) as z", db.Model(&models.AirDropInfo{}).Select("sum(amount) as amount,bind_addr").Group("bind_addr")).
		Order("z.amount desc").Limit(5).
		Find(&airDropSumAmounts).Error
	if err != nil {
		return nil, err
	}
	bindAddrs := make([]string, 0)
	for _, v := range airDropSumAmounts {
		bindAddrs = append(bindAddrs, v.BindAddr)
	}
	airDropInfos := make([]*models.AirDropInfo, 0)
	err = db.Where("bind_addr in ?", bindAddrs).
		Find(&airDropInfos).Error
	if err != nil {
		return nil, err
	}
	top5 := ""
	for _, airDropSum := range airDropSumAmounts {
		for _, v := range airDropInfos {
			if airDropSum.BindAddr == v.BindAddr {
				airDropSum.BindChainId = v.BindChainId
				airDropSum.BindAddr = basedef.Hash2Address(airDropSum.BindChainId, airDropSum.BindAddr)
				break
			}
		}
		top5 += fmt.Sprint(airDropSum.BindAddr) + "($" + fmt.Sprint(airDropSum.Amount/10000) + ")" + "</br>"
	}
	methodValues = append(methodValues, &methodValue{
		"top5", top5,
	})

	rows := make([]string, 0)
	for _, v := range methodValues {
		rows = append(rows, fmt.Sprintf(
			fmt.Sprintf("<tr>%s</tr>", strings.Repeat("<td align=\"center\">%v</td>", 2)),
			v.Method, v.Value,
		))
	}
	rb := []byte(
		fmt.Sprintf(
			`<html><body><h1>Poly transaction status</h1>
					<div>Air Drop Info</div>
						<table border="3px outset #98bf21" style="width:80%%">
						<tr>
							<th height="50px" width="60px">key</th>
							<th>value</th>
						</tr>
						%s
						</table>
				</body></html>`,
			strings.Join(rows, "\n"),
		),
	)
	return rb, nil
}

func getCountWithAmount(amount int64) (int64, error) {
	var count_greater_amount int64
	err := db.Table("(?) as z", db.Model(&models.AirDropInfo{}).Select("sum(amount) as amount").Group("bind_addr")).
		Where("z.amount > ?", amount).
		Count(&count_greater_amount).
		Error
	return count_greater_amount, err
}
