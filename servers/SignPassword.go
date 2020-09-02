package servers

import (
	"crypto/md5"
	"encoding/hex"
	"exchange_nolves/model"
	"regexp"
)

//
func GetPassword(password string) string {
	mPass := md5.New()
	mPass.Write([]byte(password))
	pass := mPass.Sum(nil)
	res :=hex.EncodeToString(pass)
	return res
}

func LoginUser(user model.User) (model.User, string) {
	if IsEmail(user.Email){
		user = SelectUser(user)
		if user.Id != 0{
			token , err :=GenerateToken(user)
			if err != nil {
				return model.User{}, ""
			}
			return user,token
		}
	}
	return model.User{}, ""
}

func EnrollUser(user model.User) (bool,model.User) {
	email := user.Email
	if IsEmail(email){
		user := InsertUser(user)
		return true,user
	}
	return false,model.User{}
}

func IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
