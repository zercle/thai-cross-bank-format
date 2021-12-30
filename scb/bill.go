package scb

// SCB: Bill Payment Host-to-Host
// interface specification (v1.4)
import (
	"fmt"
	"log"
	"time"

	thaicrossbankformat "github.com/zercle/thai-cross-bank-format"
)

type BillPaymentReq struct {
	Request    string  `json:"request"`
	User       string  `json:"user"`
	Password   string  `json:"password"`
	TranId     string  `json:"tranID"`
	TranDate   string  `json:"tranDate"`
	Channel    string  `json:"channel"`
	Account    string  `json:"account"`
	Amount     float64 `json:"amount"`
	Reference1 string  `json:"reference1"`
	Reference2 string  `json:"reference2"`
	Reference3 string  `json:"reference3"`
	BranchCode string  `json:"branchCode"`
	TerminalId string  `json:"terminalID"`
}

func (b BillPaymentReq) ToCrossBank(result thaicrossbankformat.CrossBankBillPaymentDetail) {
	result = thaicrossbankformat.CrossBankBillPaymentDetail{
		RecordType:        "D",   // Detail
		BankCode:          "014", // SCB
		CompanyAccount:    b.Account,
		Ref1:              b.Reference1,
		Ref2:              b.Reference2,
		Ref3:              b.Reference3,
		BranchNo:          b.BranchCode,
		TellerNo:          b.TerminalId,
		KindOfTransaction: "D", // Debit
		Amount:            b.Amount,
		BankRef:           b.TranId,
	}

	// convert SCB time format into RFC3339
	transactionTime, err := time.Parse(fmt.Sprintf("%s+07:00", b.TranDate), time.RFC3339)

	if err != nil {
		log.Printf("%+v", err)
	}

	result.PaymentDate = transactionTime
	result.PaymentTime = transactionTime

	return
}

type BillPaymentResp struct {
	Response   string  `json:"response"`
	ResCode    string  `json:"resCode"`
	ResMesg    string  `json:"resMesg"`
	TranId     string  `json:"tranID"`
	Reference2 string  `json:"reference2,omitempty"`
	PaymentId  string  `json:"paymentID,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
}
