package main

import(
	"code/controllers"
	"code/middleware"
	"code/database"
	"code/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port: os.Getenv("PORT")
	if port=""{
		port="8484"

	}

	app:= controllers.NewApplication(database.BookData(database.Client , "Books") , database.UserData(data));
	router:= gin.New()
	router.Use(gin.Logger())

	routes.UseRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtobasket" , app.AddToBasket())
	router.GET("/drop" , app.DropItem())
	router.GET("/checkout", app.Rent())
	router.GET("/getmefast" , app.FastRent())

	log.Fatal(router.Run(":" + port))
}