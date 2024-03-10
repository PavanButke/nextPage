package controllers

import (
	"code/database"
	"code/models"
	generate "code/tokens"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var BookCollection *mongo.Collection = database.BookData(database.Client, "Books")
var Validate = validator.New()

func VerifyPassword(userpassword string, givenpassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := Validate.Struct(user)

		if validationErr != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		//check if user with email exists
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"User Already exists with this email": err})
		}

		phone, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		if phone > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"User Already exists with this phone": err})
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		//token generation
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		//remaining props
		user.UserBasket = make([]models.BookUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Borrow_Status = make([]models.Borrow, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)

		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Create User"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Successfully Signed In")
	}
}

func Login(username string, email string, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 130*time.Second)
		defer cancel()

		var user models.User
		var founduser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Username Or Password is Incorrect!": err})
			return
		}

		PasswordValid, msg := VerifyPassword(*user.Password, *founduser.Password)

		defer cancel()
		if !PasswordValid {
			c.JSON(http.StatusInternalServerError, gin.H{"err": msg})
			return
		}

		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *&founduser.User_ID)

		defer cancel()

		//whenever user login update all tokens
		generate.UpdateAllTokens(token, refreshToken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)
	}
}

// func bookViewerAdmin() gin.HandlerFunc {

// }

// func searchBook() gin.HandlerFunc {

// }

func searchBookByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchBooks []models.Book

		queryParam := c.Query("name")

		//check if is query is empty
		if queryParam == "" {
			log.Println("Query is Empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query is empty"})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		searchquerydb, err := BookCollection.Find(ctx, bson.M{"product_name ": bson.M{"$regex": queryParam}})
		if err != nil {
			c.IndentedJSON(404, "Something went wrong")
			return
		}

		err = searchquerydb.All(ctx, &searchBooks)
		if err != nil {
			c.IndentedJSON(404, "Something went wrong")
			return
		}
		defer searchquerydb.Close(ctx)
		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "Invalid Request!")
			return
		}
		defer cancel()
		c.IndentedJSON(200, searchBooks)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
