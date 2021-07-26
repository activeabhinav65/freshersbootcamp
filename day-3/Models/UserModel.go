package Models
type User struct {
	Id      int    `json:"id"`
	Name   string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
func (b *User) TableName() string {
	return "firsttable"
}
