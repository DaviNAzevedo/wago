package webhook

type Message struct {
	From        string           `json:"from"`                  // ID do WhatsApp do remetente (usuário)
	ID          string           `json:"id"`                    // ID da mensagem (WAMID)
	Timestamp   string           `json:"timestamp"`             // Timestamp Unix da mensagem (como string)
	Type        string           `json:"type"`                  // Tipo da mensagem: "text", "image", "audio", "video", "document", "location", "contacts", "interactive", "order", "system", "sticker", "reaction", "button", "unknown"
	Text        *Text            `json:"text,omitempty"`        // Conteúdo para mensagens do tipo "text"
	Image       *Media           `json:"image,omitempty"`       // Conteúdo para mensagens do tipo "image"
	Audio       *Media           `json:"audio,omitempty"`       // Conteúdo para mensagens do tipo "audio"
	Video       *Media           `json:"video,omitempty"`       // Conteúdo para mensagens do tipo "video"
	Document    *Media           `json:"document,omitempty"`    // Conteúdo para mensagens do tipo "document"
	Sticker     *Media           `json:"sticker,omitempty"`     // Conteúdo para mensagens do tipo "sticker"
	Location    *Location        `json:"location,omitempty"`    // Conteúdo para mensagens do tipo "location"
	Contacts    []MessageContact `json:"contacts,omitempty"`    // Conteúdo para mensagens do tipo "contacts"
	Interactive *Interactive     `json:"interactive,omitempty"` // Conteúdo para mensagens do tipo "interactive"
	Order       *Order           `json:"order,omitempty"`       // Conteúdo para mensagens do tipo "order"
	System      *System          `json:"system,omitempty"`      // Conteúdo para mensagens do tipo "system" (ex: mudança de número)
	Reaction    *Reaction        `json:"reaction,omitempty"`    // Conteúdo para mensagens do tipo "reaction"
	Button      *Button          `json:"button,omitempty"`      // Conteúdo para cliques em botões de templates (não interativos)
	Context     *Context         `json:"context,omitempty"`     // Contexto da mensagem (ex: se é uma resposta, mensagem encaminhada)
	Referral    *Referral        `json:"referral,omitempty"`    // Informações de referência (ex: anúncio que originou a conversa)
	Identity    *Identity        `json:"identity,omitempty"`    // Informações sobre mudança de identidade do usuário
	Errors      []Error          `json:"errors,omitempty"`      // Erros específicos desta mensagem
}
