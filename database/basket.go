package database

import (
	"context"
	"errors"
	"log"

	"code/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindBook         = errors.New("can't find the product")
	ErrCantDecodeBooks      = errors.New("can't find the product")
	ErrUserIdIsNotValid     = errors.New("this user is invalid")
	ErrCantUpdateUser       = errors.New("Can not add the books")
	ErrCantRemoveItemBasket = errors.New("Can not remove this book from basket ")
	ErrCantGetItem          = errors.New("Was unable to get book from basket")
	ErrCantBorrowBasketItem = errors.New("can't update the basket")
)

func AddBookToBasket(ctx context.Context, bookCollection, userCollection *mongo.Collection, bookID primitive.ObjectID, userID string) error {
	searchfromdb, err := bookCollection.Find(ctx, bson.M{"_id": bookID})
	if err != nil {
		log.Println(err)
		return ErrCantFindBook
	}
	var bookbasket []models.BookUser
	err = searchfromdb.All(ctx, &bookbasket)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeBooks
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "bookbasket", Value: bson.D{{Key: "$each", Value: bookbasket}}}}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}

func RemoveItem() gin.HandlerFunc {

}

// func getItemFromBasket() gin.HandlerFunc {

// }

// func borrowFromBasket() gin.HandlerFunc {

// }

func GrabIt() gin.HandlerFunc {

}
