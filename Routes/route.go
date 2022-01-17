package Routes

import (
	"Template/Controller"
	"Template/DataStores"
	"Template/MiddleWares"
	"Template/Services"
	"Template/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	client := DataStores.InitDataLayer()
	//defer client.Disconnect(context.TODO())

	//Login Stuff
	var loginService Services.LoginService = Services.NewLoginService()
	var jwtService Services.JWTService = Services.JWTAuthService()
	var loginController Controller.LoginController = Controller.LoginHandler(loginService, jwtService)

	///hello stuff
	helloController := Controller.NewHelloController()

	///studnet stuff
	var strepo repository.StudentRepository = repository.NewStudentRepository(client)
	controllerStudnet := Controller.NewStudnetController(strepo)

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	//register other handler from their controllers
	r.Use(MiddleWares.AuthorizeJWT())
	{
		r.GET("/securedEndPoint", helloController.SayHelloHandler)
		r.POST("/student/create", controllerStudnet.CreateStudentHnadler)
		r.POST("/student", controllerStudnet.GetStudnetByAgeHandler)
		r.POST("/student/greater", controllerStudnet.GetStudentByGreaterAgeHandler)
	}

	return r
}
