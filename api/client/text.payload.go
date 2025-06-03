package client

// TextPayload define o corpo de uma mensagem de texto.
type TextPayload struct {
	PreviewURL bool   `json:"preview_url,omitempty"` // Se true, permite que o WhatsApp gere uma prévia para URLs no texto.
	Body       string `json:"body"`                  // Conteúdo textual da mensagem. UTF-8, até 4096 caracteres.
}
