package Models
type Order struct {
	Costumer_id      int   `json:"costumer_id"`
	Product_name    string `json:"product_name"`
	Quantity  int `json:"quantity"`
	Price   string `json:"price"`
	ID int `json:"id"`
}
func (b *Order) TableName() string {
	return "orders"
}