package ktb

import "github.com/zercle/thai-cross-bank-format/pkg/datamodels"

type Tag30Req struct {
}

func (b Tag30Req) ToTransaction() (result datamodels.Transaction, err error) {
	return
}

type Tag30Resp struct {
}
