package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type ComputerRepository struct {
	db *sql.DB
}

func NewComputerRepository(db *sql.DB) *ComputerRepository {
	return &ComputerRepository{db}
}

func (r *ComputerRepository) CreateComputer(computers *model.Computers) (int, error) {
	const op = "repository.CreateComputer"
	var id int
	err := r.db.QueryRow(`
			INSERT INTO clubs_computers (computer_number,is_near_to_next,
			                       is_near_to_prev,gpu,cpu,ram,
			                       x_pos,y_pos,instance_id,club_id)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
			RETURNING id;
	`, computers.ComputerNumber,
		computers.IsNearToNext,
		computers.IsNearToPrev,
		computers.Gpu,
		computers.Cpu,
		computers.Ram,
		computers.XPos,
		computers.YPos,
		computers.InstanceId,
		computers.ClubId).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil

}

func (r *ComputerRepository) GetComputerById(computer_id int) (*model.Computers, error) {
	const op = "repository.GetComputerById"
	computers := &model.Computers{}
	row := r.db.QueryRow(`
			SELECT id,computer_number,
			       is_near_to_next,
			       is_near_to_prev,
			       gpu,
			       cpu,
			       ram,
			       x_pos,
			       y_pos,
			       instance_id,
			       club_id
			FROM clubs_computers
			WHERE id = $1
	`, computer_id)
	err := row.Scan(&computers.Id, &computers.ComputerNumber,
		&computers.IsNearToNext,
		&computers.IsNearToPrev,
		&computers.Gpu,
		&computers.Cpu,
		&computers.Ram,
		&computers.XPos,
		&computers.YPos,
		&computers.InstanceId,
		&computers.ClubId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return computers, nil
}

func (r *ComputerRepository) DeleteComputerById(computer_id int) error {
	const op = "repository.DeleteComputerById"
	_, err := r.db.Exec(`
			DELETE FROM clubs_computers
			WHERE id = $1
	`, computer_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *ComputerRepository) GetAllComputers() ([]*model.Computers, error) {
	const op = "repository.GetAllComputers"
	var listComputers []*model.Computers
	rows, err := r.db.Query(`
			SELECT id,computer_number,
			       is_near_to_next,
			       is_near_to_prev,
			       gpu,
			       cpu,
			       ram,
			       x_pos,
			       y_pos,
			       instance_id,
			       club_id
			FROM clubs_computers
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		computers := &model.Computers{}
		err = rows.Scan(&computers.Id, &computers.ComputerNumber,
			&computers.IsNearToNext,
			&computers.IsNearToPrev,
			&computers.Gpu,
			&computers.Cpu,
			&computers.Ram,
			&computers.XPos,
			&computers.YPos,
			&computers.InstanceId,
			&computers.ClubId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		listComputers = append(listComputers, computers)
	}
	return listComputers, nil
}

func (r *ComputerRepository) UpdateComputerById(computer_id int, computer *model.Computers) error {
	const op = "repository.UpdateComputerById"
	_, err := r.db.Exec(`
			UPDATE clubs_computers
			SET computer_number = $1,
			    is_near_to_next = $2,
			    is_near_to_prev = $3,
			    gpu = $4,
			    cpu = $5,
			    ram = $6,
			    x_pos = $7,
			    y_pos = $8,
			    instance_id = $9,
			    club_id = $10
			WHERE id = $11
	`, computer.ComputerNumber,
		computer.IsNearToNext,
		computer.IsNearToPrev,
		computer.Gpu,
		computer.Cpu,
		computer.Ram,
		computer.XPos,
		computer.YPos,
		computer.InstanceId,
		computer.ClubId, computer_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *ComputerRepository) GetComputerByComputerNumber(computer_number int) (*model.Computers, error) {
	const op = "repository.GetComputerByComputerNumber"
	computers := &model.Computers{}
	row := r.db.QueryRow(`
			SELECT id,computer_number,
			       is_near_to_next,
			       is_near_to_prev,
			       gpu,
			       cpu,
			       ram,
			       x_pos,
			       y_pos,
			       instance_id,
			       club_id
			FROM clubs_computers
			WHERE computer_number = $1
	`, computer_number)
	err := row.Scan(&computers.Id, &computers.ComputerNumber,
		&computers.IsNearToNext,
		&computers.IsNearToPrev,
		&computers.Gpu,
		&computers.Cpu,
		&computers.Ram,
		&computers.XPos,
		&computers.YPos,
		&computers.InstanceId,
		&computers.ClubId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: computer with number %d not found", op, computer_number)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return computers, nil
}
