package c2p

import (
	"encoding/json"

	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

type TokenReq struct {
}

type TokenResp struct {
}

type BillPaymentResp struct {
	MerchantID                    string      `json:"merchantID"`
	InvoiceNo                     string      `json:"invoiceNo"`
	Amount                        json.Number `json:"amount"`
	MmcpAmount                    json.Number `json:"mcpAmount"`
	McpFxRate                     json.Number `json:"mcpFxRate"`
	McpCurrencyCode               string      `json:"mcpCurrencyCode"`
	CurrencyCode                  string      `json:"currencyCode"`
	TransactionDateTime           string      `json:"transactionDateTime"`
	AgentCode                     string      `json:"agentCode"`
	ChannelCode                   string      `json:"channelCode"`
	ApprovalCode                  string      `json:"approvalCode"`
	ReferenceNo                   string      `json:"referenceNo"`
	TranRef                       string      `json:"tranRef"`
	CardNo                        string      `json:"cardNo"`
	CardToken                     string      `json:"cardToken"`
	IssuerCountry                 string      `json:"issuerCountry"`
	IssuerBank                    string      `json:"issuerBank"`
	Eci                           string      `json:"eci"`
	InstallmentPeriod             string      `json:"installmentPeriod"`
	InterestType                  string      `json:"interestType"`
	InterestRate                  json.Number `json:"interestRate"`
	InstallmentMerchantAbsorbRate json.Number `json:"installmentMerchantAbsorbRate"`
	RecurringUniqueID             string      `json:"recurringUniqueID"`
	FxAmount                      json.Number `json:"fxAmount"`
	FxRate                        json.Number `json:"fxRate"`
	FxCurrencyCode                string      `json:"fxCurrencyCode"`
	UserDefined1                  string      `json:"userDefined1"`
	UserDefined2                  string      `json:"userDefined2"`
	UserDefined3                  string      `json:"userDefined3"`
	UserDefined4                  string      `json:"userDefined4"`
	UserDefined5                  string      `json:"userDefined5"`
	AcquirerReferenceNo           string      `json:"acquirerReferenceNo"`
	AcquirerMerchantID            string      `json:"acquirerMerchantId"`
	IdempotencyID                 string      `json:"idempotencyID"`
	PaymentScheme                 string      `json:"paymentScheme"`
	IdempotencyNo                 string      `json:"idempotencyNo"`
	RespCode                      string      `json:"respCode"`
	RespDesc                      string      `json:"respDesc"`
}

func (b BillPaymentResp) ToTransaction() (result datamodels.Transaction, err error) {
	return
}
