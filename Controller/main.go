package Controller

import (
	"Template/Routes"
)

func main() {

	port := "8080"
	engineRouter := Routes.InitRoutes()
	engineRouter.Run(":" + port)

}
