package service

import (
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
)

type ClubService struct {
	repository  repository.ClubRepository
	userService UserService
}

func NewClubService(repository repository.ClubRepository, userService UserService) *ClubService {
	return &ClubService{repository, userService}
}

func (s *ClubService) CreateClub(club *model.Club) (int, error) {
	const op = "service.CreateClub"
	_, err := s.userService.GetUserById(club.UserId)
	if err != nil {
		return -1, fmt.Errorf("user with id: %d not found", club.UserId)
	}
	id, err := s.repository.CreateClub(club)
	if err != nil {
		return -1, fmt.Errorf("Error creating club:%s", err)
	}
	return id, nil
}

func (s *ClubService) GetClubById(id int) (*model.Club, error) {
	const op = "service.GetClubById"

	club, err := s.repository.GetClubById(id)
	if err != nil {
		return nil, fmt.Errorf("Error getting club")
	}
	return club, nil
}

func (s *ClubService) GetAllClubs() ([]*model.Club, error) {
	const op = "service.GetAllClubs"
	clubs, err := s.repository.GetAllClubs()
	if err != nil {
		return nil, fmt.Errorf("Error getting clubs")
	}
	if len(clubs) == 0 {
		return nil, fmt.Errorf("No clubs found")
	}
	return clubs, nil
}

func (s *ClubService) UpdateClub(id int, club *model.Club) error {
	const op = "service.UpdateClub"
	_, err := s.userService.GetUserById(club.UserId)
	if err != nil {
		return fmt.Errorf("user with id: %d not found", club.UserId)
	}
	_, err = s.repository.GetClubById(id)
	if err != nil {
		return fmt.Errorf("Error getting club")
	}
	err = s.repository.UpdateClub(id, club)
	if err != nil {
		return fmt.Errorf("Error updating club")
	}
	return nil
}
func (s *ClubService) DeleteClub(id int) error {
	const op = "service.DeleteClub"
	_, err := s.repository.GetClubById(id)
	if err != nil {
		return fmt.Errorf("Error club with id: %d not found", id)
	}
	err = s.repository.DeleteClubById(id)
	if err != nil {
		return fmt.Errorf("Error deleting club")
	}
	return nil
}
