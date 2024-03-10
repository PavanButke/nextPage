package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"code/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	bookCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(bookCollection, userCollection *mongo.Collection) *Application {

	return &Application{
		bookCollection: bookCollection,
		userCollection: userCollection,
	}

}

func (app *Application) AddToBasket() gin.HandlerFunc {

	return func(c *gin.Context) {
		//check if book is available
		bookQueryID := c.Query("id")
		if bookQueryID == "" {
			log.Println("Book Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Book Id is empty!"))
			return

		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User Id is empty!"))
			return

		}

		//
		bookID, err := primitive.ObjectIDFromHex(bookQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		database.AddBookToBasket(ctx, app.bookCollection, app.userCollection, bookID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}
		c.IndentedJSON(http.StatusOK, "Succesfully Added to  Basket!")
	}
}

func RemoveItem(app *Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookQueryID := c.Query("id")
		if bookQueryID == "" {
			log.Println("Book Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Book Id is empty!"))
			return

		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User Id is empty!"))
			return

		}

		//
		bookID, err := primitive.ObjectIDFromHex(bookQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		err = database.RemoveItem(ctx, app.bookCollection, app.userCollection, bookID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, "Successfully Removed from  Basket!")
	}
}

// func GetItemFromBasket() gin.HandlerFunc {

// }

// func BorrowFromBasket() gin.HandlerFunc {

// }

func (app *Application) GrabIt() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookQueryID := c.Query("id")
		if bookQueryID == "" {
			log.Println("Book Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Book Id is empty!"))
			return

		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User Id is Empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User Id is empty!"))
			return

		}

		//
		bookID, err := primitive.ObjectIDFromHex(bookQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		err = database.GrabIt(ctx, app.bookCollection, app.userCollection, bookID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Succesfully Added a Book!")
	}
}
