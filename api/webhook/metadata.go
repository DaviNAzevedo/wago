package webhook

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"` // Número de telefone formatado para exibição
	PhoneNumberID      string `json:"phone_number_id"`      // ID do número de telefone da conta empresarial
}
