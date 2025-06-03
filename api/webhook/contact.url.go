package webhook

type ContactURL struct {
	URL  *string `json:"url,omitempty"`  // URL
	Type *string `json:"type,omitempty"` // Tipo (HOME, WORK)
}
