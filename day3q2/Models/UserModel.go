package Models
type User struct {
	Id      int   `json:"id"`
	Name    string `json:"name"`
	Lastname    string `json:"lastname"`
	Dob    int `json:"dob"`
	Address string `json:"address"`
	Subject  string `json:"subject"`
	Marks    string `json:"marks"`
}
func (b *User) TableName() string {
	return "students"
}