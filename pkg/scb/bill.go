package scb

// SCB: Bill Payment Host-to-Host
// interface specification (v1.4)
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-proxy/pkg/datamodels"
)

type BillPaymentReq struct {
	Request    string      `json:"request"`
	User       string      `json:"user"`
	Password   string      `json:"password"`
	TranId     string      `json:"tranID"`
	TranDate   string      `json:"tranDate"`
	Channel    string      `json:"channel"`
	Account    string      `json:"account"`
	Amount     json.Number `json:"amount"`
	Reference1 string      `json:"reference1"`
	Reference2 string      `json:"reference2"`
	Reference3 string      `json:"reference3"`
	BranchCode string      `json:"branchCode"`
	TerminalId string      `json:"terminalID"`
}

func (b BillPaymentReq) ToTransaction(result datamodels.Transaction, err error) {
	// convert SCB time format into RFC3339
	transactionTime, err := time.Parse(fmt.Sprintf("%s+07:00", b.TranDate), time.RFC3339)

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
		PayeeBankCode: "014", // SCB
		PayeeAcc:      b.Account,
		Reference1:    b.Reference1,
		Reference2:    b.Reference2,
		Reference3:    b.Reference3,
		PayeeBranchNo: b.BranchCode,
		Terminal:      b.TerminalId,
		Amount:        b.Amount,
		TxRef:         b.TranId,
		TxDateTime:    transactionTime,
	}

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
