package routers

import (
	"github.com/gin-gonic/gin"
	"clothingecommerce/controllers"
)

func AdminRouter (incomingRoutes *gin.Engine){
	incomingRoutes.POST("/admin/signin" , controllers.AdminLogin())
}