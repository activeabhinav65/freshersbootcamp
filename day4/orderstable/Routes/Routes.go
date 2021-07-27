package Routes
import (
	"github.com/gin-gonic/gin"
	"orderstable/Controllers"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetOrders)
		grp1.POST("user", Controllers.CreateOrder)
		grp1.GET("user/:costumer_id", Controllers.GetOrderByID)
		grp1.PUT("user/:costumer_id", Controllers.UpdateOrder)
		grp1.DELETE("user/:costumer_id", Controllers.DeleteOrder)
	}
	return r
}