package Routes
import (
	"github.com/gin-gonic/gin"
	"productstable/Controllers"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetProducts)
		grp1.POST("user", Controllers.CreateProduct)
		grp1.GET("user/:id", Controllers.GetProductByID)
		grp1.PUT("user/:id", Controllers.UpdateProduct)
		grp1.DELETE("user/:id", Controllers.DeleteProduct)
		grp1.POST("user1",Controllers.GetProductByProductName)
	}
	return r
}
