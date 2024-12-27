package controllers

import (
	"clothingecommerce/db"
	"clothingecommerce/models"
	"clothingecommerce/utils"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx , cancel := context.WithTimeout(context.Background() , 100*time.Second)
		defer cancel()

		var loginInput struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.BindJSON(&loginInput); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError , "Error in the Request")
			return
		}

		var user models.User

		err := db.GetDB().QueryRowContext(ctx ,
		"SELECT password_hash , image_url , role , name FROM users WHERE email =$1" , loginInput.Email).Scan(&user.Password , &user.ImageUrl , &user.Role , &user.Name)

		if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
		}

		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}
		passwordIsValid , msg := utils.VerifyPassword(user.Password , loginInput.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"Error" : msg})
			return
		}
		token , err := utils.GenerateAllTokens(user.Name , user.Role , user.Email)
		c.JSON(http.StatusOK, gin.H{
			"message": "Admin login successful",
			"token":   token,
		})
	}
} 