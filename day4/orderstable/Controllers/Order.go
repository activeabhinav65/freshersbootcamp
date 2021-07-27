package Controllers
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orderstable/Models"
)
func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func GetOrderByID(c *gin.Context) {
	/*var order []Models.Order
	err := Models.GetAllOrders(&order)
	costumer_id := c.Params.ByName("costumer_id")

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}*/

	costumer_id := c.Params.ByName("costumer_id")
	var order Models.Order
	err := Models.GetOrderByID(&order, costumer_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func UpdateOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	c.BindJSON(&order)
	err = Models.UpdateOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}