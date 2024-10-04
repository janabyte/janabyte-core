package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type RoleService struct {
	repository repository.RolesRepository
}

func NewRoleService(repository repository.RolesRepository) *RoleService {
	return &RoleService{repository}
}

func (s *RoleService) CreateRole(roles *model.Roles) (int, error) {
	const op = "RoleService.CreateRole"
	role, _ := s.repository.GetRoleByName(roles.Name)
	if role != nil {
		return -1, fmt.Errorf("%s already exists", roles.Name)
	}
	id, err := s.repository.CreateRole(roles)
	if err != nil {
		return -1, fmt.Errorf("error creating role")
	}
	return id, nil
}
func (s *RoleService) DeleteRole(id int) error {
	const op = "RoleService.DeleteRole"
	role, _ := s.repository.GetRoleById(id)
	if role == nil {
		return fmt.Errorf("role with id: %d not found", id)
	}
	err := s.repository.DeleteRoleById(id)
	if err != nil {
		return fmt.Errorf("error deleting role with id %d: %s", id, err)
	}
	return nil
}

func (s *RoleService) GetRoleById(id int) (*model.Roles, error) {
	const op = "RoleService.GetRole"
	role, err := s.repository.GetRoleById(id)
	if err != nil {
		return nil, fmt.Errorf("error getting role with id %d: %s", id, err)
	}
	if role == nil {
		return nil, fmt.Errorf("role with id %d not found", id)
	}
	return role, nil
}

func (s *RoleService) GetAllRoles() ([]*model.Roles, error) {
	const op = "RoleService.GetAllRoles"
	roles, err := s.repository.GetAllRoles()
	if err != nil {
		return nil, fmt.Errorf("error getting all roles: %s", err)
	}
	if len(roles) == 0 {
		return nil, fmt.Errorf("no roles found")
	}
	return roles, nil
}

func (s *RoleService) UpdateRoleById(id int, role *model.Roles) error {
	const op = "RoleService.UpdateRoleById"
	roles, _ := s.repository.GetRoleById(id)
	if roles == nil {
		return fmt.Errorf("role with id %d not found", id)
	}
	err := s.repository.UpdateRole(id, role)
	if err != nil {
		return fmt.Errorf("error updating role with id %d: %s", id, err)
	}
	return nil
}
