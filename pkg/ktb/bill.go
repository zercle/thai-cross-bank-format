package ktb

// KTB: Direct Link
// WebService Message Specification Version 5.1
import (
	"encoding/xml"

	"github.com/zercle/thai-cross-bank-format/pkg/datamodels"
)

type BillPaymentReq struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    BillPaymentReqBody
}

type BillPaymentReqBody struct {
	XMLName     xml.Name `xml:"Body"`
	RequestCode BillPaymentReqCode
}

type BillPaymentReqCode struct {
	XMLName  xml.Name `xml:"RequestCode"`
	TranxId  string   `xml:"tranxid,omitempty"`
	Ref1     string   `xml:"ref1"`
	Datetime string   `xml:"datetime"`
	Prodcode string   `xml:"prodcode"`
	Effdate  string   `xml:"effdate"`
	Ref4     string   `xml:"ref4"`
	Bankref  string   `xml:"bankref"`
	Ref3     string   `xml:"ref3"`
	Comcode  string   `xml:"comcode"`
	Command  string   `xml:"command"`
	User     string   `xml:"user"`
	Channel  string   `xml:"channel"`
	Cusname  string   `xml:"cusname,omitempty"`
	Password string   `xml:"password"`
	Ref2     string   `xml:"ref2"`
	Amount   float64  `xml:"amount,omitempty"`
	Bankcode int      `xml:"bankcode"`
}

func (b BillPaymentReqCode) ToTransaction() (result datamodels.Transaction, err error) {
	return
}

type BillPaymentResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    BillPaymentRespBody
}

type BillPaymentRespBody struct {
	XMLName     xml.Name `xml:"Body"`
	RequestCode BillPaymentRespCode
}

type BillPaymentRespCode struct {
	XMLName  xml.Name `xml:"RequestCodeResponse"`
	TranxId  string   `xml:"tranxid"`
	Bankref  string   `xml:"bankref"`
	Respcode string   `xml:"respcode"`
	Respmsg  string   `xml:"respmsg"`
	Balance  string   `xml:"balance,omitempty"`
	Cusname  string   `xml:"cusname,omitempty"`
	Info     string   `xml:"info,omitempty"`
	Print1   string   `xml:"print1,omitempty"`
	Print2   string   `xml:"print2,omitempty"`
	Print3   string   `xml:"print3,omitempty"`
	Print4   string   `xml:"print4,omitempty"`
	Print5   string   `xml:"print5,omitempty"`
	Print6   string   `xml:"print6,omitempty"`
	Print7   string   `xml:"print7,omitempty"`
}
