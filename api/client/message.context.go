package client

// MessageContext é usado para enviar uma resposta a uma mensagem específica.
type MessageContext struct {
	MessageID string `json:"message_id"` // WAMID da mensagem à qual se está respondendo
}
