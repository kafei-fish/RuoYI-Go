package controllers

import (
	"RuoYi-Go/db"
	"RuoYi-Go/models"
	"RuoYi-Go/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var users []models.User
	db := db.GetDB()
	result := db.Find(&users)
	fmt.Println("数据库查询结果:", result)

	util.Success(c, users)

}
