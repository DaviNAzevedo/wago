package webhook

type ButtonReply struct {
	ID    string `json:"id"`    // ID do botão clicado
	Title string `json:"title"` // Texto do botão clicado
}
