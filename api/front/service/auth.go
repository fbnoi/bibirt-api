package service

import (
	"bibirt-api/api/front/dao"
	"bibirt-api/api/front/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Uid string `json:"uid"`
}

type LoginFromTokenParam struct {
	TokenString string `json:"token" binding:"required"`
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /register/anonymous [post]
func RegisterAsAnonymous(c *gin.Context) {
	user := dao.NewTmpUser()
	tokenStr, err := newToken(user).SignedString(sign_key)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "data": tokenStr})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /login [post]
func LoginFromToken(c *gin.Context) {
	var param LoginFromTokenParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	token, _ := jwt.ParseWithClaims(param.TokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(sign_key), nil
	})
	if !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "token invalid"})
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "token invalid"})
		return
	}
	user, err := dao.FindUserByUuid(claims.Uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "user not exist"})
		return
	}
	tokenStr, err := newToken(user).SignedString(sign_key)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "system error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": tokenStr})
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
