package webhook

type MessageContact struct {
	Name      ContactName      `json:"name"`                // Nome do contato
	Phones    []ContactPhone   `json:"phones,omitempty"`    // Telefones do contato
	Emails    []ContactEmail   `json:"emails,omitempty"`    // Emails do contato
	Urls      []ContactURL     `json:"urls,omitempty"`      // URLs (websites) do contato
	Birthday  *string          `json:"birthday,omitempty"`  // Data de nascimento (AAAA-MM-DD)
	Addresses []ContactAddress `json:"addresses,omitempty"` // Endereços do contato
	Org       *ContactOrg      `json:"org,omitempty"`       // Organização (empresa) do contato
}
