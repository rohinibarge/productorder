package api

type Product struct {
	Id                 int      `json:"id,omitempty"`
	Title              string   `json:"title,omitempty"`
	Description        string   `json:"description,omitempty"`
	Price              float32  `json:"price,omitempty"`
	DiscountPercentage float32  `json:"discountpercentage,omitempty"`
	Rating             float32  `json:"rating,omitempty"`
	Stock              int      `json:"stock,omitempty"`
	Brand              string   `json:"brand,omitempty"`
	Category           string   `json:"category,omitempty"`
	Thumbnail          string   `json:"thumbnail,omitempty"`
	Images             []string `json:"images,omitempty"`
}
type ProductList struct {
	Products []Product `json:"products,omitempty"`
}
