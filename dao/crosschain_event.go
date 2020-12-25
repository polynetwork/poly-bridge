package dao

import "poly-swap/models"

type CrossChainEventDao interface {
	UpdateEvents(chain *models.Chain, wrapperTransactions []*models.WrapperTransaction, srcTransactions []*models.SrcTransaction, polyTransactions []*models.PolyTransaction, dstTransactions []*models.DstTransaction) error
	GetChain(chainId uint64) (*models.Chain, error)
	UpdateChain(chain *models.Chain) error
}
