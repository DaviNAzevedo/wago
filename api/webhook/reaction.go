package webhook

type Reaction struct {
	MessageID string  `json:"message_id"`      // ID da mensagem que recebeu a reação
	Emoji     *string `json:"emoji,omitempty"` // Emoji da reação (vazio se a reação foi removida)
}
