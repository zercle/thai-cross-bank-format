// BBL: Bill Payment 2Ways
// Payment Notification Service(v5.0)
package bbl

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/encoding/json"

	"github.com/zercle/thai-cross-bank-format/pkg/bankofthailand"
	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

// BBL responseCode and responseMesg
var RespCode = map[string]string{
	"000": "Success",
	"052": "Biller ID or Service Code not register",
	"054": "System Unavailable",
	"210": "Time out",
	"211": "Invalid data",
	"212": "Duplicate transaction reference",
	"213": "Invalid amount",
	"214": "Pls. Contact service provider",
	"331": "Invalid reference1/reference2",
	"332": "Invalid payee ID",
	"341": "Service Provider not ready",
	"411": "Reference code expired",
	"412": "Already paid",
	"888": "Other Error",
}

var Channel = map[string]string{
	"A": "ATM",
	"I": "Internet",
	"P": "Phone",
	"C": "Counter Bank",
}

const (
	// BBL Request-Ref
	HeaderRequestRef = "Request-Ref"
	// BBL Transmit-Date-Time format yyyy-MM-dd'T'HH:mm:ss.SSS+07:00
	// eg: 2006-01-02T15:04:05.999Z07:00
	HeaderTransmitTime = "Transmit-Date-Time"
)

// BBL: Bill Payment request from bank
type BillPaymentReq struct {
	PayeeId    string      `json:"payeeId"`
	TransDate  string      `json:"transDate"`
	TransTime  string      `json:"transTime"`
	TransRef   string      `json:"transRef"`
	Channel    string      `json:"channel"`
	TermId     string      `json:"termId"`
	Amount     json.Number `json:"amount"`
	Reference1 string      `json:"reference1"`
	Reference2 string      `json:"reference2"`
	Reference3 string      `json:"reference3"`
	FromBank   string      `json:"fromBank"`
	RetryFlag  string      `json:"retryFlag"`
}

// Transform to transaction
func (b BillPaymentReq) ToTransaction() (result datamodels.Transaction, err error) {
	// convert BBL time format into RFC3339
	transactionTime, err := time.ParseInLocation(time.RFC3339, fmt.Sprintf("%sT%s+07:00", b.TransDate, b.TransTime), time.Local)

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
		PayeeBankCode: bankofthailand.BankCode["BBL"], // BBL
		PayeePID:      b.PayeeId,
		PayerBankCode: b.FromBank,
		Reference1:    b.Reference1,
		Reference2:    b.Reference2,
		Reference3:    b.Reference3,
		Channel:       b.Channel,
		Terminal:      b.TermId,
		Amount:        b.Amount,
		TxRef:         b.TransRef,
		TxDateTime:    transactionTime,
	}

	if channel, ok := Channel[b.Channel]; ok {
		result.Channel = channel
	}

	return
}

// BBL: Bill Payment response to bank
type BillPaymentResp struct {
	ResponseCode  string  `json:"responseCode"`
	ResponseMesg  string  `json:"responseMesg"`
	RspAmount     float64 `json:"rspAmount"`
	RspReference1 string  `json:"rspReference1"`
	RspReference2 string  `json:"rspReference2"`
	RspReference3 string  `json:"rspReference3"`
	UserData1     string  `json:"userData1"`
	UserData2     string  `json:"userData2"`
	UserData3     string  `json:"userData3"`
	UserData4     string  `json:"userData4"`
	UserData5     string  `json:"userData5"`
}
