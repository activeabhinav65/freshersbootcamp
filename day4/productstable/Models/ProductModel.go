package Models
type Product struct {
	Id      int   `json:"id"`
	Product_Name    string `json:"product_name"`
	Price   string `json:"price"`
	Quantity  int `json:"quantity"`
}
type Order struct{
	Id int `json:"id"`
	Product_Name string  `json:"product_name"`
	Quantity int `json:"quantity"`
	Costumer_Id int `json:"costumer_id"`
}
func (b *Product) TableName() string {
	return "products"
}
func (b *Order) TableName() string {
	return "orders"
}