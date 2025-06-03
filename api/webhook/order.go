package webhook

type Order struct {
	CatalogID    string        `json:"catalog_id"`     // ID do catálogo
	ProductItems []ProductItem `json:"product_items"`  // Itens do pedido
	Text         *string       `json:"text,omitempty"` // Texto opcional enviado com o pedido
}
