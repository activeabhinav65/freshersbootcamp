package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"productstable/Config"
	"productstable/Models"
	"productstable/Routes"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	r := Routes.SetupRouter()
	r.Run()
}