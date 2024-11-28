package jwt

import (
	"github.com/577683719/common/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("server.jwtSecret"))
var expired = 24 * time.Hour

type Claims struct {
	//用户id
	UserID int64 `json:"user_id"`
	//用户名
	UserName string `json:"user_name"`
	//账号
	AccountNo string `json:"account_no"`
	//手机号
	Phone string `json:"phone"`
	//邮箱
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(userID int64, userName string, accountNo string, phone string, nickname string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expired)
	claims := Claims{
		UserID:    userID,
		UserName:  userName,
		AccountNo: accountNo,
		Phone:     phone,
		Nickname:  nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "38384-SearchEngine",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Conf.Server.JwtSecret))
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Server.JwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func GetUserIdByToken(token string) int64 {
	claims, err := ParseToken(token)
	if err != nil {
		panic("token错误")
	}
	return claims.UserID

}
