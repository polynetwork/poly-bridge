package transactions

import (
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/conf"
	"poly-bridge/models"
	"time"
)

func GetStuckTxs(db *gorm.DB, redis *cacheRedis.RedisCache, pageSize, pageNo, from int) ([]*models.TxHashChainIdPair, int, error) {
	tt := time.Now().Unix()
	end := tt - conf.GlobalConfig.EventEffectConfig.HowOld
	if from == 0 {
		from = 3
	}
	endBsc := tt - conf.GlobalConfig.EventEffectConfig.HowOld2

	txs := make([]*models.TxHashChainIdPair, 0)
	var count int64

	var polyProxies []string
	for k := range conf.PolyProxy {
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
			txs[i].DstChainId == basedef.NEO3_CROSSCHAIN_ID ||
			txs[i].SrcChainId == basedef.NEO3N3T5_CROSSCHAIN_ID ||
			txs[i].DstChainId == basedef.NEO3N3T5_CROSSCHAIN_ID) && txs[i].WrapperId == 0 {
			count--
			txs = append(txs[:i], txs[i+1:]...)
			logs.Info("skip %s, because it is a NEO/NEO3 tx with no wrapper_transactions", hash)
			continue
		}

		exists, _ := redis.Exists(cacheRedis.MarkTxAsSkipPrefix + hash)
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

func GetSrcPolyDstRelation(db *gorm.DB, tx *models.TxHashChainIdPair) (*models.SrcPolyDstRelation, error) {
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
