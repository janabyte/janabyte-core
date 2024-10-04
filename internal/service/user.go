package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/logger"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
	"github.com/aidosgal/janabyte/janabyte-core/internal/utils"
	"log"
)

type UserService struct {
	repository repository.UserRepository
	roles      repository.RolesRepository
}

func NewUserService(repository repository.UserRepository, roles repository.RolesRepository) *UserService {
	return &UserService{repository, roles}
}

func (service *UserService) CreateUser(user *model.User) (int, error) {
	const op = "UserService.CreateUser"
	roles, err := service.roles.GetRoleById(user.RoleID)
	if err != nil || roles == nil {
		return -1, fmt.Errorf("Role with id %d does not exist", user.RoleID)
	}
	existLogUser, err := service.repository.GetUserByLogin(user.Login)
	if err != nil {
		return -1, fmt.Errorf("error checking login %s: %s", op, err)
	}
	if existLogUser != nil {
		// Log the existence of the user
		log.Printf("User with login %s already exists", user.Login)
		return -1, fmt.Errorf("user %s already exists", user.Login)
	}

	existsEmail, err := service.repository.GetUserByEmail(user.Email)
	if err != nil {
		return -1, fmt.Errorf("error checking email %s: %s", op, err)
	}
	if existsEmail != nil {
		log.Printf("User with email %s already exists", user.Email)
		return -1, fmt.Errorf("user with email %s already exists", user.Email)
	}

	existsPhoneUser, err := service.repository.GetUserByPhone(user.Phone)
	if err != nil {
		return -1, fmt.Errorf("error checking phone: %s", err)
	}
	if existsPhoneUser != nil {
		log.Printf("User with phone %s already exists", user.Phone)
		return -1, fmt.Errorf("user with phone %s already exists", user.Phone)
	}

	err = utils.CheckPhoneNumber(user.Password)
	if err != nil {
		return -1, fmt.Errorf("error creating user: %v", err)
	}
	err = utils.CheckEmail(user.Email)
	if err != nil {
		return -1, fmt.Errorf("error creating user: %v", err)
	}
	err = utils.CheckPhoneNumber(user.Phone)
	if err != nil {
		return -1, fmt.Errorf("error creating user: %s ", err)
	}
	err = utils.IsValidPassword(user.Password)
	if err != nil {
		return -1, fmt.Errorf("error creating user: %s", err)
	}

	id, err := service.repository.CreateUser(user)
	if err != nil {
		return -1, fmt.Errorf("failed to create user: %v ", err)
	}
	logger.SetupLogger().Debug("User created", map[string]interface{}{"user": user})
	return id, nil
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	const op = "service.GetAllUsers"
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error with repository: %s %s", err, op)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil

}

func (s *UserService) DeleteUserById(id int) error {
	const op = "service.DeleteUserById"
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return fmt.Errorf("Error with repository: %s %s", err, op)
	}
	if user == nil {
		return fmt.Errorf("User with id: %d does not exists, %s %s", id, err, op)
	}
	err = s.repository.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("cant delete user %s %s", err, op)
	}
	return nil

}

func (s *UserService) UpdateUserById(id int, user *model.User) error {
	//if user.Id == 0 {
	//	user.Id = id
	//}
	roles, err := s.roles.GetRoleById(user.RoleID)
	if err != nil || roles == nil {
		return fmt.Errorf("Role with id %d does not exist", user.RoleID)
	}
	get_user, err := s.repository.GetUserById(id)
	if err != nil {
		return fmt.Errorf("error getting user by id %d: %s", id, err)
	}
	if get_user == nil {
		return fmt.Errorf("user with id: %d does not exists %s", id, err)
	}
	hashedPassword := get_user.Password
	if user.Password != "" {
		hashedPassword, err = utils.HashUserPassword(user.Password)
		if err != nil {
			return fmt.Errorf("error hashing password: %s", err)
		}
	}

	existsLoginUser, err := s.repository.GetUserByLogin(user.Login)
	if err != nil {
		return fmt.Errorf("error getting user by login %s: %s", user.Login, err)
	}
	if existsLoginUser != nil {
		return fmt.Errorf("user with login %s already exists", user.Login)
	}
	existsEmailUser, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("error getting user by email %s: %s", user.Email, err)
	}
	if existsEmailUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	existsPhoneUser, err := s.repository.GetUserByPhone(user.Phone)
	if err != nil {
		return fmt.Errorf("error getting user by phone %s: %s", user.Phone, err)
	}
	if existsPhoneUser != nil {
		return fmt.Errorf("user with phone %s already exists", user.Phone)
	}

	//checking fields
	err = utils.CheckPhoneNumber(user.Password)
	if err != nil {
		return fmt.Errorf("error creating user  %v", err)
	}
	err = utils.CheckEmail(user.Email)
	if err != nil {
		return fmt.Errorf("error creating user  %v", err)
	}
	err = utils.CheckPhoneNumber(user.Phone)
	if err != nil {
		return fmt.Errorf("error creating user %v", err)
	}

	err = utils.IsValidPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error creating user  %v", err)
	}

	//

	err = s.repository.UpdateUser(id, user)
	if err != nil {
		return fmt.Errorf("error with updating user: %s", err)
	}
	err = s.repository.SetPassword(id, hashedPassword)
	if err != nil {
		return fmt.Errorf("Error setting password: %s", err)
	}
	return nil

}

func (s *UserService) GetUserById(id int) (*model.User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user by id %d: %s", id, err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with id: %d does not exists", id)
	}
	return user, nil
}

//func (s *UserService) GetUserByLogin(login string) (*model.User, error) {
//	const op = "service.GetUserByLogin"
//	user, err := s.repository.GetUserByLogin(login)
//	if err != nil {
//		return nil, fmt.Errorf("error getting user by login %s: %s", login, err)
//	}
//	if user == nil {
//		return nil, fmt.Errorf("user with login %s does not exists", login)
//	}
//	return user, nil
//}

//rep - service - handler
