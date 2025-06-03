package webhook

type ChangeValue struct {
	MessagingProduct string    `json:"messaging_product"` // Deve ser "whatsapp"
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts,omitempty"` // Informações sobre o usuário que interagiu com a empresa
	Messages         []Message `json:"messages,omitempty"` // Array de mensagens recebidas (normalmente uma por notificação)
	Statuses         []Status  `json:"statuses,omitempty"` // Array de atualizações de status de mensagens
	Errors           []Error   `json:"errors,omitempty"`   // Erros de nível superior relacionados à notificação
}
