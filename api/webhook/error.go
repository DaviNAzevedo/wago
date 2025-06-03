package webhook

type Error struct {
	Code      int        `json:"code"`                 // Código de erro do WhatsApp
	Title     string     `json:"title"`                // Título do erro
	Message   *string    `json:"message,omitempty"`    // Mensagem de erro legível (nem sempre presente)
	ErrorData *ErrorData `json:"error_data,omitempty"` // Dados adicionais do erro
	Details   *string    `json:"details,omitempty"`    // Detalhes do erro (às vezes presente em vez de ErrorData)
}
