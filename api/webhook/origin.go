package webhook

type Origin struct {
	Type string `json:"type"` // Tipo de origem: "user_initiated", "business_initiated", "referral_conversion", "authentication"
}
