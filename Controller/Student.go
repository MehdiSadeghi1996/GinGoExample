package Controller

import (
	"Template/entity"
	"Template/entity/RequestModel"
	"Template/repository"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudnetController interface {
	CreateStudentHnadler(ctx *gin.Context)
	GetStudnetByAgeHandler(ctx *gin.Context)
	GetStudentByGreaterAgeHandler(ctx *gin.Context)
}

type studnetController struct {
	studentRepository repository.StudentRepository
}

func NewStudnetController(stRepo repository.StudentRepository) StudnetController {
	return &studnetController{
		studentRepository: stRepo,
	}
}

func (r *studnetController) CreateStudentHnadler(ctx *gin.Context) {

	var studnet entity.Student
	err := ctx.ShouldBind(&studnet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)

	}

	err = r.studentRepository.CreateStudnet(context.TODO(), studnet)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, "Created")
}

func (r *studnetController) GetStudnetByAgeHandler(ctx *gin.Context) {

	var reqModel RequestModel.StudnetRequestModel
	err := ctx.ShouldBindJSON(&reqModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)

	}

	stude, err := r.studentRepository.FindByAgeOne(context.TODO(), reqModel.Age)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, stude)

}

func (r *studnetController) GetStudentByGreaterAgeHandler(ctx *gin.Context) {
	var reqModel RequestModel.StudnetRequestModel
	err := ctx.ShouldBindJSON(&reqModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)

	}
	stude, err := r.studentRepository.FindByGreaterAge(context.TODO(), reqModel.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, stude)
}
