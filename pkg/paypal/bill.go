package paypal

import "time"

type TokenReq struct {
}

type TokenResp struct {
}

type OrderReq struct {
	Intent        string
	PerchaseUnits []PerchaseUnit
}

type OrderResp struct {
	Id     string
	Status string
	Links  []Link
}

type AuthorizedOrder struct {
	Id            string
	Status        string
	Payer         Payer
	PerchaseUnits []PerchaseUnit
	Links         []Link
}

type CaptureResp struct {
	Id            string
	Status        string
	Payer         Payer
	PerchaseUnits []PerchaseUnit
	Links         []Link
}

type Webhook struct {
	Id              string
	CreateTime      time.Time
	ResourceType    string
	EventVersion    string
	EventType       string
	Summary         string
	ResourceVersion string
	Resource        Resource
	Links           []Link
}

type WebhookReq struct {
	Url        string
	EventTypes []EventType
}

type WebhookResp struct {
	Id         string
	Url        string
	EventTypes []EventType
	Links      []Link
}

type WebhookVerifyReq struct {
	TransmissionId   string
	TransmissionTime time.Time
	CertUrl          string
	AuthAlgo         string
	TransmissionSig  string
	WebhookId        string
	WebhookEvent     WebhookEvent
}

type WebhookVerifyResp struct {
	VerificationStatus string
}

func (c WebhookVerifyResp) Bool() bool {
	return c.VerificationStatus == "SUCCESS"
}
