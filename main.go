package main

import (
	"github.com/gin-gonic/gin"
	"v0/controller"
	"v0/servers"
)

func main()  {
	servers.Init()
	App()
}
/*
md5
*/
func App()  {
	app := gin.Default()
	//注册
	app.POST("/enroll",controller.UserEnroll)
	//登录
	app.POST("/login/user",controller.UserLogin)
	//随机文章
	app.GET("/read/page/random",controller.RandPage)

	//单独查看一篇
	app.GET("/read/page/one/:id",controller.ReadOne)
	//user
	user :=app.Group("/user",controller.Verify)
	//写
	user.POST("write/title",controller.WriteTitleContext)
	//查看故事线
	user.GET("read/all/:nid",controller.ReadThisNovelAll)

	app.GET("/get/user/:uid",controller.GetUser)

	app.GET("/get/novel/:uid",controller.GetNovels)

	app.Run(":8080")
}