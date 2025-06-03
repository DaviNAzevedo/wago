package webhook

type ContactAddress struct {
	Street      *string `json:"street,omitempty"`       // Rua
	City        *string `json:"city,omitempty"`         // Cidade
	State       *string `json:"state,omitempty"`        // Estado/Província
	Zip         *string `json:"zip,omitempty"`          // CEP
	Country     *string `json:"country,omitempty"`      // País
	CountryCode *string `json:"country_code,omitempty"` // Código do país (ex: US, BR)
	Type        *string `json:"type,omitempty"`         // Tipo (HOME, WORK)
}
