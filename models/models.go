package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json: "_id": bson:"_id"`
	First_Name      *string            `json: "first_name"	validate:"required , min=2 , max=25"`
	Last_Name       *string            `json: "last_name"	validate:"required	, min=2 , max=25"`
	Password        *string            `json: "password" 	validate:"required , min=6"`
	Email           *string            `json: "email"		validate:"email , required"`
	Phone           *string            `json: "phone"`
	Token           *string            `json: "token"`
	Refresh_Token   *string            `json: "refreshed_token`
	Created_At      time.Time          `json: "created_at"`
	Updated_At      time.Time          `json: "updated_at"`
	User_ID         *string            `json: "user_id"`
	UserBasket      []BookUser         `json: userbasket bson:"userbasket"`
	Address_Details []Address          ``
	Borrow_Status   []Borrow
}

type Book struct {
	Book_ID         primitive.ObjectID `bson:"_id"`
	Book_Name       *string            `json: "book_name" validate:"required"`
	Author          *string            `json: "author"`
	Price           *uint64            `json: "price"`
	//may be a set of arr for img
	Image           *string            `json: "image"	validate:"required"`
	Recommendations int                `json: "recommendections"`
}

type BookUser struct {
	Book_ID   primitive.ObjectID `bson: "_id"`
	Book_Name *string            `json: "book_name" bson: "book_name"`
	Author    *string            `json: "author" bson:"author" `
	Price     *uint64            `json: "price" bson:"price"`
	Image     *string            `json: "image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json: "house_name" bson:"house_name"`
	Street     *string            `json: "street_name" bson:"street_name"`
	City       *string            `json: "city" bson:"city" validate:"required"`
	Pincode    *string            `json: "pincode" bson:"pincode"`
}

type Borrow struct {
	Borrow_ID     primitive.ObjectID `bson: "_id"`
	Borrow_Basket []BookUser         `json: "borrow_basket" bson:"borrow_basket"`
	Return_Date   time.Time          `json: "return_date" bson:"return_date"`
	Borrowed_At   time.Time          `json: "borrowed_at bson:"borrowed_at"`
	Price         *uint64            `json: "price" bson:"price"`
	Discount      int                `json: "discount" bson:"discount"`
	//Payment_Method Payment            `json: "payment_method" bson:"payment_method"`
}

// type Payment struct {
// 	Digital
// 	COD
// }
