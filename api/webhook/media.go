package webhook

type Media struct {
	Caption  *string `json:"caption,omitempty"`  // Legenda (para imagem, vídeo, documento)
	Filename *string `json:"filename,omitempty"` // Nome do arquivo (para documento)
	MimeType string  `json:"mime_type"`          // Tipo MIME da mídia
	Sha256   string  `json:"sha256"`             // Hash SHA256 da mídia
	ID       string  `json:"id"`                 // ID da mídia no WhatsApp
	Voice    *bool   `json:"voice,omitempty"`    // Para áudio, indica se é uma mensagem de voz
}
