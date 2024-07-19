package models

type SnsMessageType struct {
	From          string `json:"from"`
	To            string `json:"to"`
	TemplateValue string `json:"templateValue"`
}
