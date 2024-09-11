package auth

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
	)




func Createjwt(email string, secret string) (string, error){
	signingKey := []byte("secret")	
	claims := jwt.MapClaims{
		"Email": email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString(signingKey)

	if err != nil{
		return "" ,err
	}

	return tokenString, nil
}
