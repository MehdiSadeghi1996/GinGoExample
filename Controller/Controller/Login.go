package Controller

import (
	"Template/Services"
	"Template/entity"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService Services.LoginService
	jWtService   Services.JWTService
}

func LoginHandler(loginService Services.LoginService, jWtService Services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential entity.Users
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated, err := controller.loginService.Login(credential.Username, credential.Password)
	if err != nil {
		return ""
	}

	if isUserAuthenticated.Role == "admin" {
		return controller.jWtService.GenerateToken(credential.Username, true)
	} else {
		return controller.jWtService.GenerateToken(credential.Username, false)
	}
	return ""
}
