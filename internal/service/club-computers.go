package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type ComputerService struct {
	repository repository.ComputerRepository
	club       repository.ClubRepository
	instance   repository.InstanceRepository
}

func NewComputerService(repository repository.ComputerRepository, club repository.ClubRepository, instance repository.InstanceRepository) *ComputerService {
	return &ComputerService{repository: repository, club: club, instance: instance}
}

func (s *ComputerService) CreateComputer(computer *model.Computers) (int, error) {
	club, err := s.club.GetClubById(computer.ClubId)
	if err != nil || club == nil {
		return -1, fmt.Errorf("Club with id:%d does not exist", computer.ClubId)
	}
	instance, err := s.instance.GetInstanceById(computer.InstanceId)
	if err != nil || instance == nil {
		return -1, fmt.Errorf("Instance with id:%d does not exist", computer.InstanceId)
	}
	getComputer, err := s.repository.GetComputerByComputerNumber(computer.ComputerNumber)
	if getComputer != nil {
		return -1, fmt.Errorf("Computer with number:%d already exists", computer.ComputerNumber)
	}
	id, err := s.repository.CreateComputer(computer)
	if err != nil {
		return -1, fmt.Errorf("Error creating computer")
	}
	return id, nil
}

func (s *ComputerService) GetComputerById(id int) (*model.Computers, error) {
	computer, err := s.repository.GetComputerById(id)
	if err != nil {
		return nil, fmt.Errorf("Error getting computer")
	}
	return computer, nil
}

func (s *ComputerService) GetAllComputers() ([]*model.Computers, error) {
	listComputer, err := s.repository.GetAllComputers()
	if err != nil {
		return nil, fmt.Errorf("Error getting computers")
	}
	if len(listComputer) == 0 {
		return nil, fmt.Errorf("No computers found")
	}
	return listComputer, nil
}

func (s *ComputerService) DeleteComputerById(id int) error {
	computer, err := s.repository.GetComputerById(id)
	if err != nil || computer == nil {
		return fmt.Errorf("Computer with id:%d does not exist", id)
	}
	err = s.repository.DeleteComputerById(id)
	if err != nil {
		return fmt.Errorf("Error deleting computer")
	}
	return nil
}

func (s *ComputerService) UpdateComputer(id int, computer *model.Computers) error {
	computerGet, err := s.repository.GetComputerById(id)
	if err != nil || computerGet == nil {
		return fmt.Errorf("Computer with id:%d does not exist", id)
	}
	instance, err := s.instance.GetInstanceById(computer.InstanceId)
	if err != nil || instance == nil {
		return fmt.Errorf("Instance with id:%d does not exist", computer.InstanceId)
	}
	getComputer, err := s.repository.GetComputerByComputerNumber(computer.ComputerNumber)
	if getComputer != nil && computerGet.ComputerNumber != computer.ComputerNumber { //check if updated computer_number is still the same
		return fmt.Errorf("Computer with id:%d already exists", computer.ComputerNumber)
	}

	club, err := s.club.GetClubById(computer.ClubId)
	if err != nil || club == nil {
		return fmt.Errorf("Club with id:%d does not exist", computer.ClubId)
	}
	err = s.repository.UpdateComputerById(id, computer)
	if err != nil {
		return fmt.Errorf("Error updating computer")
	}
	return nil
}
