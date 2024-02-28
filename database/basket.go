package database

import(

)

var(
	ErrCantFindBook= errors.New("can't find the product")
	ErrCantDecodeBooks= errors.New("can't find the product")
	ErrUserIdIsNotValid= errors.New("this user is invalid")
	ErrCantUpdateUser= errors.New("Can not add the books")
	ErrCantRemoveItemBasket= errors.New("Can not remove this book from basket ")
	ErrCantGetItem= errors.New("Was unable to get book from basket")
	ErrCantBorrowBasketItem= errors.New("can't update the basket")
)

func addToBasket() gin.HandlerFunc {

}

func removeItem() gin.HandlerFunc {

}

func getItemFromBasket() gin.HandlerFunc {

}

func borrowFromBasket() gin.HandlerFunc{

}

func grabIt() gin.HandlerFunc{
	
}