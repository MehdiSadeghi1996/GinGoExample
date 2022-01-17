package Controller

import "github.com/gin-gonic/gin"

type HelloController interface {
	SayHelloHandler(ctx *gin.Context)
}

type helloController struct {
}

func NewHelloController() HelloController {
	return &helloController{}
}

func (r *helloController) SayHelloHandler(c *gin.Context) {
	c.JSON(200, "Hello to You!")
}
