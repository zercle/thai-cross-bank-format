package bbl

// BBL: Bill Payment 2Ways
// Payment Notification Service(v5.0)
import thaicrossbankformat "github.com/zercle/thai-cross-bank-format"

type BillPaymentReq struct {
	PayeeId    string `json:"payeeId"`
	TransDate  string `json:"transDate"`
	TransTime  string `json:"transTime"`
	TransRef   string `json:"transRef"`
	Channel    string `json:"channel"`
	TermId     string `json:"termId"`
	Amount     string `json:"amount"`
	Reference1 string `json:"reference1"`
	Reference2 string `json:"reference2"`
	Reference3 string `json:"reference3"`
	FromBank   string `json:"fromBank"`
}

func (b BillPaymentReq) ToCrossBank(result thaicrossbankformat.CrossBankBillPaymentDetail) {
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
