package client

// ContactsPayload define o corpo de uma mensagem de contatos.
type ContactsPayload struct {
	Contacts []ContactPayload `json:"contacts"` // Array de objetos de contato. Pelo menos um é necessário.
}

// ContactPayload representa um único contato a ser enviado.
type ContactPayload struct {
	Addresses []AddressPayload `json:"addresses,omitempty"`
	Birthday  *string          `json:"birthday,omitempty"` // Formato AAAA-MM-DD
	Emails    []EmailPayload   `json:"emails,omitempty"`
	Name      NamePayload      `json:"name"` // Obrigatório
	Org       *OrgPayload      `json:"org,omitempty"`
	Phones    []PhonePayload   `json:"phones,omitempty"`
	Urls      []URLPayload     `json:"urls,omitempty"`
}

// AddressPayload representa um endereço de um contato.
type AddressPayload struct {
	Street      *string `json:"street,omitempty"`
	City        *string `json:"city,omitempty"`
	State       *string `json:"state,omitempty"`
	Zip         *string `json:"zip,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"` // Código do país de duas letras.
	Type        *string `json:"type,omitempty"`         // Padrão: HOME. Outros: WORK.
}

// EmailPayload representa um email de um contato.
type EmailPayload struct {
	Email *string `json:"email,omitempty"`
	Type  *string `json:"type,omitempty"` // Padrão: HOME. Outros: WORK.
}

// NamePayload representa o nome de um contato.
type NamePayload struct {
	FormattedName string  `json:"formatted_name"` // Obrigatório. Nome completo formatado.
	FirstName     *string `json:"first_name,omitempty"`
	LastName      *string `json:"last_name,omitempty"`
	MiddleName    *string `json:"middle_name,omitempty"`
	Suffix        *string `json:"suffix,omitempty"`
	Prefix        *string `json:"prefix,omitempty"`
}

// OrgPayload representa a organização de um contato.
type OrgPayload struct {
	Company    *string `json:"company,omitempty"`
	Department *string `json:"department,omitempty"`
	Title      *string `json:"title,omitempty"`
}

// PhonePayload representa um telefone de um contato.
type PhonePayload struct {
	Phone *string `json:"phone,omitempty"` // Número de telefone.
	Type  *string `json:"type,omitempty"`  // Padrão: CELL. Outros: HOME, WORK, MAIN.
	WaID  *string `json:"wa_id,omitempty"` // ID do WhatsApp (número de telefone).
}

// URLPayload representa uma URL (website) de um contato.
type URLPayload struct {
	URL  *string `json:"url,omitempty"`
	Type *string `json:"type,omitempty"` // Padrão: HOME. Outros: WORK.
}
