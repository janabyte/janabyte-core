package repository

import (
	"database/sql"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type RolesRepository struct {
	db *sql.DB
}

func NewRolesRepository(db *sql.DB) *RolesRepository {
	return &RolesRepository{db}
}

func (r *RolesRepository) CreateRole(roles *model.Roles) (int, error) {
	const op = "repository.CreateRole"
	var id int
	err := r.db.QueryRow(`
			INSERT INTO roles (name)
			VALUES ($1)
			RETURNING id
	`, roles.Name).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Error creating role %s: %w", roles.Name, err)
	}
	return id, nil
}

func (r *RolesRepository) DeleteRoleById(id int) error {
	const op = "repository.DeleteRole"
	_, err := r.db.Exec(`
			DELETE FROM roles
			WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *RolesRepository) GetAllRoles() ([]*model.Roles, error) {
	const op = "repository.GetAllRoles"
	var roles []*model.Roles
	rows, err := r.db.Query(`
			SELECT id,name
			FROM roles
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		role := &model.Roles{}
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		roles = append(roles, role)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return roles, nil
}

func (r *RolesRepository) GetRoleById(id int) (*model.Roles, error) {
	const op = "repository.GetRoleById"
	role := &model.Roles{}
	row := r.db.QueryRow(`
			SELECT id,name
			FROM roles
			WHERE id = $1
	`, id)
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return role, nil

}

func (r *RolesRepository) UpdateRole(id int, role *model.Roles) error {
	const op = "repository.UpdateRole"
	_, err := r.db.Exec(`
			UPDATE roles
			SET name = $1
			WHERe id = $2
	`, role.Name, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *RolesRepository) GetRoleByName(name string) (*model.Roles, error) {
	const op = "repository.GetRoleByName"
	row, err := r.db.Query(`
			SELECT id,name
			FROM roles
			WHERE name = $1
	`, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer row.Close()
	role := &model.Roles{}
	if row.Next() {
		err = row.Scan(&role.Id, &role.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}
	return role, nil

}
