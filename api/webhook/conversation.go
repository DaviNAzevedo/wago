package webhook

type Conversation struct {
	ID                  string  `json:"id"`                             // ID da conversa
	Origin              Origin  `json:"origin"`                         // Origem da conversa
	ExpirationTimestamp *string `json:"expiration_timestamp,omitempty"` // Timestamp Unix de expiração da conversa
}
