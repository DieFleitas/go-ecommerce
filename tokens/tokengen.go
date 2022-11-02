package tokens

import (
	"log"
	"os"
	"time"

	"github.com/DieFleitas/ecommerce-akhil/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email, firstname, lastname, uid string) (signedtoken, signedrefreshtoken string, err error) {
	claims := &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshtoken, nil

}

func ValidateToken(){

}

func UpdateAllTokens(){

}