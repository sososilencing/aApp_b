package controller

import (
	"encoding/json"
	"exchange_nolves/model"
	"exchange_nolves/servers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context)  {
	uid := c.Param("uid")
	id,err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400,nil)
		c.Abort()
	}
	user := model.User{
		Id:       id,
	}
	user = servers.GetUser(user)
	userJson,err := json.Marshal(user)
	if err != nil {
		c.JSON(400,nil)
		c.Abort()
	}
	c.JSON(200,model.GetMessage(200,string(userJson)))
}

func GetNovels(c *gin.Context)  {
	uid := c.Param("uid")
	id,err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400,nil)
		c.Abort()
	}
	user := model.User{
		Id:       id,
	}
	novels := servers.GetNovels(user);
	novelsJson,err := json.Marshal(novels)
	if err != nil {
		c.JSON(400,nil)
		c.Abort()
	}
	c.JSON(200,model.GetMessage(999,string(novelsJson)))
}
//post
func UserLogin(c *gin.Context)  {
	user := model.User{}
	userJson := c.PostForm("user")

	json.Unmarshal([]byte(userJson),&user)
	
	user , Token := servers.LoginUser(user)
	if Token == ""{
		c.JSON(200,model.GetMessage(412,"账号或密码错误"))
	}else {
		token := model.UserToken{
			Token:     Token,
			User:       user,
		}

		by,err :=json.Marshal(token)
		if err != nil {
			c.JSON(200,model.GetMessage(400,""))
		}else {
			c.JSON(200,model.GetMessage(200,string(by)))
		}
	}
	// 进行验证
}
//post
func UserEnroll(c *gin.Context)  {
	user := model.User{}
	userJson := c.PostForm("user")
	fmt.Println(userJson)
	json.Unmarshal([]byte(userJson),&user)
	// 一些判断
	ok,user :=servers.EnrollUser(user)
	fmt.Println(user)
	if ok{
		by,_ := json.Marshal(user)

		c.JSON(200,model.GetMessage(200, string(by)))
	}else {
		c.JSON(200,model.GetMessage(400,"你的邮箱或者xxx不符合规范"))
	}
}
//定义为cookies
func Verify(c *gin.Context)  {
	token := c.Request.Header.Get("cook")
	loginToken, err := servers.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loginToken)
	//判断logintoken 有没有值
	if loginToken!=nil {
		c.Next()
	}else {
		c.Abort()
	}
}
//post
func WriteTitleContext(c *gin.Context){
	mNovel := &model.MNovels{
		Uid:      0,
		Nid:      0,
		Title:    "",
		Context:  "",
		Pid: 0,
	}

	mNovelJson := c.PostForm("novel")

	json.Unmarshal([]byte(mNovelJson),mNovel)
	mNovel.Context = SensitiveWordFilter(mNovel.Context)
	//一些判断 然后存入数据库
	mNovel = servers.WriteMNovel(*mNovel)

	if mNovel == nil{
		c.JSON(400,model.GetMessage(400,""))
	}else {
		by , _ := json.Marshal(mNovel)
		c.JSON(200,model.GetMessage(200,string(by)))
	}
}
