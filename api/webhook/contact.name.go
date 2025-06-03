package webhook

type ContactName struct {
	FormattedName string  `json:"formatted_name"`        // Nome completo formatado
	FirstName     *string `json:"first_name,omitempty"`  // Primeiro nome
	LastName      *string `json:"last_name,omitempty"`   // Sobrenome
	MiddleName    *string `json:"middle_name,omitempty"` // Nome do meio
	Suffix        *string `json:"suffix,omitempty"`      // Sufixo (Jr., Sr.)
	Prefix        *string `json:"prefix,omitempty"`      // Prefixo (Dr., Mr.)
}
