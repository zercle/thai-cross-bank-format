package bbl

import thaicrossbankformat "github.com/zercle/thai-cross-bank-format"

type Tag30Req struct {
}

func (b Tag30Req) ToCrossBank(result thaicrossbankformat.CrossBankBillPaymentDetail) {
	return
}

type Tag30Resp struct {
}
