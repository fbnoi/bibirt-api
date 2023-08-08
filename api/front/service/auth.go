package service

import (
	"bibirt-api/api/front/dao"
	"bibirt-api/api/front/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var (
	_sign_key = ""
)

type Claims struct {
	jwt.RegisteredClaims

	Uid string `json:"uid"`
}

func RegisterAsAnonymous() (string, error) {
	user := dao.NewTmpUser()
	return newToken(user).SignedString(_sign_key)
}

func LoginFromToken(tokenString string) (string, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(_sign_key), nil
	})
	if token.Valid {
		claims, ok := token.Claims.(*Claims)
		if !ok {
			return "", errors.New("service.LoginFromToken: error claims")
		}
		user, err := dao.FindUserByUuid(claims.Uid)
		if err != nil {
			return "", errors.New("user not found")
		}
		return newToken(user).SignedString(_sign_key)
	}
	return "", errors.New("token string invalid")
}

func newToken(user *model.User) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		Uid: user.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "client_auth",
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
}
