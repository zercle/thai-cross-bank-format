package scb

// SCB: QR Payment
// API Specification for Payment Confirmation (v1.0.2.2)
import (
	"log"
	"time"

	"encoding/json"

	"github.com/zercle/thai-cross-bank-format/pkg/bankofthailand"
	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
	"github.com/zercle/thai-cross-bank-format/pkg/utils"
)

// Tag30ProxyType channel that customer pay
var Tag30ProxyType = map[string]string{
	"MSISDN":    "Mobile number",
	"NATID":     "Citizen ID, Tax ID, Government Customer ID",
	"BILLERID":  "Biller ID for Promptpay Bill Pay method",
	"EWALLETID": "E-Wallet ID",
}

// Tag30RespCode response code use in SCB's API
var Tag30RespCode = map[string]string{
	"1000": "Success",
	"1101": "Missing required parameters",
	"1102": "Invalid parameters entered",
	"1103": "Empty string input not supported",
	"1104": "Requested entity record does not exist",
	"1105": "Unrecognized field name was entered - Please check spelling and/or refer to the API docs for correct name",
	"1111": "Data entry duplicated with existing",
	"4101": "Current channel is not supported",
	"8101": "Invalid response from downstream service",
	"8102": "Payment API error code",
	"8901": "Database error",
	"8902": "Error getting mysql_pool connection",
	"9100": "Missing required authorization credentials",
	"9300": "Invalid/expired temporary token",
	"9500": "Invalid authorization credentials",
	"9503": "Invalid access rights",
	"9700": "Generic server side error",
	"9900": "Server is currently unavailable because traffic overload or it is down for maintenance",
}

type Tag30Req struct {
	EventCode              string      `json:"eventCode"`
	TransactionType        string      `json:"transactionType"`
	ReverseFlag            string      `json:"reverseFlag"`
	PayeeProxyId           string      `json:"payeeProxyId"`
	PayeeProxyType         string      `json:"payeeProxyType"`
	PayeeAccountNumber     string      `json:"payeeAccountNumber"`
	PayeeName              string      `json:"payeeName"`
	PayerProxyId           string      `json:"payerProxyId"`
	PayerProxyType         string      `json:"payerProxyType"`
	PayerAccountNumber     string      `json:"payerAccountNumber"`
	PayerName              string      `json:"payerName"`
	SendingBankCode        string      `json:"sendingBankCode"`
	ReceivingBankCode      string      `json:"receivingBankCode"`
	Amount                 json.Number `json:"amount"`
	TransactionId          string      `json:"transactionId"`
	FastEasySlipNumber     string      `json:"fastEasySlipNumber"`
	TransactionDateAndTime string      `json:"transactionDateandTime"`
	BillPaymentRef1        string      `json:"billPaymentRef1"`
	BillPaymentRef2        string      `json:"billPaymentRef2"`
	BillPaymentRef3        string      `json:"billPaymentRef3"`
	CurrencyCode           string      `json:"currencyCode"`
	EquivalentAmount       json.Number `json:"equivalentAmount"`
	EquivalentCurrencyCode string      `json:"equivalentCurrencyCode"`
	ExchangeRate           string      `json:"exchangeRate"`
	ChannelCode            string      `json:"channelCode"`
	PartnerTransactionId   string      `json:"partnerTransactionId"`
	TepaCode               string      `json:"tepaCode"`
}

func (b Tag30Req) ToTransaction() (result datamodels.Transaction, err error) {
	// convert SCB time format into RFC3339
	transactionTime, err := time.ParseInLocation(utils.RFC3339Milli, b.TransactionDateAndTime, time.Local)

	if err != nil {
		log.Printf("%+v", err)
	}

	// check amount
	_, err = b.Amount.Float64()
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	result = datamodels.Transaction{
		PayeeBankCode: bankofthailand.BankCode["SCB"], // SCB
		PayeeAcc:      b.PayeeAccountNumber,
		PayeeName:     b.PayeeName,
		PayeePID:      b.PayeeProxyId,
		PayerBankCode: b.SendingBankCode,
		PayerAcc:      b.PayerAccountNumber,
		PayerName:     b.PayerName,
		PayerPID:      b.PayerProxyId,
		Reference1:    b.BillPaymentRef1,
		Reference2:    b.BillPaymentRef2,
		Reference3:    b.BillPaymentRef3,
		Amount:        b.Amount,
		CurrencyCode:  b.CurrencyCode,
		Channel:       b.ChannelCode,
		TxRef:         b.TransactionId,
		TxDateTime:    transactionTime,
	}

	return
}

// https://developer.scb/#/documents/api-reference-index/references/generic-response-codes.html
type Tag30Resp struct {
	ResCode       string `json:"resCode"`
	ResDesc       string `json:"resDesc"`
	TransactionId string `json:"transactionId"`
	ConfirmId     string `json:"confirmId"`
}
