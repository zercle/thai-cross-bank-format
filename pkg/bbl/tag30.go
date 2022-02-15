package bbl

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

type Tag30Req struct {
	BillerId   string      `json:"payeeId"`
	TransDate  string      `json:"transDate"`
	TransTime  string      `json:"transTime"`
	TermType   string      `json:"termType"`
	Amount     json.Number `json:"amount"`
	Reference1 string      `json:"reference1"`
	Reference2 string      `json:"reference2"`
	Reference3 string      `json:"reference3"`
	FromBank   string      `json:"fromBank"`
	FromBranch string      `json:"fromBranch"`
	FromName   string      `json:"fromName"`
	TxnType    string      `json:"txnType"`
	RetryFlag  string      `json:"retryFlag"`
}

func (b Tag30Req) ToTransaction(result datamodels.Transaction, err error) {
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
		PayerBankCode: b.FromBank,
		PayerBranchNo: b.FromBranch,
		PayerName:     b.FromName,
		Reference1:    b.Reference1,
		Reference2:    b.Reference2,
		Reference3:    b.Reference3,
		Terminal:      b.TermType,
		Amount:        b.Amount,
		TxType:        b.TxnType,
		TxDateTime:    transactionTime,
	}

	billerId := []rune(b.BillerId)

	if len(billerId) >= 13 {
		result.PayeePID = string(billerId[:13])
	}
	return
}

type Tag30Resp struct {
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
