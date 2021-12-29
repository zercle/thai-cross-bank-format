package scb

// SCB: Bill Payment Host-to-Host
// interface specification (v1.4)
import thaicrossbankformat "github.com/zercle/thai-cross-bank-format"

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
	return
}

type BillPaymentResp struct {
	Response string `json:"response"`
	ResCode  string `json:"resCode"`
	ResMesg  string `json:"resMesg"`
	TranId   string `json:"tranID"`
	// Reference2 string `json:"reference2"`
	// PaymentId  string `json:"paymentID"`
}
