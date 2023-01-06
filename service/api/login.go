package api

import (
	"github.com/Tiveron999/Wasa/service/database"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

func storeKey() {
	// Read the secret key file into a byte slice
	secretKey, err := ioutil.ReadFile("secret.key")
	if err != nil {
		log.Fatal(err)
	}
	token := database.getItem("token")
	// Use the byte slice as the secret key
	jwt.Parse(token, secretKey)
}
