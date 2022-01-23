package datamodels

import (
	"encoding/json"
	"time"

	bot "github.com/zercle/thai-cross-bank-proxy/pkg/bankofthailand"
)

type Transaction struct {
	PayeePID      string      `json:"payeePID"`
	PayeeBankCode string      `json:"payeeBankCode"`
	PayeeBranchNo string      `json:"payeeBranchNo"`
	PayeeAcc      string      `json:"payeeAcc"`
	PayeeName     string      `json:"payeeName"`
	PayerPID      string      `json:"payerPID"`
	PayerBankCode string      `json:"payerBankCode"`
	PayerBranchNo string      `json:"payerBranchNo"`
	PayerAcc      string      `json:"payerAcc"`
	PayerName     string      `json:"payerName"`
	Reference1    string      `json:"reference1"`
	Reference2    string      `json:"reference2"`
	Reference3    string      `json:"reference3"`
	Amount        json.Number `json:"amount"`
	CurrencyCode  string      `json:"currencyCode"`
	Channel       string      `json:"channelCode"`
	Terminal      string      `json:"terminalCode"`
	TxRef         string      `json:"txRef"`
	TxDateTime    time.Time   `json:"txDateTime"`
	UserDefined1  string      `json:"userDefined1"`
	UserDefined2  string      `json:"userDefined2"`
	UserDefined3  string      `json:"userDefined3"`
	UserDefined4  string      `json:"userDefined4"`
	UserDefined5  string      `json:"userDefined5"`
}

func (b Transaction) ToCrossBank(result bot.CrossBankBillPaymentDetail) {
	result = bot.CrossBankBillPaymentDetail{
		RecordType:        "D", // Detail
		BankCode:          b.PayeeBankCode,
		BranchNo:          b.PayeeBranchNo,
		CompanyAccount:    b.PayeeAcc,
		PaymentDate:       b.TxDateTime.Format("2006-01-02"),
		PaymentTime:       b.TxDateTime.Format("15:04:05"),
		CustomerName:      b.PayerName,
		Ref1:              b.Reference1,
		Ref2:              b.Reference2,
		Ref3:              b.Reference3,
		TellerNo:          b.Terminal,
		KindOfTransaction: "D", // Debit
		Amount:            b.Amount,
		SendingBankCode:   b.PayerBankCode,
	}

	if len(b.PayeePID) >= 15 {
		result.BillerId = b.PayeePID
	}
	
	return
}
