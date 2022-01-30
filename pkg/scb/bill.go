package scb

// SCB: Bill Payment Host-to-Host
// interface specification (v1.4)
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

// PaymentChannel channel that customer pay
var PaymentChannel = map[string]string{
	"ATM":  "Automatic Teller Machine",
	"CDM":  "Cash Deposit Machine",
	"PHON": "ATM Phone",
	"ENET": "SCB EasyNet",
	"TELE": "SCB EasyPhone (Telebank)",
	"TELL": "Counter",
	"PTNR": "Partners' outlet",
}

// BillPaymentRespCode response code use in SCB's API
var BillPaymentRespCode = map[string]string{
	"0000": "Success",
	"1xxx": "Invalid input data group",
	"1000": "Invalid data",
	"1001": "Invalid reference1",
	"1002": "Invalid reference2",
	"1003": "Invalid reference3",
	"1004": "Invalid amount",
	"2xxx": "Unable to process group",
	"2000": "Unable to process transaction",
	"2001": "Duplicate transaction",
	"2002": "Over due",
	"9xxx": "System error group",
	"9000": "System error",
	"9001": "System is busy",
	"9002": "Time out",
}

type RequestType string

const (
	ReqTypeVerify  RequestType = "verify"
	ReqTypeConfirm RequestType = "confirm"
	ReqTypeCancel  RequestType = "cancel"
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
