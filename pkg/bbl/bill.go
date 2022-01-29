package bbl

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-proxy/pkg/datamodels"
)

var RespCode = map[string]string{
	"000": "Success",
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

const (
	HeaderRequestRef = "Request-Ref"
	// yyyy-MM-dd'T'HH:mm:ss.SSS+07:00
	// 2006-01-02T15:04:05.999Z07:00
	HeaderTransmitTime = "Transmit-Date-Time"
)

// BBL: Bill Payment 2Ways
// Payment Notification Service(v5.0)

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

func (b BillPaymentReq) ToTransaction(result datamodels.Transaction, err error) {
	// convert BBL time format into RFC3339
	transactionTime, err := time.Parse(fmt.Sprintf("%sT%s+07:00", b.TransDate, b.TransTime), time.RFC3339)

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

	return
}

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
