package client

// LocationPayload define o corpo de uma mensagem de localização.
type LocationPayload struct {
	Longitude float64 `json:"longitude"` // Longitude da localização.
	Latitude  float64 `json:"latitude"`  // Latitude da localização.
	Name      string  `json:"name"`      // Nome do local.
	Address   string  `json:"address"`   // Endereço do local.
}
