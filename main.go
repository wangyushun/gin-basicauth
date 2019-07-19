package main

import (
	"gin-basicauth/basicauth"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:cnic.cn@tcp(159.226.235.106:3306)/goharbor?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&UserProfile{})

	app := gin.Default()
	app.Use(basicauth.BasicAuth(db, &UserProfile{})) //input db and user model
	app.GET("/api/v1/users/", getUserHandler)

	app.Run(":9000")
}

func getUserHandler(ctx *gin.Context) {

	user, exists := ctx.Get(basicauth.AuthUserKey)
	if !exists {
		ctx.JSON(401, "user not authenticated")
		return
	}
	u, ok := user.(*UserProfile)
	if ok {
		ctx.JSON(200, u)
		return
	}
	ctx.JSON(401, "user not authenticated")
}
