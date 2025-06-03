package webhook

type Identity struct {
	Acknowledged     bool   `json:"acknowledged"`      // Se a mudança de identidade foi reconhecida pelo destinatário
	CreatedTimestamp string `json:"created_timestamp"` // Timestamp Unix de quando a identidade mudou
	Hash             string `json:"hash"`              // Hash da nova identidade
}
