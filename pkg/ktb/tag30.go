package ktb

import (
	bot "github.com/zercle/thai-cross-bank-proxy/pkg/bankofthailand"
)

type Tag30Req struct {
}

func (b Tag30Req) ToCrossBank(result bot.CrossBankBillPaymentDetail) {
	return
}

type Tag30Resp struct {
}
