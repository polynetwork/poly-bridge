package explorer

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"math/big"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/models"
	"poly-bridge/utils/decimal"
)

type DefiLlamaController struct {
	web.Controller
}

func getTVLAmount(chain uint64) (amount string, err error) {
	amount, err = cacheRedis.Redis.GetChainTvl(chain)
	if err == nil {
		logs.Info("getTVLAmount chain with Redis,chain:", chain, "amount:", amount)
		return
	}
	lockTokens := make([]*models.LockTokenStatistic, 0)
	db.Where("chain_id = ?", chain).
		Find(&lockTokens)
	tVLChain := big.NewInt(0)
	for _, lockToken := range lockTokens {
		tVLChain.Add(tVLChain, &lockToken.InAmountUsd.Int)
	}
	amount = decimal.NewFromBigInt(tVLChain, -4).StringFixed(2)
	err = cacheRedis.Redis.SetChainTvl(chain, amount)
	if err != nil {
		logs.Error("getTVLAmount SetChainTvl err,chain:", chain, err)
	}
	return amount, nil
}

func getTVLtotalAmount() (amount string, err error) {
	totalChain := uint64(0)
	amount, err = cacheRedis.Redis.GetChainTvl(totalChain)
	if err == nil {
		logs.Info("getTVLtotalAmount chain with Redis,chain:", totalChain, "amount:", amount)
		return
	}
	lockTokens := make([]*models.LockTokenStatistic, 0)
	db.Where("in_amount_usd <> '0'").
		Find(&lockTokens)
	tVLChain := big.NewInt(0)
	for _, lockToken := range lockTokens {
		if lockToken.InAmountUsd.Cmp(big.NewInt(0)) > 0 {
			tVLChain.Add(tVLChain, &lockToken.InAmountUsd.Int)
		}
	}
	amount = decimal.NewFromBigInt(tVLChain, -4).StringFixed(2)
	err = cacheRedis.Redis.SetChainTvl(totalChain, amount)
	if err != nil {
		logs.Error("getTVLtotalAmount SetChainTvl err,chain:", totalChain, err)
	}
	return amount, nil
}

func (c *DefiLlamaController) GetTVLTotal() {
	tvlAmount, err := getTVLtotalAmount()
	if err != nil {
		logs.Error("GetTVLTotal err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLEthereum() {
	tvlAmount, err := getTVLAmount(basedef.ETHEREUM_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLEthereum err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLOntology() {
	tvlAmount, err := getTVLAmount(basedef.ONT_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLOntology err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLNeo() {
	tvlAmount, err := getTVLAmount(basedef.NEO_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLNeo err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLCarbon() {
	tvlAmount, err := getTVLAmount(basedef.SWITCHEO_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLCarbon err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLBNBChain() {
	tvlAmount, err := getTVLAmount(basedef.BSC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLBNBChain err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLHeco() {
	tvlAmount, err := getTVLAmount(basedef.HECO_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLHeco err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLOKC() {
	tvlAmount, err := getTVLAmount(basedef.OK_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLOKC err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLNeo3() {
	tvlAmount, err := getTVLAmount(basedef.NEO3_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLNeo3 err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLPolygon() {
	tvlAmount, err := getTVLAmount(basedef.MATIC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLPolygon err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLPalette() {
	tvlAmount, err := getTVLAmount(basedef.PLT_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLPalette err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLArbitrum() {
	tvlAmount, err := getTVLAmount(basedef.ARBITRUM_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLArbitrum err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLGnosisChain() {
	tvlAmount, err := getTVLAmount(basedef.XDAI_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLGnosisChain err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLZilliqa() {
	tvlAmount, err := getTVLAmount(basedef.ZILLIQA_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLZilliqa err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLAvalanche() {
	tvlAmount, err := getTVLAmount(basedef.AVAX_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLAvalanche err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLFantom() {
	tvlAmount, err := getTVLAmount(basedef.FANTOM_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLFantom err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLOptimistic() {
	tvlAmount, err := getTVLAmount(basedef.OPTIMISTIC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLOptimistic err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLAndromeda() {
	tvlAmount, err := getTVLAmount(basedef.METIS_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLAndromeda err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLBoba() {
	tvlAmount, err := getTVLAmount(basedef.BOBA_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLBoba err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLOasis() {
	tvlAmount, err := getTVLAmount(basedef.OASIS_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLOasis err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLHarmony() {
	tvlAmount, err := getTVLAmount(basedef.HARMONY_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLHarmony err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLHSC() {
	tvlAmount, err := getTVLAmount(basedef.HSC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLHSC err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLBytomSidechain() {
	tvlAmount, err := getTVLAmount(basedef.BYTOM_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLBytomSidechain err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
func (c *DefiLlamaController) GetTVLKCC() {
	tvlAmount, err := getTVLAmount(basedef.KCC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLKCC err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLStarcoin() {
	tvlAmount, err := getTVLAmount(basedef.STARCOIN_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLStarcoin err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLKava() {
	tvlAmount, err := getTVLAmount(basedef.KAVA_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLKava err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLCube() {
	tvlAmount, err := getTVLAmount(basedef.CUBE_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLCube err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLZkSync() {
	tvlAmount, err := getTVLAmount(basedef.ZKSYNC_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLCube err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLCelo() {
	tvlAmount, err := getTVLAmount(basedef.CELO_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLCube err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}

func (c *DefiLlamaController) GetTVLClover() {
	tvlAmount, err := getTVLAmount(basedef.CLOVER_CROSSCHAIN_ID)
	if err != nil {
		logs.Error("GetTVLCube err", err)
	}
	c.Data["json"] = tvlAmount
	c.ServeJSON()
}
