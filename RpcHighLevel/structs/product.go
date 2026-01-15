package structs

type Product struct {
	ProductID int     `gorm:"primaryKey;column:product_id" json:"product_id"`
	Name      string  `gorm:"column:name;not null" json:"name"`
	Price     float64 `gorm:"column:price;not null" json:"price"`
	Quantity  int     `gorm:"column:quantity;not null" json:"quantity"`
}

// TableName specifies the table name for Product model
func (Product) TableName() string {
	return "products"
}

type ProductResponse struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
