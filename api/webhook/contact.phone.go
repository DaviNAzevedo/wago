package webhook

type ContactPhone struct {
	Phone *string `json:"phone,omitempty"` // Número de telefone
	Type  *string `json:"type,omitempty"`  // Tipo (HOME, WORK, CELL)
	WaID  *string `json:"wa_id,omitempty"` // ID do WhatsApp associado (se houver)
}
