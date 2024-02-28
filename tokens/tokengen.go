package tokens

import(
	jwt "github.com/dgrijalva/jwt-go"

)

type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid
	jwt.StandardClaims
}

func TokenGenerator()
{

}

func ValidateToken()
{

}

func UpdateAllTokens()
{
	
}