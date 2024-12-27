package routers

import (
	"github.com/gin-gonic/gin"
	"clothingecommerce/controllers"
)

func AuthRouter (incomingRoutes *gin.Engine){
	incomingRoutes.POST("/signup" , controllers.Register())
	incomingRoutes.POST("/signin" , controllers.Login())
	incomingRoutes.GET("/fetchusers" , controllers.FetchAllUsers())
}