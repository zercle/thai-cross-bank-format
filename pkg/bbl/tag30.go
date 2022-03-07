// BBL: Smart Bill Payment
// Payment Notification Service(v0.8)
package bbl

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

var TermType = map[string]string{
	"10": "IVR",
	"20": "KIOSK",
	"30": "ATM",
	"40": "EDC/POS",
	"50": "COUNTER",
	"60": "IBANKING",
	"70": "CDM",
	"80": "MBANKING",
}

var TxnType = map[string]string{
	"C": "payment transaction",
	"D": "payment was cancelled/deleted",
}

var RetryFlag = map[string]string{
	"Y": "retry/resend message",
	"N": "original message",
}

type Tag30Req struct {
	BillerId     string      `json:"billerId"`
	TransDate    string      `json:"transDate"`
	TransTime    string      `json:"transTime"`
	TermType     string      `json:"termType"`
	Amount       json.Number `json:"amount"`
	Reference1   string      `json:"reference1"`
	Reference2   string      `json:"reference2"`
	Reference3   string      `json:"reference3"`
	FromBank     string      `json:"fromBank"`
	FromBranch   string      `json:"fromBranch"`
	FromName     string      `json:"fromName"`
	BankRef      string      `json:"bankRef"`
	ApprovalCode string      `json:"approvalCode"`
	TxnType      string      `json:"txnType"`
	RetryFlag    string      `json:"retryFlag"`
}

func (b Tag30Req) ToTransaction() (result datamodels.Transaction, err error) {
	// convert BBL time format into RFC3339
	transactionTime, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%s+07:00", b.TransDate, b.TransTime))

	if err != nil {
		log.Printf("%+v", err)
		return
	}

	// check amount
	_, err = b.Amount.Float64()
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	result = datamodels.Transaction{
		PayeeBankCode: "002", // BBL
		PayerBankCode: b.FromBank,
		PayerBranchNo: b.FromBranch,
		PayerName:     b.FromName,
		Reference1:    b.Reference1,
		Reference2:    b.Reference2,
		Reference3:    b.Reference3,
		Terminal:      b.TermType,
		Amount:        b.Amount,
		TxRef:         b.BankRef,
		TxDateTime:    transactionTime,
	}

	billerId := []rune(b.BillerId)

	if len(billerId) >= 13 {
		result.PayeePID = string(billerId[:13])
	}

	if termType, ok := TermType[b.TermType]; ok {
		result.Channel = termType
	}

	return
}

type Tag30Resp struct {
	ResponseCode string `json:"responseCode"`
	ResponseMesg string `json:"responseMesg"`
}
