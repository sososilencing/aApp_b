package servers

import (
	"github.com/gbrlsnchs/jwt/v3"
	"time"
	"v0/model"
)

func getHs() *jwt.HMACSHA {
	hs  := jwt.NewHS256([]byte("你真的猜不到哈"))
	return hs
}
/*
生成验证token
 */
func GenerateToken(user model.User)  (string,error){
	nowtime :=  time.Now();
	expireTime := nowtime.Add(7*24*time.Hour);
	pl := model.LoginToken{
		Payload: jwt.Payload{
			Issuer:         "iiRoxi",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(expireTime),
			NotBefore:      nil,
			IssuedAt:       jwt.NumericDate(nowtime),
			JWTID:          "",
		},
		ID:      user.Id,
		Email:   user.Email,
	}

	token ,err := jwt.Sign(pl,getHs())
	return string(token),err
}
/*
验证token 是否 完整
 */
func VerifyToken(token string) (*model.LoginToken, error) {
	pl := &model.LoginToken{}
	_ ,err :=jwt.Verify([]byte(token),getHs(),pl)
	return pl,err
}