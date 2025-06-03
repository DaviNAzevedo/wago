package webhook

type ListReply struct {
	ID          string  `json:"id"`                    // ID do item da lista selecionado
	Title       string  `json:"title"`                 // Texto do item da lista selecionado
	Description *string `json:"description,omitempty"` // Descrição do item da lista (se houver)
}
