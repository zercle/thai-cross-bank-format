package datamodels

import (
	"encoding/json"
	"time"

	both "github.com/zercle/thai-cross-bank-format/pkg/bankofthailand"
)

type Transaction struct {
	PayeePID      string      `json:"payeePID"`      // หมายเลขประจำตัวผู้รับเงิน
	PayeeBankCode string      `json:"payeeBankCode"` // ธนาคารผู้รับเงิน
	PayeeBranchNo string      `json:"payeeBranchNo"` // สาขาผู้รับเงิน
	PayeeAcc      string      `json:"payeeAcc"`      // บัญชีผู้รับเงิน
	PayeeName     string      `json:"payeeName"`     // ชื่อผู้รับเงิน
	PayerPID      string      `json:"payerPID"`      // หมายเลขประจำตัวผู้ชำระเงิน
	PayerBankCode string      `json:"payerBankCode"` // ธนาคารผู้ชำระเงิน
	PayerBranchNo string      `json:"payerBranchNo"` // สาขาผู้ชำระเงิน
	PayerAcc      string      `json:"payerAcc"`      // บัญชีผู้ชำระเงิน
	PayerName     string      `json:"payerName"`     // ชื่อผู้ชำระเงิน
	Reference1    string      `json:"reference1"`    // หมายเลขลูกค้า 1
	Reference2    string      `json:"reference2"`    // หมายเลขลูกค้า 2
	Reference3    string      `json:"reference3"`    // หมายเลขลูกค้า 3
	Amount        json.Number `json:"amount"`        // จำนวนเงิน
	CurrencyCode  string      `json:"currencyCode"`  // รหัสสกุลเงิน
	Channel       string      `json:"channelCode"`   // ช่องทางชำระเงิน
	Terminal      string      `json:"terminalCode"`  // หมายเลขช่องทางชำระเงิน
	KindOfTx      string      `json:"kindOfTx"`      // ประเภทของการชำระเงิน
	TxType        string      `json:"txType"`        // ประเภทของรายการ
	TxRef         string      `json:"txRef"`         // หมายเลขประจำรายการ
	TxDateTime    time.Time   `json:"txDateTime"`    // เวลาทำรายการ
	UserDefined1  string      `json:"userDefined1"`  // อื่น ๆ 1
	UserDefined2  string      `json:"userDefined2"`  // อื่น ๆ 2
	UserDefined3  string      `json:"userDefined3"`  // อื่น ๆ 3
	UserDefined4  string      `json:"userDefined4"`  // อื่น ๆ 4
	UserDefined5  string      `json:"userDefined5"`  // อื่น ๆ 5
}

func (b Transaction) ToCrossBank(result both.CrossBankBillPaymentDetail) {
	result = both.CrossBankBillPaymentDetail{
		RecordType:     both.RecordDetail, // Detail
		BankCode:       b.PayeeBankCode,
		BranchNo:       b.PayeeBranchNo,
		CompanyAccount: b.PayeeAcc,
		// PaymentDate:       b.TxDateTime.Format("2006-01-02"),
		// PaymentTime:       b.TxDateTime.Format("15:04:05"),
		PaymentDate:       time.Date(b.TxDateTime.Year(), b.TxDateTime.Month(), b.TxDateTime.Day(), 0, 0, 0, 0, nil),
		PaymentTime:       time.Date(0, 0, 0, b.TxDateTime.Hour(), b.TxDateTime.Minute(), b.TxDateTime.Second(), 0, nil),
		CustomerName:      b.PayerName,
		Ref1:              b.Reference1,
		Ref2:              b.Reference2,
		Ref3:              b.Reference3,
		TellerNo:          b.Terminal,
		KindOfTransaction: both.KindDebit,
		Amount:            b.Amount,
		SendingBankCode:   b.PayerBankCode,
	}

	if len(b.PayeePID) >= 15 {
		result.BillerID = b.PayeePID
	}
}
