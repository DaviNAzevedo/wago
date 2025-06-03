package client

// TemplatePayload define o corpo de uma mensagem de template.
// Documentação: https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates
type TemplatePayload struct {
	Name       string              `json:"name"`                 // Obrigatório. Nome do template.
	Language   TemplateLanguage    `json:"language"`             // Obrigatório. Objeto de idioma.
	Components []TemplateComponent `json:"components,omitempty"` // Componentes do template (header, body, footer, buttons).
}

// TemplateLanguage define o idioma de um template.
type TemplateLanguage struct {
	Code string `json:"code"` // Obrigatório. Código do idioma (ex: "en_US", "pt_BR").
	// Policy: "deterministic" (recomendado) ou "fallback".
}

// TemplateComponent define um componente de um template (header, body, footer, buttons).
type TemplateComponent struct {
	Type       string                  `json:"type"`                 // "header", "body", "footer", "button"
	SubType    *string                 `json:"sub_type,omitempty"`   // Para botões: "quick_reply", "url", "call", "copy_code", "catalog", "mpm". Para carrosséis: "carousel".
	Parameters []TemplateParameter     `json:"parameters,omitempty"` // Parâmetros para preencher variáveis no template.
	Index      *string                 `json:"index,omitempty"`      // Para botões de URL dinâmicos ou botões de ação. "0", "1", "2", etc.
	Cards      []TemplateCardComponent `json:"cards,omitempty"`      // Para componentes de carrossel.
}

// TemplateCardComponent define um card dentro de um componente de carrossel.
type TemplateCardComponent struct {
	Components []TemplateComponent `json:"components"` // Cada card tem seus próprios componentes (header, body, buttons)
}

// TemplateParameter define um parâmetro para um componente de template.
// O tipo do parâmetro determina qual campo será usado.
type TemplateParameter struct {
	Type     string             `json:"type"` // "text", "currency", "date_time", "image", "document", "video", "payload" (para botões de resposta rápida ou mpm)
	Text     *string            `json:"text,omitempty"`
	Currency *CurrencyParameter `json:"currency,omitempty"`
	DateTime *DateTimeParameter `json:"date_time,omitempty"`
	Image    *MediaPayload      `json:"image,omitempty"`    // Apenas Link ou ID
	Document *DocumentPayload   `json:"document,omitempty"` // Apenas Link ou ID
	Video    *MediaPayload      `json:"video,omitempty"`    // Apenas Link ou ID
	Payload  *string            `json:"payload,omitempty"`  // Para botões de resposta rápida ou mpm.
}

// CurrencyParameter define um parâmetro do tipo moeda.
type CurrencyParameter struct {
	FallbackValue string `json:"fallback_value"` // Valor de fallback (ex: "$10.99")
	Code          string `json:"code"`           // Código da moeda ISO 4217 (ex: "USD")
	Amount1000    int64  `json:"amount_1000"`    // Valor multiplicado por 1000 (ex: 10990 para $10.99)
}

// DateTimeParameter define um parâmetro do tipo data/hora.
type DateTimeParameter struct {
	FallbackValue string `json:"fallback_value"` // Valor de fallback (ex: "February 25, 1977")
	// Você pode especificar um ou mais dos seguintes, dependendo do formato do seu template:
	DayOfWeek  *int    `json:"day_of_week,omitempty"` // 1 (Segunda) a 7 (Domingo)
	Year       *int    `json:"year,omitempty"`
	Month      *int    `json:"month,omitempty"` // 1 a 12
	DayOfMonth *int    `json:"day_of_month,omitempty"`
	Hour       *int    `json:"hour,omitempty"`      // 0 a 23
	Minute     *int    `json:"minute,omitempty"`    // 0 a 59
	Calendar   *string `json:"calendar,omitempty"`  // "GREGORIAN", "SOLAR_HIJRI"
	Timestamp  *int64  `json:"timestamp,omitempty"` // Timestamp Unix em SEGUNDOS
}
