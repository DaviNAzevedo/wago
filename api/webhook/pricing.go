package webhook

type Pricing struct {
	Billable     bool   `json:"billable"`      // Se a conversa é cobrável
	PricingModel string `json:"pricing_model"` // Modelo de precificação (ex: "CBP")
	Category     string `json:"category"`      // Categoria da conversa (igual a Origin.Type)
}
