package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)




func HashPassword(password string)  string{
	bytes , err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil{
		log.Panic(err)
	}
	return string(bytes)

}
func VerifyPassword(userPassword string,providedPassword string)(bool,string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword),[]byte(providedPassword))
	check := true
	msg := ""
	if err != nil {
		fmt.Println(err)
		msg = "Email or Password Incorrect"
		check = false
	}
	return check,msg
}