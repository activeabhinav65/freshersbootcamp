package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"productstable/Models"
	"strconv"
	"sync"
	"time"
)
func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByID(c *gin.Context) {
	ik := c.Params.ByName("id")
	var product Models.Product
	var id, _ = strconv.Atoi(ik)
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
func GetProductByProductName(c *gin.Context) {
	var mutex = &sync.Mutex{}
				mutex.Lock()
				order := Models.Order{}
				err := c.BindJSON(&order)

				if err != nil {
					panic("Failed")
				}

				id := order.Id
				i1 := order.Quantity

				product := Models.Product{}
				err = Models.GetProductByID(&product, id)
				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
				}

				if product.Quantity < i1 {
					fmt.Println("Order Failed")
				} else {
					fmt.Println("Order Success")

					err := Models.CreateProduct1(&order)
					product.Quantity = product.Quantity - i1;
					product1 := product
					ik := c.Params.ByName("id")
					id, _ := strconv.Atoi(ik)
					err1 := Models.GetProductByID(&product1, id)
					if err1 != nil {
						c.JSON(http.StatusNotFound, product1)
					}
					c.BindJSON(&product1)
					err1 = Models.UpdateProduct(&product1, string(id))
					if err1 != nil {
						c.AbortWithStatus(http.StatusNotFound)
					} else {
						c.JSON(http.StatusOK, product1)
					}
					if err != nil {
						fmt.Println(err.Error())
						c.AbortWithStatus(http.StatusNotFound)
					} else {
						c.JSON(http.StatusOK, order)
					}
				}
				mutex.Unlock()
				time.Sleep(time.Second)
	/*order := Models.Order{}
	err := c.BindJSON(&order)

	if err != nil {
		panic("Failed")
	}

    id:= order.Id
    i1:= order.Quantity

	product := Models.Product{}
	err = Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if product.Quantity < i1 {
		fmt.Println("Order Failed")
	} else {
		fmt.Println("Order Success")

		err := Models.CreateProduct1(&order)
        product.Quantity=product.Quantity-i1;
        product1 :=product
		ik := c.Params.ByName("id")
		id, _ := strconv.Atoi(ik)
		err1 := Models.GetProductByID(&product1, id)
		if err1 != nil {
			c.JSON(http.StatusNotFound, product1)
		}
		c.BindJSON(&product1)
		err1 = Models.UpdateProduct(&product1, string(id))
		if err1 != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, product1)
		}
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, order)
		}
	}*/
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	ik := c.Params.ByName("id")
	id, _ := strconv.Atoi(ik)
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, string(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}