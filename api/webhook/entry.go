package webhook

type Entry struct {
	ID      string   `json:"id"` // ID da Conta Empresarial do WhatsApp (WABA ID)
	Changes []Change `json:"changes"`
}
