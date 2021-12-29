package thaicrossbankformat

// Bank Of Thailand: cross bank bill payment
import "time"

type CrossBankBillPayment struct {
	Header  CrossBankBillPaymentHeader   `json:"header"`
	Details []CrossBankBillPaymentDetail `json:"details"`
	Total   CrossBankBillPaymentTotal    `json:"total"`
}

type CrossBankBillPaymentHeader struct {
	RecordType     string    `json:"recordType"`
	SequenceNo     int       `json:"sequenceNo"`
	BankCode       string    `json:"bankCode"`
	CompanyAccount string    `json:"companyAccount"`
	CompanyName    string    `json:"companyName"`
	EffectiveDate  time.Time `json:"effectiveDate"`
	ServiceCode    string    `json:"serviceCode"`
	// Spare          string
}

type CrossBankBillPaymentDetail struct {
	RecordType        string    `json:"recordType"`
	SequenceNo        int       `json:"sequenceNo"`
	BankCode          string    `json:"bankCode"`
	CompanyAccount    string    `json:"companyAccount"`
	PaymentDate       time.Time `json:"paymentDate"`
	PaymentTime       time.Time `json:"paymentTime"`
	CustomerName      string    `json:"customerName"`
	Ref1              string    `json:"ref1"`
	Ref2              string    `json:"ref2"`
	Ref3              string    `json:"ref3"`
	BranchNo          string    `json:"branchNo"`
	TellerNo          string    `json:"tellerNo"`
	KindOfTransaction string    `json:"kindOfTransaction"`
	TransactionCode   string    `json:"transactionCode"`
	ChequeNo          string    `json:"chequeNo"`
	Amount            float64   `json:"amount"`
	ChequeBankCode    string    `json:"chequeBankCode"`
	// Spare          string
	BillerId string `json:"billerId"`
	// Spare          string
	SendingBankCode string `json:"sendingBankCode"`
	NewChequeNo     string `json:"newChequeNo"`
}

type CrossBankBillPaymentTotal struct {
	RecordType             string  `json:"recordType"`
	SequenceNo             int     `json:"sequenceNo"`
	BankCode               string  `json:"bankCode"`
	CompanyAccount         string  `json:"companyAccount"`
	TotalDebitAmount       string  `json:"totalDebitAmount"`
	TotalCreditAmount      float64 `json:"totalCreditAmount"`
	TotalCreditTransaction float64 `json:"totalCreditTransaction"`
	// Spare                  string
}
