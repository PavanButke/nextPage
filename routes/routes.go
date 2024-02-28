package routes

import (
	"code/controllers"
	
	"github.com/gin-gonic/gin"
)

func UserRouter(incocmingRoutes *gin.Engine) {
	incocmingRoutes.POST("/users/signup", controllers.SignUp())
	incocmingRoutes.POST("users/login", controllers.Login())
	incocmingRoutes.POST("/admin/addbooks", controllers.ProductViewerAdmin())
	incocmingRoutes.POST("user/booksview", controllers.SearchProduct())
	incocmingRoutes.POST("users/search", controllers.SearchProductByQuery())
}
