package tokens

import (
	"os"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/form3tech-oss/jwt-go"
)

type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid
	jwt.StandardClaims
}

func TokenGenerator(){
	claims:= &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(48)).Unix(),
	}

	var SECRET_KEY = os.Getenv("SECRET_KEY")

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims).SignedString([]byte(SECRET_KEY))

	if err != nil{
		return "","" , err
	}

	refreshToken, err:= jwt.NewWithClaims(jwt.SigningMethodHS256 , claims).SignedString([]byte(SECRET_KEY))
	if err != nil{
		return "","" , err
	}
	return token,refreshToken , err
}

func ValidateToken(signedtoken string)(clainmclaims *SignedDetails , msg string){
	token , err := jwt.ParseWithClaims(signedtoken , &SignedDSignedDetails{}, func(token *jwt.Token)(interface{},error){
		return []byte(SECRET_KEY),nil
	});

	if err != nil{
		msg = err.Error()
		return
	}

	claims , ok := token.Claims.(*SignedDSignedDetails)
	if !ok{
		msg= "Invalid Token"
		return
	}

	claims.ExpiresAt < time.Now().Local().unix(){
		msg ="Token Expired"
		return
	}
}

func UpdateAllTokens(signedtoken string , signedfreshtoken string , userid string){

	var ctx , cance := context.WithTimeout(context.Background(), 100*time.Second)

	var updateobj primtive.D

	updateobj = append(updateobj , bson.E{Key:"token" , Value: signedtoken})
	updateobj = append(updateobj , bson.E{Key:"refresh_token", Value: signedfreshtoken})
	update_at, _ := time.Parse(time.RFC3339, time.Now().Format);

	upsert := true

	filter := bson.M{"user_id": : userid }
	opt := options.UpdateOne(ctx , filter , bson.D{
		{Key:"$set" , Value: updateobj},
	}, &opt )

	defer cancel()

	if err != nil {
		log.panic(err)
		return
	}
}