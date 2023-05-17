package paypal

import (
	"encoding/json"
	"time"
)

type PerchaseUnit struct {
	ReferenceId string
	Shipping    Shipping
	Payments    []Payment
	Amount      Amount
}

type Payment struct {
	Authorizations []Authorization
	Captures       []Capture
}

type Authorization struct {
	Id               string
	Status           string
	Amount           Amount
	SellerProtection SellerProtection
	ExpirationTime   time.Time
	CreateTime       time.Time
	UpdateTime       time.Time
	Links            []Link
}

type Capture struct {
	Id                        string
	Status                    string
	Amount                    Amount
	SellerProtection          SellerProtection
	FinalCapture              bool
	DisbursementMode          string
	SellerReceivableBreakdown SellerReceivableBreakdown
	ExpirationTime            time.Time
	CreateTime                time.Time
	UpdateTime                time.Time
	Links                     []Link
}

type Amount struct {
	CurrencyCode string
	Value        json.Number
	Details      AmountDetail
}

type AmountDetail struct {
	SubTotal json.Number
}

type Link struct {
	Href   string
	Rel    string
	Method string
}

type Payer struct {
	Name         Name
	EmailAddress string
	PayerId      string
}

type Name struct {
	GivenName string
	Surname   string
}

type Shipping struct {
	Address Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	AdminArea1   string
	AdminArea2   string
	PostalCode   string
	CountryCode  string
}

type SellerProtection struct {
	Status            string
	DisputeCategories []string
}

type SellerReceivableBreakdown struct {
	GrossAmount Amount
	PaypalFee   Amount
	NetAmount   Amount
}

type EventType struct {
	Name string
}

type WebhookEvent struct {
	Id           string
	CreateTime   time.Time
	ResourceType string
	EventType    string
	Summary      string
	Resource     Resource
}

type Resource struct {
	Id            string
	CreateTime    time.Time
	UpdateTime    time.Time
	State         string
	Amont         Amount
	ParentPayment string
	ValidUntil    time.Time
	Links         []Link
}
