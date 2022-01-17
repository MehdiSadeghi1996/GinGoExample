package Services

import (
	"Template/DataStores"
	"Template/entity"
	"Template/repository"
	"context"
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleInstance *repetitive

type repetitive struct {
	allUsers []entity.Users
}

func GetInstance() *repetitive {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &repetitive{allUsers: getAllUsersCredentialsFromDb()}
		} else {
			//fmt.Println("Single instance already created.")
		}
	} else {
		//fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func getAllUsersCredentialsFromDb() []entity.Users {
	client := DataStores.InitDataLayer()
	defer client.Disconnect(context.TODO())

	usrep := repository.NewUserRepository(client)

	getuser, err := usrep.ListUsers(context.TODO())
	if err != nil {
		panic(err)
	}
	return getuser
}
