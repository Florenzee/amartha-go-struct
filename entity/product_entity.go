package entity

type ProductEntity struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	TaxRate     float64 `json:"tax_rate"`
}
