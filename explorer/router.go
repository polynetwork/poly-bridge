package explorer

import (
	"github.com/beego/beego/v2/server/web"
)

func GetRouter() web.LinkNamespace {
	bot := &BotController{}
	go bot.RunChecks()

	ns := web.NSNamespace("/explorer",
		web.NSRouter("/getcrosstx", &ExplorerController{}, "get:GetCrossTx"),
		web.NSRouter("/getassetstatistic", &ExplorerController{}, "get:GetAssetStatistic"),
		web.NSRouter("/gettransferstatistic", &ExplorerController{}, "get:GetTransferStatistic"),
		web.NSRouter("/getexplorerinfo/", &ExplorerController{}, "get:GetExplorerInfo"),
		web.NSRouter("/getcrosstxlist/", &ExplorerController{}, "post:GetCrossTxList"),
		web.NSRouter("/gettokentxlist/", &ExplorerController{}, "post:GetTokenTxList"),
		web.NSRouter("/getaddresstxlist/", &ExplorerController{}, "post:GetAddressTxList"),
		web.NSRouter("/getlocktokenlist/", &ExplorerController{}, "get:GetLockTokenList"),
		web.NSRouter("/getlocktokeninfo/", &ExplorerController{}, "get:GetLockTokenInfo"),
		web.NSRouter("/getetheffectuser/", &ExplorerController{}, "post:GetEthEffectUser"),
		web.NSRouter("/bot/", &BotController{}, "get:BotPage"),
		web.NSRouter("/bottxs/", &BotController{}, "get:GetTxs"),
		web.NSRouter("/botcheck/", &BotController{}, "get:CheckTxs"),
		web.NSRouter("/botcheckfee/", &BotController{}, "post:CheckFees"),
		web.NSRouter("/botfinishtx/", &BotController{}, "get:FinishTx"),
		web.NSRouter("/botmarkunmarktxaspaid/", &BotController{}, "get:MarkUnMarkTxAsPaid"),
		web.NSRouter("/botlistlargetx/", &BotController{}, "get:ListLargeTxPage"),
		web.NSRouter("/botlistnodestatus/", &BotController{}, "get:ListNodeStatusPage"),
		web.NSRouter("/botignorenodestatusalarm/", &BotController{}, "get:IgnoreNodeStatusAlarm"),
		web.NSRouter("/botlistrelayeraccountstatus/", &BotController{}, "get:ListRelayerAccountStatus"),
		web.NSRouter("/getTVLTotal/", &DefiLlamaController{}, "get:GetTVLTotal"),
		web.NSRouter("/getTVLEthereum/", &DefiLlamaController{}, "get:GetTVLEthereum"),
		web.NSRouter("/getTVLOntology/", &DefiLlamaController{}, "get:GetTVLOntology"),
		web.NSRouter("/getTVLNeo/", &DefiLlamaController{}, "get:GetTVLNeo"),
		web.NSRouter("/getTVLCarbon/", &DefiLlamaController{}, "get:GetTVLCarbon"),
		web.NSRouter("/getTVLBNBChain/", &DefiLlamaController{}, "get:GetTVLBNBChain"),
		web.NSRouter("/getTVLHeco/", &DefiLlamaController{}, "get:GetTVLHeco"),
		web.NSRouter("/getTVLOKC/", &DefiLlamaController{}, "get:GetTVLOKC"),
		web.NSRouter("/getTVLNeo3/", &DefiLlamaController{}, "get:GetTVLNeo3"),
		web.NSRouter("/getTVLPolygon/", &DefiLlamaController{}, "get:GetTVLPolygon"),
		web.NSRouter("/getTVLPalette/", &DefiLlamaController{}, "get:GetTVLPalette"),
		web.NSRouter("/getTVLArbitrum/", &DefiLlamaController{}, "get:GetTVLArbitrum"),
		web.NSRouter("/getTVLGnosisChain/", &DefiLlamaController{}, "get:GetTVLGnosisChain"),
		web.NSRouter("/getTVLZilliqa/", &DefiLlamaController{}, "get:GetTVLZilliqa"),
		web.NSRouter("/getTVLAvalanche/", &DefiLlamaController{}, "get:GetTVLAvalanche"),
		web.NSRouter("/getTVLFantom/", &DefiLlamaController{}, "get:GetTVLFantom"),
		web.NSRouter("/getTVLOptimistic/", &DefiLlamaController{}, "get:GetTVLOptimistic"),
		web.NSRouter("/getTVLAndromeda/", &DefiLlamaController{}, "get:GetTVLAndromeda"),
		web.NSRouter("/getTVLBoba/", &DefiLlamaController{}, "get:GetTVLBoba"),
		web.NSRouter("/getTVLOasis/", &DefiLlamaController{}, "get:GetTVLOasis"),
		web.NSRouter("/getTVLHarmony/", &DefiLlamaController{}, "get:GetTVLHarmony"),
		web.NSRouter("/getTVLHSC/", &DefiLlamaController{}, "get:GetTVLHSC"),
		web.NSRouter("/getTVLBytomSidechain/", &DefiLlamaController{}, "get:GetTVLBytomSidechain"),
		web.NSRouter("/getTVLKCC/", &DefiLlamaController{}, "get:GetTVLKCC"),
		web.NSRouter("/getTVLStarcoin/", &DefiLlamaController{}, "get:GetTVLStarcoin"),
		web.NSRouter("/getTVLKava/", &DefiLlamaController{}, "get:GetTVLKava"),
		web.NSRouter("/getTVLCube/", &DefiLlamaController{}, "get:GetTVLCube"),
		web.NSRouter("/getTVLZkSync/", &DefiLlamaController{}, "get:GetTVLZkSync"),
		web.NSRouter("/getTVLCelo/", &DefiLlamaController{}, "get:GetTVLCelo"),
		web.NSRouter("/getTVLClover/", &DefiLlamaController{}, "get:GetTVLClover"),
	)
	return ns
}
