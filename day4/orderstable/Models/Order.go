package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"orderstable/Config"
)
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}
func GetOrderByID(order *Order, costumer_id string) (err error) {
	if err = Config.DB.Where("costumer_id = ?", costumer_id).First(order).Error; err != nil {
		return err
	}
	return nil
}
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}
func DeleteOrder(order *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(order)
	return nil
}