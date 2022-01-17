package Services

import (
	"Template/entity"
	"errors"
	"github.com/ahmetb/go-linq/v3"
)

type LoginService interface {
	Login(username string, password string) (entity.Users, error)
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() *loginService {
	return &loginService{
		authorizedPassword: "",
		authorizedUsername: "",
	}
}

func (service *loginService) Login(username string, password string) (entity.Users, error) {

	getFromRepSer := GetInstance()

	userTargeted := linq.From(getFromRepSer.allUsers).Where(func(c interface{}) bool {
		return c.(entity.Users).Username == username && c.(entity.Users).Password == password
	}).Select(func(c interface{}) interface{} {
		return c.(entity.Users)
	}).First()

	if userTargeted != nil {
		return userTargeted.(entity.Users), nil
	} else {
		return entity.Users{Username: "", Password: ""}, errors.New("not found user")
	}

}
