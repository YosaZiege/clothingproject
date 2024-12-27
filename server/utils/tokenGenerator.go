package utils

import (
	"clothingecommerce/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct{
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims

}
func GenerateAllTokens(username string , role string, email string ) (signedToken string,err error) {
	claims := &SignedDetails{
		Email: email,
		Role: role,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetEnv("JWT_SECRET", "")))
	if err != nil {
		log.Panic("Error generating main token:", err)
		return "", err
	}
	// refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.GetEnv("JWT_SECRET", "")))
	// 	if err != nil {
	// 		log.Panic("Error generating refresh token:", err)
	// 		return "", "", err
	// 	}
	return signedToken , nil
}