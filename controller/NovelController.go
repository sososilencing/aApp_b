package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"v0/model"
	"v0/servers"
)

//一个敏感词过滤
func SensitiveWordFilter(words string)  string{
	word :=servers.SWF()
	for _ ,v := range  word{
		words = strings.ReplaceAll(words,v.Words,"*")
	}

	// 从数据库 拿出 过滤的敏感信息 然后 一一 替换掉
	return words
}

//这里有个问题不知道 是不是主页面
func ReadThisNovelAll(c *gin.Context)  {
	nid :=c.Param("nid")
	n , err := strconv.Atoi(nid)
	novel := servers.SelectOneNovel(n)

	if err != nil {
		c.JSON(400,model.GetMessage(400,""))
	}else {
		//开始寻找全部

		if novel.Nid == 0 {
			c.JSON(400,model.GetMessage(400,""))
		}else {
			oneNovel, novels := servers.FindOneNovel(*novel)

			one := model.One{
				oneNovel,
				novels,
			}
			by, err :=json.Marshal(one)
			if err != nil {
				c.Abort()
			}
			fmt.Println(string(by))
			c.JSON(200,model.GetMessage(200, string(by)))
		}
	}
}
func ReadOne(c *gin.Context)  {
	nid := c.Param("id")
	fmt.Println(nid)
	id , err := strconv.Atoi(nid)
	if err != nil {
		c.JSON(400,"")
		c.Abort()
	}
	novel := &model.MNovels{
		Nid:     id,
	}
	novel = servers.SelectMNovel(novel)
	c.JSON(200,novel)
}

func RandPage(c *gin.Context)  {
	novel :=servers.RandOne()
	if novel !=nil{
		by , _ := json.Marshal(novel)
		c.JSON(200,model.GetMessage(200,string(by)))
	}else {
		c.JSON(200,model.GetMessage(400,"未知错误"))
	}
}

func Comment(c *gin.Context)  {
	comment := model.Comment{
		Id:      0,
		Uid:     0,
		Pid:     0,
		Comment: "",
	}
	commentJson := c.PostForm("comment")

	json.Unmarshal([]byte(commentJson),&comment)

	//敏感词过滤？？？
}