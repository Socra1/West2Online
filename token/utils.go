package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// 定义密钥
var JWTsecret = []byte("JWT")

// 定义token的claims
type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"user_name"`
	PassWord  string `json:"password"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, username string, password string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)
	Cla := Claims{
		Id:       id,
		Username: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	} //token有效期为一天
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Cla) //jwt密钥生成token
	s, err := tokenClaims.SignedString(JWTsecret)
	return s, err
}

// 验证token

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
