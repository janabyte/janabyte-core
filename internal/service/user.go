package service

import (
	"fmt"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type UserService struct {
    repository repository.UserRepository  
}

func NewUserService(repository repository.UserRepository) *UserService {
    return  &UserService{repository}
}

func (service *UserService) CreateUser(user model.User) error {
    _, err := service.repository.GetUserByPhone(user.Phone)
    if err == nil {
        return fmt.Errorf("User with this phone already exists")
    }

    //hashPassword
    //Beka попробуй реализовать тут
    hashPassword := user.Password

    err = service.repository.CreateUser(model.User{
        Name: user.Name,
        Email: user.Email,
        Phone: user.Phone,
        Password: hashPassword,
    })
    if err != nil {
        return err
    }

    return nil
}
