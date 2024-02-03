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

	app:= controllers.NewApplication(database.ProductData(database.Client , "Products") , database.UserData(data));
	router:= gin.New()
	router.Use(gin.Logger())
}