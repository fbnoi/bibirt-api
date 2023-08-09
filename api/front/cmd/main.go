package main

import (
	"bibirt-api/api/front/dao"
	"bibirt-api/api/front/service"

	"github.com/gin-gonic/gin"
)

func bootstrap() {
	dao.Bootstrap()
	service.Bootstrap()
}

func main() {
	bootstrap()
	router := gin.Default()
	g := router.Group("/api/v1/auth")
	g.POST("/register/anonymous", service.RegisterAsAnonymous)
	g.POST("/login", service.LoginFromToken)
	router.Run(":8080")
}
