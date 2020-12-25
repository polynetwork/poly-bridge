package stake_dao

import (
	"encoding/json"
	"fmt"
	"poly-swap/models"
)

type StakeDao struct {
	chains map[uint64]*models.Chain
}

func NewStakeDao() *StakeDao {
	stakeDao := &StakeDao{
	}
	chains := make(map[uint64]*models.Chain)
	chains[2] = &models.Chain{
		ChainId: 2,
		Name:    "Ethereum",
		Height:  9329384,
	}
	stakeDao.chains = chains
	return stakeDao
}

func (dao *StakeDao) UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error {
	{
		json, _ := json.Marshal(chain)
		fmt.Printf("chain: %s\n", json)
	}
	{
		json, _ := json.Marshal(wrapperTransactions)
		fmt.Printf("wrapperTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(srcTransactions)
		fmt.Printf("srcTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(polyTransactions)
		fmt.Printf("polyTransactions: %s\n", json)
	}
	{
		json, _ := json.Marshal(dstTransactions)
		fmt.Printf("dstTransactions: %s\n", json)
	}
	return nil
}

func (dao *StakeDao) GetChain(chainId uint64) (*models.Chain, error) {
	return dao.chains[chainId], nil
}

func (dao *StakeDao) UpdateChain(chain *models.Chain) error {
	dao.chains[chain.ChainId] = chain
	return nil
}
