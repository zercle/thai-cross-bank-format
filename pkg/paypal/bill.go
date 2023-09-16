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
	ID     string
	Status string
	Links  []Link
}

type AuthorizedOrder struct {
	ID            string
	Status        string
	Payer         Payer
	PerchaseUnits []PerchaseUnit
	Links         []Link
}

type CaptureResp struct {
	ID            string
	Status        string
	Payer         Payer
	PerchaseUnits []PerchaseUnit
	Links         []Link
}

type Webhook struct {
	ID              string
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
	URL        string
	EventTypes []EventType
}

type WebhookResp struct {
	ID         string
	URL        string
	EventTypes []EventType
	Links      []Link
}

type WebhookVerifyReq struct {
	TransmissionID   string
	TransmissionTime time.Time
	CertURL          string
	AuthAlgo         string
	TransmissionSig  string
	WebhookID        string
	WebhookEvent     WebhookEvent
}

type WebhookVerifyResp struct {
	VerificationStatus string
}

func (c WebhookVerifyResp) Bool() bool {
	return c.VerificationStatus == "SUCCESS"
}
