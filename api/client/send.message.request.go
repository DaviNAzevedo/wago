package client

// SendMessageRequest é a estrutura de nível superior para enviar uma mensagem.
// Documentação: https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages
type SendMessageRequest struct {
	MessagingProduct string `json:"messaging_product"`        // Sempre "whatsapp"
	RecipientType    string `json:"recipient_type,omitempty"` // Padrão: "individual". Outra opção: "group" (atualmente não suportado pela Cloud API para envio)
	To               string `json:"to"`                       // ID do WhatsApp do destinatário (número de telefone)
	Type             string `json:"type,omitempty"`           // Tipo da mensagem: "text", "image", "audio", "video", "document", "location", "contacts", "interactive", "template", "sticker"
	// Se Type não for especificado, o WhatsApp tentará inferir a partir dos campos preenchidos.
	// É recomendado sempre especificar o Type.

	Text        *TextPayload        `json:"text,omitempty"`
	Image       *MediaPayload       `json:"image,omitempty"`
	Audio       *MediaPayload       `json:"audio,omitempty"`
	Video       *MediaPayload       `json:"video,omitempty"`
	Document    *DocumentPayload    `json:"document,omitempty"`
	Sticker     *MediaPayload       `json:"sticker,omitempty"`
	Location    *LocationPayload    `json:"location,omitempty"`
	Contacts    *ContactsPayload    `json:"contacts,omitempty"`
	Interactive *InteractivePayload `json:"interactive,omitempty"`
	Template    *TemplatePayload    `json:"template,omitempty"`
	Context     *MessageContext     `json:"context,omitempty"`     // Para responder a uma mensagem específica
	PreviewURL  *bool               `json:"preview_url,omitempty"` // Para mensagens de texto, se deve gerar preview de URL. Padrão: false.
}
