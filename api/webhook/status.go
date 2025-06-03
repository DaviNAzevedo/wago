package webhook

type Status struct {
	ID           string        `json:"id"`                     // ID da mensagem (WAMID) para a qual o status se refere
	RecipientID  string        `json:"recipient_id"`           // ID do WhatsApp do destinatário
	Status       string        `json:"status"`                 // Status: "sent", "delivered", "read", "failed", "warning"
	Timestamp    string        `json:"timestamp"`              // Timestamp Unix da atualização de status
	Conversation *Conversation `json:"conversation,omitempty"` // Informações sobre a conversa (se aplicável)
	Pricing      *Pricing      `json:"pricing,omitempty"`      // Informações de precificação da conversa
	Errors       []Error       `json:"errors,omitempty"`       // Erros (se status="failed")
}
