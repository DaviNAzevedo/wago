package webhook

type ContactEmail struct {
	Email *string `json:"email,omitempty"` // Endereço de email
	Type  *string `json:"type,omitempty"`  // Tipo (PERSONAL, WORK)
}
