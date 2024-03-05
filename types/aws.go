package types

type Email struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}
