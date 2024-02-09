package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct(
	ID 						primitive.ObjectID
	First_Name				*string
	Last_Name				*string
	Password				*string
	Email					*string
	Phone					*string
	Token					*string
	Refresh_Token			*string
	Created_At				time.Time
	Updated_At				time.Time
	User_ID					*string
	UserBasket				[]BookUser
	Address_Details			[]Address
	Borrow_Status			[]Borrow
)

type Book struct{
	Book_ID				primitive.ObjectID
	Book_Name			*string
	Author				*string
	Price				*uint64
	Image				*string
	Recommendations		int
}

type BookUser struct{
	Book_ID				primitive.ObjectID
	Book_Name			*string
	Author				*string
	Price				*uint64
	Image				*string

}

type Address struct{
	Address_ID 				primitive.ObjectID
	House					*string
	Street					*string
	City					*string
	Pincode					*string
}

type Borrow struct{
	Borrow_ID			primitive.ObjectID
	Borrow_Basket		[]BookUser
	Borrowed_At			time.Time
	Price				*uint64
	Discount			int
	Payment_Method		Payment
} 

type Payment struct{
	Digital		
	COD
}