package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

var AppSecret = ""       //viper.GetString会设置这个值(32byte长度)
var AppIss = "inception" //这个值会被viper.GetString重写
var UserClaims = "UserClaims"

// 自定义payload结构体,不建议直接使用 dgrijalva/jwt-go `jwt.StandardClaims`结构体.因为他的payload包含的用户信息太少.
type userStdClaims struct {
	jwt.StandardClaims
	//*model.User
	BaseClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	AuthorityId uint
}

// Valid 实现 `type Claims interface` 的 `Valid() error` 方法,自定义校验内容
func (c userStdClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}
	if !c.VerifyIssuer(AppIss, true) {
		return errors.New("token's issuer is wrong")
	}
	if c.ID < 1 {
		return errors.New("invalid user in jwt")
	}
	return
}

func GenerateToken(m BaseClaims, d time.Duration) (string, error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.ID),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		BaseClaims:     m,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

// ParseToken 解析payload的内容,得到用户信息
// gin-middleware 会使用这个方法
func ParseToken(tokenString string) (*BaseClaims, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.BaseClaims, nil
}
