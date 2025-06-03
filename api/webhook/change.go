package webhook

type Change struct {
	Value ChangeValue `json:"value"`
	Field string      `json:"field"` // Ex: "messages", "message_template_status_update", etc.
}
