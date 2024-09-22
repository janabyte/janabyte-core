package service

import (
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type UserManipulator interface {
	GetAllUsers() ([]*model.User, error)
	CreateUser(user *model.User) error
	GetUserByLogin(login string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	GetUserById(id int) (*model.User, error)
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository}
}

//func (service *UserService) CreateUser(user model.User) error {
//	const op = "UserService.CreateUser"
//	existUser, err := service.repository.GetUserByLogin(user.Login)
//	if err == nil && existUser != nil {
//		return fmt.Errorf("User with: %s already exists :%s", user.Login, op)
//	}
//	err = service.repository.CreateUser(&user)
//	if err != nil {
//		return fmt.Errorf("Failed to create user: %v :%s", err, op)
//	}
//
//	return nil
//}

//func (service *UserService)
