package scb

// SCB: QR Payment
// API Specification for Payment Confirmation (v1.0.2.2)
import (
	"log"
	"time"

	thaicrossbankformat "github.com/zercle/thai-cross-bank-format"
)

type Tag30Req struct {
	PayeeProxyId           string  `json:"payeeProxyId"`
	PayeeProxyType         string  `json:"payeeProxyType"`
	PayeeAccountNumber     string  `json:"payeeAccountNumber"`
	PayerProxyId           string  `json:"payerProxyId"`
	PayerProxyType         string  `json:"payerProxyType"`
	PayerAccountNumber     string  `json:"payerAccountNumber"`
	SendingBankCode        string  `json:"sendingBankCode"`
	ReceivingBankCode      string  `json:"receivingBankCode"`
	Amount                 float64 `json:"amount"`
	TransactionId          string  `json:"transactionId"`
	FastEasySlipNumber     string  `json:"fastEasySlipNumber"`
	TransactionDateAndTime string  `json:"transactionDateandTime"`
	BillPaymentRef1        string  `json:"billPaymentRef1"`
	BillPaymentRef2        string  `json:"billPaymentRef2"`
	BillPaymentRef3        string  `json:"billPaymentRef3"`
	ThaiQRTag              string  `json:"thaiQRTag"`
	MerchantId             string  `json:"merchantId"`
	MerchantPAN            string  `json:"merchantPAN"`
	ConsumerPAN            string  `json:"consumerPAN"`
	CurrencyCode           string  `json:"currencyCode"`
}

func (b Tag30Req) ToCrossBank(result thaicrossbankformat.CrossBankBillPaymentDetail) {
	result = thaicrossbankformat.CrossBankBillPaymentDetail{
		RecordType:        "D",   // Detail
		BankCode:          "014", // SCB
		CompanyAccount:    b.PayeeAccountNumber,
		Ref1:              b.BillPaymentRef1,
		Ref2:              b.BillPaymentRef2,
		Ref3:              b.BillPaymentRef3,
		KindOfTransaction: "D", // Debit
		Amount:            b.Amount,
		BillerId:          b.PayeeProxyId,
		SendingBankCode:   b.SendingBankCode,
		BankRef:           b.TransactionId,
	}

	// convert SCB time format into RFC3339
	transactionTime, err := time.Parse(b.TransactionDateAndTime, thaicrossbankformat.RFC3339Mili)

	if err != nil {
		log.Printf("%+v", err)
	}

	result.PaymentDate = transactionTime
	result.PaymentTime = transactionTime
	return
}

type Tag30Resp struct {
	ResCode       string `json:"resCode"`
	ResDesc       string `json:"resDesc"`
	TransactionId string `json:"transactionId"`
	ConfirmId     string `json:"confirmId"`
}
