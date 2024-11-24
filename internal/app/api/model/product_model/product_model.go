package product_model

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	UrlImage string  `json:"urlImage"`
}
