package webhook

type Location struct {
	Latitude  float64 `json:"latitude"`          // Latitude da localização
	Longitude float64 `json:"longitude"`         // Longitude da localização
	Name      *string `json:"name,omitempty"`    // Nome do local
	Address   *string `json:"address,omitempty"` // Endereço do local
	URL       *string `json:"url,omitempty"`     // URL para abrir a localização em um mapa
}
