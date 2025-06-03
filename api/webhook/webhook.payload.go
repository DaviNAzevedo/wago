package webhook

type WebhookPayload struct {
	Object string  `json:"object"` // Deve ser "whatsapp_business_account"
	Entry  []Entry `json:"entry"`
}
