package main

import (
	"log"
	"os"

	"code/controllers"
	"code/database"
	"code/middleware"
	"code/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8484"

	}

	app := controllers.NewApplication(database.BookData(database.Client, "Books"), database.UserData(database.Client, "Users"))
	
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	router.GET("/addtobasket", app.AddToBasket())
	// router.GET("/drop", app.DropItem())
	// router.GET("/checkout", app.Rent())
	// router.GET("/getmefast", app.FastRent())

	log.Fatal(router.Run(":" + port))
}
