package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type InstanceService struct {
	repository repository.InstanceRepository
}

func NewServiceInstance(repository repository.InstanceRepository) *InstanceService {
	return &InstanceService{repository}
}

func (s *InstanceService) CreateInstance(instance *model.Instance) (int, error) {
	inst, err := s.repository.GetInstanceByName(instance.Name)
	if err == nil || inst != nil {
		return -1, fmt.Errorf("instance %s already exists", instance.Name)
	}
	id, err := s.repository.CreateInstance(instance)
	if err != nil {
		return -1, fmt.Errorf("Error creating instance")
	}
	return id, nil
}

func (s *InstanceService) GetInstanceById(id int) (*model.Instance, error) {
	instance, err := s.repository.GetInstanceById(id)
	if err != nil {
		return nil, fmt.Errorf("Error getting instance")
	}
	return instance, nil
}

func (s *InstanceService) DeleteInstanceById(id int) error {
	instance, err := s.repository.GetInstanceById(id)
	if err != nil || instance == nil {
		return fmt.Errorf("Instance does not exists")
	}
	err = s.repository.DeleteInstanceById(id)
	if err != nil {
		return fmt.Errorf("Error deleting instance")
	}
	return nil
}

func (s *InstanceService) GetAllInstances() ([]*model.Instance, error) {
	instances, err := s.repository.GetAllInstances()
	if err != nil {
		return nil, fmt.Errorf("Error getting all instances")
	}
	if len(instances) == 0 {
		return nil, fmt.Errorf("No instances found")
	}
	return instances, nil
}

func (s *InstanceService) UpdateInstanceById(id int, instance *model.Instance) error {
	instanceGet, err := s.repository.GetInstanceById(id)
	if err != nil || instanceGet == nil {
		return fmt.Errorf("Instance does not exists")
	}
	err = s.repository.UpdateInstanceById(id, instance)
	if err != nil {
		return fmt.Errorf("Error updating instance")
	}
	return nil
}
