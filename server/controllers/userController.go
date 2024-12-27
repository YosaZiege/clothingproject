package controllers

import (
	"clothingecommerce/db"
	"clothingecommerce/models"
	"clothingecommerce/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)





func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set a context with a timeout for the database operation
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		// Parse the JSON body into the user model
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Format"})
			return
		}

		// Validate the user data here (you may want to add validation logic)

		// Hash the password before storing it
		password := utils.HashPassword(user.Password)
		
		user.Password = password

		// Insert the new user into the database
		// Use ExecContext instead of QueryRowContext for INSERT queries
		_, err := db.GetDB().ExecContext(ctx,
			"INSERT INTO users (name, email, password_hash, image_url, role) VALUES($1, $2, $3, $4, $5)",
			user.Name, user.Email, user.Password, user.ImageUrl, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register the user"})
			log.Println(err)
			return
		}

		// Generate JWT token after the user has been successfully created
		token, err := utils.GenerateAllTokens(user.Name, user.Role, user.Email) // Ensure you're passing user.ID
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Token"})
			return
		}

		// Return a success message with the generated token
		c.JSON(http.StatusOK, gin.H{
			"message": "Registration successful",
			"token":   token,
		})
	}
}


func Login() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx , cancel := context.WithTimeout(context.Background() , 100*time.Second)
		defer cancel()		
		var loginData struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		var user models.User
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		query := "SELECT id, name, password_hash,email,role,image_url FROM public.users where name=$1 "
		err := db.GetDB().QueryRowContext(ctx,query,loginData.Username).Scan(&user.ID,&user.Name,&user.Password,&user.Email,&user.Role,&user.ImageUrl)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"Error" : "Invalid email or password"})
			}
			log.Println("Database error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		passwordIsValid , msg := utils.VerifyPassword(user.Password , loginData.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"Error" : msg})
			return
		}

		token , err := utils.GenerateAllTokens(user.Name , user.Role , user.Email)

		if err != nil{
			log.Println("Database error Generating token:", err)
		}

			c.JSON(http.StatusOK, gin.H{
			"message":       "Login successful",
			"token":token,
			"role": user.Role,
			"userId": user.ID,
			"image_url": user.ImageUrl,
			})

	}
}
func FetchAllUsers() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx , cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var users []models.User

		query := "SELECT id , email , role , name , image_url FROM users;"

		rows , err := db.GetDB().QueryContext(ctx , query)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError , gin.H{"Error" : "Could not retrieve Users "})
            return
		}
		defer rows.Close()

		for rows.Next()	{
			var user models.User

			if err := rows.Scan(&user.ID , &user.Email,&user.Role ,&user.Name , &user.ImageUrl); err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError , gin.H{"Error" : "Scanning Error of Products"})
				return
			}
			users = append(users, user)
		}
		if err = rows.Err() ; err != nil{
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError , gin.H{"Error" : "Error with Retrieving the users"})
			return
		}

		c.JSON(http.StatusOK , users)
	}
}
