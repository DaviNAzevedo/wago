package webhook

type ReferredProduct struct {
	CatalogID         string `json:"catalog_id"`          // ID do catálogo do produto
	ProductRetailerID string `json:"product_retailer_id"` // ID do produto no varejista
}
