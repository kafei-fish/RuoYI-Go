package routes

import (
	"RuoYi-Go/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//配置路由组
	r.GET("/login", controllers.CreateUser)
	r.Run(":8080")
}
