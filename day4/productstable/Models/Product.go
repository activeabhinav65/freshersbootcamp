package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"productstable/Config"
)
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error;
	err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
func CreateProduct1(order *Order) (err error) {
	if err = Config.DB.Create(order).Error;
	err!=nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
func GetProductByProductName(order *Order,productname string) (err error) {
	if err = Config.DB.Where("product_name= ?",productname).First(order).Error;
	err!=nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}