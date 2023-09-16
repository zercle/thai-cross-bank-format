package bankofthailand

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
)

const (
	RecordHeader string = "H"
	RecordDetail string = "D"
	RecordTotal  string = "T"
	KindDebit    string = "D"
	KindCredit   string = "C"
)

// Bank Of Thailand: cross bank bill payment

type CrossBankBillPayment struct {
	Total   CrossBankBillPaymentTotal    `json:"total"`
	Header  CrossBankBillPaymentHeader   `json:"header"`
	Details []CrossBankBillPaymentDetail `json:"details"`
}

type CrossBankBillPaymentHeader struct {
	RecordType     string      `json:"recordType"`
	SequenceNo     json.Number `json:"sequenceNo"`
	BankCode       string      `json:"bankCode"`
	CompanyAccount string      `json:"companyAccount"`
	CompanyName    string      `json:"companyName"`
	EffectiveDate  time.Time   `json:"effectiveDate"`
	ServiceCode    string      `json:"serviceCode"`
	Spare          string      `json:"spare"`
}

type CrossBankBillPaymentDetail struct {
	RecordType        string      `json:"recordType"`
	SequenceNo        json.Number `json:"sequenceNo"`
	BankCode          string      `json:"bankCode"`
	CompanyAccount    string      `json:"companyAccount"`
	PaymentDate       time.Time   `json:"paymentDate"`
	PaymentTime       time.Time   `json:"paymentTime"`
	CustomerName      string      `json:"customerName"`
	Ref1              string      `json:"reference1"`
	Ref2              string      `json:"reference2"`
	Ref3              string      `json:"reference3"`
	BranchNo          string      `json:"branchNo"`
	TellerNo          string      `json:"tellerNo"`
	KindOfTransaction string      `json:"kindOfTransaction"`
	TransactionCode   string      `json:"transactionCode"`
	ChequeNo          string      `json:"chequeNo"`
	Amount            json.Number `json:"amount"`
	ChequeBankCode    string      `json:"chequeBankCode"`
	Spare1            string      `json:"spare1"`
	BillerID          string      `json:"billerId"`
	Spare2            string      `json:"spare2"`
	SendingBankCode   string      `json:"sendingBankCode"`
	NewChequeNo       string      `json:"newChequeNo"`
}

type CrossBankBillPaymentTotal struct {
	RecordType             string      `json:"recordType"`
	SequenceNo             json.Number `json:"sequenceNo"`
	BankCode               string      `json:"bankCode"`
	CompanyAccount         string      `json:"companyAccount"`
	TotalDebitAmount       json.Number `json:"totalDebitAmount"`
	TotalDebitTransaction  json.Number `json:"totalDebitTransaction"`
	TotalCreditAmount      json.Number `json:"totalCreditAmount"`
	TotalCreditTransaction json.Number `json:"totalCreditTransaction"`
	Spare                  string      `json:"spare"`
}

func ConvertTxtToStruct(source io.Reader) (result CrossBankBillPayment, err error) {
	// set timezone to Thailand
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return
	}
	// Thai bank file encode in TIS-620
	cp874Decoder := charmap.Windows874.NewDecoder()
	reader := cp874Decoder.Reader(source)
	scanner := bufio.NewScanner(reader)
	var line []rune
	var headerLine CrossBankBillPaymentHeader
	var detailLines = make([]CrossBankBillPaymentDetail, 0)
	var totalLine CrossBankBillPaymentTotal
	var detailLine CrossBankBillPaymentDetail
