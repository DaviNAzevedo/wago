package webhook

type ProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"` // ID do produto no varejista
	Quantity          string `json:"quantity"`            // Quantidade (como string)
	ItemPrice         string `json:"item_price"`          // Preço unitário (como string)
	Currency          string `json:"currency"`            // Moeda (ex: USD, BRL)
}
