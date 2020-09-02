package servers

import (
	"exchange_nolves/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const SQL = "root:@tcp(localhost:3306)/nolves?charset=utf8"

func Init() {
	db , _ = gorm.Open("mysql",SQL)

	db.SingularTable(true)
}

func InsertUser(user model.User)  model.User{
	db.Create(&user)
	return user
}

func UpdateUser(user model.User){
	db.Where("id = ?",user.Id).Update(user)
}

func GetUser(user model.User)  model.User{
	sql := "select id,email,nickname,sex from user where id = ?"
	db.Raw(sql,user.Id).Scan(&user)
	return user
}

func GetNovels(user model.User)  []model.MNovels{
	sql := "select nid,uid,title,pid,context from m_novels where uid = ?"
	novels := []model.MNovels{}
	db.Raw(sql,user.Id).Scan(&novels)
	return novels
}

func SelectUser(user model.User) model.User {
	sql :="select id,email,nickname,sex from user where email = ? and password = ?"
	db.Raw(sql,user.Email,user.Password).Scan(&user)
	return user
}
//敏感词过滤
func SWF() []model.Words {
	words := []model.Words{}
	sql := "select words from words"
	db.Raw(sql).Scan(&words)
	return words
}

//创建一个主题目
func InsertMNovel(novel model.MNovels) model.MNovels {
	db.Create(&novel)
	return novel
}
func UpdateMNovel(novel model.MNovels)  {
	//一个所有权的问题
}

//找一个章节
func SelectOneNovel(nid int) *model.MNovels {
	sql := "select nid,uid,title,pid from m_novels where nid = ?"
	novel := &model.MNovels{}
	db.Raw(sql,nid).Scan(&novel)
	return novel
}

//找 一个章节内容
func SelectMNovel(novel *model.MNovels) *model.MNovels {
	sql := "select nid,uid,title,context,pid from m_novels where nid = ?"
	db.Raw(sql,novel.Nid).Scan(&novel)
	return novel
}


//找章节的衍生
func SelectChildrenNovel(nid int) []*model.Novels {
	novels := []*model.Novels{}
	sql := "select nid,uid,title,pid from m_novels where pid = ?"
	db.Raw(sql,nid).Scan(&novels)
	return novels
}

func RandNovel() model.MNovels {
	sql := "SELECT " +
		"nid,uid,title,context,pid " +
		"FROM m_novels AS t1 " +
		"JOIN (SELECT round( rand( ) * ( ( SELECT MAX( nid ) FROM m_novels )" +
		" - ( SELECT MIN( nid ) FROM m_novels ) ) ) " +
		"+ ( SELECT MIN( nid ) FROM m_novels ) AS nid2 ) AS t2 " +
		"WHERE " +
		"t1.nid >= t2.nid2 ORDER BY t1.nid LIMIT 1"
	//fmt.Println(sql)
	novel := model.MNovels{}
	db.Raw(sql).Scan(&novel)
	//fmt.Println(novel)
	return novel
}