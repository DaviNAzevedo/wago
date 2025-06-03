package webhook

type Button struct {
	Payload string `json:"payload"` // Payload definido no botão do template
	Text    string `json:"text"`    // Texto do botão clicado
}
