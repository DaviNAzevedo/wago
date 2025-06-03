package client

// InteractivePayload define o corpo de uma mensagem interativa.
// Documentação: https://developers.facebook.com/docs/whatsapp/guides/interactive-messages
type InteractivePayload struct {
	Type   string             `json:"type"` // "button", "list", "product", "product_list", "flow"
	Header *InteractiveHeader `json:"header,omitempty"`
	Body   InteractiveBody    `json:"body"` // Obrigatório
	Footer *InteractiveFooter `json:"footer,omitempty"`
	Action InteractiveAction  `json:"action"` // Obrigatório
}

// InteractiveHeader define o cabeçalho de uma mensagem interativa.
type InteractiveHeader struct {
	Type     string           `json:"type"`               // "text", "video", "image", "document"
	Text     *string          `json:"text,omitempty"`     // Necessário se type="text". Até 60 caracteres.
	Video    *MediaPayload    `json:"video,omitempty"`    // Necessário se type="video".
	Image    *MediaPayload    `json:"image,omitempty"`    // Necessário se type="image".
	Document *DocumentPayload `json:"document,omitempty"` // Necessário se type="document".
}

// InteractiveBody define o corpo de uma mensagem interativa.
type InteractiveBody struct {
	Text string `json:"text"` // Obrigatório. Conteúdo principal. UTF-8. Até 1024 caracteres.
}

// InteractiveFooter define o rodapé de uma mensagem interativa.
type InteractiveFooter struct {
	Text string `json:"text"` // Obrigatório. Conteúdo do rodapé. UTF-8. Até 60 caracteres.
}

// InteractiveAction define a ação de uma mensagem interativa.
type InteractiveAction struct {
	Button             *string              `json:"button,omitempty"`               // Necessário para type="list". Texto do botão da lista. Até 20 caracteres.
	Buttons            []InteractiveButton  `json:"buttons,omitempty"`              // Necessário para type="button". Máximo de 3 botões.
	Sections           []InteractiveSection `json:"sections,omitempty"`             // Necessário para type="list". Mínimo 1, máximo 10 seções.
	CatalogID          *string              `json:"catalog_id,omitempty"`           // Necessário para type="product_list".
	ProductRetailerID  *string              `json:"product_retailer_id,omitempty"`  // Necessário para type="product".
	Mode               *string              `json:"mode,omitempty"`                 // Para Flows. "draft" ou "published".
	FlowMessageVersion *string              `json:"flow_message_version,omitempty"` // Para Flows.
	FlowToken          *string              `json:"flow_token,omitempty"`           // Para Flows.
	FlowID             *string              `json:"flow_id,omitempty"`              // Para Flows.
	FlowCTA            *string              `json:"flow_cta,omitempty"`             // Para Flows.
	FlowAction         *string              `json:"flow_action,omitempty"`          // Para Flows. "navigate" ou "data_exchange".
	FlowActionPayload  *FlowActionPayload   `json:"flow_action_payload,omitempty"`  // Para Flows.
}

// FlowActionPayload define o payload para uma ação de Flow.
type FlowActionPayload struct {
	Screen string                 `json:"screen"`         // Obrigatório. ID da tela inicial do Flow.
	Data   map[string]interface{} `json:"data,omitempty"` // Dados para pré-preencher a tela.
}

// InteractiveButton define um botão em uma mensagem interativa do tipo "button".
type InteractiveButton struct {
	Type  string      `json:"type"` // Sempre "reply"
	Reply ButtonReply `json:"reply"`
}

// ButtonReply define a resposta de um botão interativo.
type ButtonReply struct {
	ID    string `json:"id"`    // Obrigatório. ID único do botão. Até 256 caracteres.
	Title string `json:"title"` // Obrigatório. Texto do botão. Até 20 caracteres.
}

// InteractiveSection define uma seção em uma mensagem interativa do tipo "list".
type InteractiveSection struct {
	Title *string   `json:"title,omitempty"` // Título da seção. Até 24 caracteres.
	Rows  []ListRow `json:"rows"`            // Obrigatório. Mínimo 1, máximo 10 linhas por seção.
}

// ListRow define uma linha em uma seção de lista interativa.
type ListRow struct {
	ID          string  `json:"id"`                    // Obrigatório. ID único da linha. Até 200 caracteres.
	Title       string  `json:"title"`                 // Obrigatório. Texto da linha. Até 24 caracteres.
	Description *string `json:"description,omitempty"` // Descrição da linha. Até 72 caracteres.
}
