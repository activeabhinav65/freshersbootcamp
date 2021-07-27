package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"orderstable/Config"
	"orderstable/Models"
	"orderstable/Routes"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Order{})
	r := Routes.SetupRouter()
	r.Run()
}