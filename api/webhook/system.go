package webhook

type System struct {
	Body     string  `json:"body"`                // Descrição do evento do sistema
	Identity *string `json:"identity,omitempty"`  // Hash SHA256 da nova chave de identidade (para "customer_identity_changed")
	NewWaID  *string `json:"new_wa_id,omitempty"` // Novo ID do WhatsApp do cliente (para "customer_changed_number")
	WaID     *string `json:"wa_id,omitempty"`     // Antigo ID do WhatsApp do cliente (obsoleto)
	Type     string  `json:"type"`                // Tipo de evento: "customer_changed_number", "customer_identity_changed", "group_participant_add", etc.
	Customer *string `json:"customer,omitempty"`  // ID do WhatsApp do cliente que mudou de número
}
