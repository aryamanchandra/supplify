package models

type ProductCheckout struct {
	ProductID    string `json:"product_id"`
	Buyer        string `json:"buyer"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Product struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	Title        string `json:"title"`
	Brand        string `json:"brand"`
	Seller       string `json:"seller"`
	UPC          string `json:"upc"`
	MFD          string `json:"mfd"`
	Cost         string `json:"cost"`
	Availability bool   `json:"availability"`
}
