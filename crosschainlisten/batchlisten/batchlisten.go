package batchlisten

import "poly-bridge/models"

type BatchListen struct {
	chWrapperTx         chan *models.WrapperTransaction
	chSrcTx             chan *models.SrcTransaction
	chDstTx             chan *models.DstTransaction
	WrapperTransactions []*models.WrapperTransaction
	SrcTransactions     []*models.SrcTransaction
	DstTransactions     []*models.DstTransaction
}

func NewBatchListen(size int, done func()) *BatchListen {
	b := &BatchListen{
		chWrapperTx:         make(chan *models.WrapperTransaction, size),
		chSrcTx:             make(chan *models.SrcTransaction, size),
		chDstTx:             make(chan *models.DstTransaction, size),
		WrapperTransactions: make([]*models.WrapperTransaction, 0),
		SrcTransactions:     make([]*models.SrcTransaction, 0),
		DstTransactions:     make([]*models.DstTransaction, 0),
	}
	go func() {
		b.Schedule()
		done()
	}()
	return b
}

func (b *BatchListen) Schedule() {
	// 从 channel 接收数据
	for i := range b.chWrapperTx {
		b.WrapperTransactions = append(b.WrapperTransactions, i)
	}
	for i := range b.chSrcTx {
		b.SrcTransactions = append(b.SrcTransactions, i)
	}
	for i := range b.chDstTx {
		b.DstTransactions = append(b.DstTransactions, i)
	}
}

func (b *BatchListen) Close() {
	close(b.chWrapperTx)
	close(b.chSrcTx)
	close(b.chDstTx)
}

func (b *BatchListen) AddWrapperTx(v *models.WrapperTransaction) {
	b.chWrapperTx <- v
}
func (b *BatchListen) AddSrcTx(v *models.SrcTransaction) {
	b.chSrcTx <- v
}
func (b *BatchListen) AddDstTx(v *models.DstTransaction) {
	b.chDstTx <- v
}
