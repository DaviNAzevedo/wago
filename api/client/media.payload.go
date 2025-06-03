package client

// MediaPayload é usado para enviar mensagens de mídia (imagem, áudio, vídeo, sticker).
// Você pode enviar mídia por ID (se já foi carregada) ou por URL.
type MediaPayload struct {
	ID      string  `json:"id,omitempty"`      // ID da mídia previamente carregada.
	Link    string  `json:"link,omitempty"`    // URL da mídia. HTTP ou HTTPS. O WhatsApp fará o download.
	Caption *string `json:"caption,omitempty"` // Legenda para a mídia (apenas para imagem, vídeo). Até 1024 caracteres.
}
