package webhook

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"` // ID do WhatsApp do usuário (número de telefone)
}
