package client

// DocumentPayload é usado para enviar mensagens de documento.
// Você pode enviar por ID (se já foi carregado) ou por URL.
type DocumentPayload struct {
	ID       string  `json:"id,omitempty"`       // ID do documento previamente carregado.
	Link     string  `json:"link,omitempty"`     // URL do documento. HTTP ou HTTPS.
	Caption  *string `json:"caption,omitempty"`  // Legenda para o documento. Até 1024 caracteres.
	Filename *string `json:"filename,omitempty"` // Nome do arquivo como aparecerá para o usuário.
}
