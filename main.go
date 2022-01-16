package main

import (
	"Template/Controller"
	"Template/MiddleWares"
	"Template/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecuredHandler(c *gin.Context) {
	c.JSON(200, "areeeeeeeee")
}

func main() {

	var loginService Services.LoginService = Services.NewLoginService()
	var jwtService Services.JWTService = Services.JWTAuthService()
	var loginController Controller.LoginController = Controller.LoginHandler(loginService, jwtService)

	server := gin.New()
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.Use(MiddleWares.AuthorizeJWT())
	{
		server.GET("/securedEndPoint", SecuredHandler)
	}

	port := "8080"
	server.Run(":" + port)

}