scanLoop:
	for scanner.Scan() {
		line = []rune(scanner.Text())
		// fmt.Printf("type: %s\n", string(line[0:1]))
		// fmt.Printf("seq no: %s\n", string(line[1:7]))
		// fmt.Printf("bank code: %s\n", string(line[7:10]))
		// fmt.Printf("payee acc: %s\n", string(line[10:20]))
		recordType := string(line[0:1])
		seqNo := string(line[1:7])
		seqNo = strings.TrimSpace(seqNo)
		seqNo = strings.TrimLeft(seqNo, "0")
		bankCode := string(line[7:10])
		payeeAcc := string(line[10:20])
		switch recordType {
		case RecordHeader:
			// fmt.Printf("payee name: %s\n", string(line[20:60]))
			// fmt.Printf("data date: %s\n", string(line[60:68]))
			// fmt.Printf("service code: %s\n", string(line[68:76]))
			// fmt.Printf("etc: %s\n", string(line[76:256]))
			payeeName := string(line[20:60])
			dataDateStr := string(line[64:68]) + "-" + string(line[62:64]) + "-" + string(line[60:62])
			dataDate, _ := time.ParseInLocation("2006-01-02", dataDateStr, loc)
			serviceCode := string(line[68:76])
			spare := string(line[76:256])

			headerLine = CrossBankBillPaymentHeader{
				RecordType:     recordType,
				SequenceNo:     json.Number(seqNo),
				BankCode:       bankCode,
				CompanyAccount: payeeAcc,
				CompanyName:    strings.TrimSpace(payeeName),
				EffectiveDate:  dataDate,
				ServiceCode:    strings.TrimSpace(serviceCode),
				Spare:          strings.TrimSpace(spare),
			}
		case RecordDetail:
			// fmt.Printf("payment date: %s\n", string(line[20:28]))
			// fmt.Printf("payment time: %s\n", string(line[28:34]))
			// fmt.Printf("cust name: %s\n", string(line[34:84]))
			// fmt.Printf("ref1: %s\n", string(line[84:104]))
			// fmt.Printf("ref2: %s\n", string(line[104:124]))
			// fmt.Printf("ref3: %s\n", string(line[124:144]))
			// fmt.Printf("branch no: %s\n", string(line[144:148]))
			// fmt.Printf("teller no: %s\n", string(line[148:152]))
			// fmt.Printf("kind of transaction: %s\n", string(line[152:153]))
			// fmt.Printf("transaction code: %s\n", string(line[153:156]))
			// fmt.Printf("cheque no: %s\n", string(line[156:163]))
			// fmt.Printf("amount: %s\n", string(line[163:176]))
			// fmt.Printf("cheque bank code: %s\n", string(line[176:179]))
			// fmt.Printf("etc1: %s\n", string(line[179:196]))
			// fmt.Printf("biller ID: %s\n", string(line[196:211]))
			// fmt.Printf("etc2: %s\n", string(line[211:243]))
			// fmt.Printf("payer bank code: %s\n", string(line[243:246]))
			// fmt.Printf("new cheque no: %s\n", string(line[246:256]))
			// "2006-01-02T15:04:05Z07:00"
			payDateStr := string(line[24:28]) + "-" + string(line[22:24]) + "-" + string(line[20:22])
			payTimeStr := string(line[28:30]) + ":" + string(line[30:32]) + ":" + string(line[32:34])
			payDate, _ := time.ParseInLocation("2006-01-02", payDateStr, loc)
			payTime, _ := time.ParseInLocation("15:04:05", payTimeStr, loc)
			custName := string(line[34:84])
			ref1 := string(line[84:104])
			ref2 := string(line[104:124])
			ref3 := string(line[124:144])
			branchNo := string(line[144:148])
			tellerNo := string(line[148:152])
			kot := string(line[152:153])
			tc := string(line[153:156])
			cn := string(line[156:163])
			amount := string(line[163:174]) + "." + string(line[174:176])
			amount = strings.TrimSpace(amount)
			amount = strings.TrimLeft(amount, "0")
			amount = fmt.Sprintf("%04s", amount)
			cbc := string(line[176:179])
			spare1 := string(line[179:196])
			billerID := string(line[196:211])
			spare2 := string(line[211:243])
			payerBankCode := string(line[243:246])
			ncn := string(line[246:256])

			detailLine = CrossBankBillPaymentDetail{
				RecordType:        recordType,
				SequenceNo:        json.Number(seqNo),
				BankCode:          bankCode,
				CompanyAccount:    payeeAcc,
				PaymentDate:       payDate,
				PaymentTime:       payTime,
				CustomerName:      strings.TrimSpace(custName),
				Ref1:              strings.TrimSpace(ref1),
				Ref2:              strings.TrimSpace(ref2),
				Ref3:              strings.TrimSpace(ref3),
				BranchNo:          strings.TrimSpace(branchNo),
				TellerNo:          strings.TrimSpace(tellerNo),
				KindOfTransaction: strings.TrimSpace(kot),
				TransactionCode:   strings.TrimSpace(tc),
				ChequeNo:          strings.TrimSpace(cn),
				Amount:            json.Number(amount),
				ChequeBankCode:    strings.TrimSpace(cbc),
				Spare1:            strings.TrimSpace(spare1),
				BillerID:          strings.TrimSpace(billerID),
				Spare2:            strings.TrimSpace(spare2),
				SendingBankCode:   strings.TrimSpace(payerBankCode),
				NewChequeNo:       strings.TrimSpace(ncn),
			}
			detailLines = append(detailLines, detailLine)
		case RecordTotal:
			// fmt.Printf("total debit amount: %s\n", string(line[20:33]))
			// fmt.Printf("total debit transaction: %s\n", string(line[33:39]))
			// fmt.Printf("total credit amount: %s\n", string(line[39:52]))
			// fmt.Printf("total credit transaction: %s\n", string(line[52:58]))
			// fmt.Printf("etc: %s\n", string(line[58:256]))

			totalDebitAmount := string(line[20:31]) + "." + string(line[31:33])
			totalDebitAmount = strings.TrimSpace(totalDebitAmount)
			totalDebitAmount = strings.TrimLeft(totalDebitAmount, "0")
			totalDebitAmount = fmt.Sprintf("%04s", totalDebitAmount)
			totalDebitTransaction := string(line[33:39])
			totalDebitTransaction = strings.TrimSpace(totalDebitTransaction)
			totalDebitTransaction = strings.TrimLeft(totalDebitTransaction, "0")
			totalCreditAmount := string(line[39:50]) + "." + string(line[50:52])
			totalCreditAmount = strings.TrimSpace(totalCreditAmount)
			totalCreditAmount = strings.TrimLeft(totalCreditAmount, "0")
			totalCreditAmount = fmt.Sprintf("%04s", totalCreditAmount)
			totalCreditTransaction := string(line[52:58])
			totalCreditTransaction = strings.TrimSpace(totalCreditTransaction)
			totalCreditTransaction = strings.TrimLeft(totalCreditTransaction, "0")
			spare := string(line[58:256])

			totalLine = CrossBankBillPaymentTotal{
				RecordType:             recordType,
				SequenceNo:             json.Number(seqNo),
				BankCode:               bankCode,
				CompanyAccount:         payeeAcc,
				TotalDebitAmount:       json.Number(totalDebitAmount),
				TotalDebitTransaction:  json.Number(totalDebitTransaction),
				TotalCreditAmount:      json.Number(totalCreditAmount),
				TotalCreditTransaction: json.Number(totalCreditTransaction),
				Spare:                  strings.TrimSpace(spare),
			}
			break scanLoop
		}
	}
	result.Header = headerLine
	result.Details = detailLines
	result.Total = totalLine
	return
}
