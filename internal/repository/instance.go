package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type InstanceRepository struct {
	db *sql.DB
}

func NewInstanceRepository(db *sql.DB) *InstanceRepository {
	return &InstanceRepository{db}
}

func (r *InstanceRepository) CreateInstance(instance *model.Instance) (int, error) {
	const op = "repository.CreateInstance"
	var id int
	err := r.db.QueryRow(`
			INSERT INTO instances(name,icon_url)
			VALUES ($1, $2)
			RETURNING id;
	`, instance.Name, instance.IconUrl).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
func (r *InstanceRepository) GetInstanceById(id int) (*model.Instance, error) {
	const op = "repository.GetInstanceById"
	instance := &model.Instance{}
	row := r.db.QueryRow(`
			SELECT id,name,icon_url
			FROM instances
			WHERE id = $1
	`, id)
	err := row.Scan(&instance.Id, &instance.Name, &instance.IconUrl)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: user with id %d not found", op, id)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return instance, nil
}

func (r *InstanceRepository) GetInstanceByName(name string) (*model.Instance, error) {
	const op = "repository.GetInstanceByName"
	instance := &model.Instance{}
	row := r.db.QueryRow(`
			SELECT id,name,icon_url
			FROM instances
			WHERE name = $1
	`, name)
	err := row.Scan(&instance.Id, &instance.Name, &instance.IconUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, err)
		} else if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}
	return instance, nil
}
func (r *InstanceRepository) GetAllInstances() ([]*model.Instance, error) {
	const op = "repository.GetAllInstances"
	var listInstances []*model.Instance
	rows, err := r.db.Query(`
			SELECT id,name,icon_url
			FROM instances
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		instance := &model.Instance{}
		err = rows.Scan(&instance.Id, &instance.Name, &instance.IconUrl)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		listInstances = append(listInstances, instance)
	}
	return listInstances, nil
}
func (r *InstanceRepository) DeleteInstanceById(id int) error {
	const op = "repository.DeleteInstanceById"
	_, err := r.db.Exec(`
			DELETE FROM instances
			WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *InstanceRepository) UpdateInstanceById(id int, instance *model.Instance) error {
	const op = "repository.UpdateInstanceById"
	_, err := r.db.Exec(`
			UPDATE instances
			SET name = $1, icon_url = $2
			WHERE id = $3
	`, instance.Name, instance.IconUrl, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
