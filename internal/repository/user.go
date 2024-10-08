package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
	"github.com/aidosgal/janabyte/janabyte-core/internal/utils"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) CreateUser(user *model.User) (int, error) {
	const op = "repository.CreateUser"
	var id int
	hashedPassword, err := utils.HashUserPassword(user.Password)
	if err != nil {
		return -1, fmt.Errorf("%s : %s", op, err)
	}
	err = repository.db.QueryRow(`
            INSERT INTO users (login,first_name,last_name, email, phone, password,role_id) 
            VALUES ($1, $2, $3, $4,$5,$6,$7)
            RETURNING id;
    `, user.Login, user.FirstName, user.LastName, user.Email, user.Phone, hashedPassword, user.RoleID).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s: %s", op, err)
	}

	return id, nil
}

func (repository *UserRepository) GetAllUsers() ([]*model.User, error) {
	const op = "repository.GetAllUsers"
	userList := []*model.User{}
	rows, err := repository.db.Query(`
            SELECT id,login,first_name,last_name,email,phone,password,role_id
			FROM users
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()
	for rows.Next() {

		user, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func (repository *UserRepository) GetUserById(id int) (*model.User, error) {
	const op = "repository.GetUserById"
	user := &model.User{}

	row := repository.db.QueryRow(`
		SELECT id,login, first_name,
		       last_name, email, phone, 
		       password,role_id 
		FROM users WHERE id = $1
	`, id)

	err := row.Scan(&user.Id, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password, &user.RoleID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: user with id %d not found", op, id)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return user, nil
}

func (repository *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}

	// Use QueryRow instead of Query for a single-row result
	row := repository.db.QueryRow(`
        SELECT id,login, first_name, last_name, email, phone, password,role_id
        FROM users
        WHERE phone = $1
    `, phone)

	// Scan the result directly
	err := row.Scan(&user.Id, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password, &user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user found, return nil and no error
			return nil, nil
		}
		return nil, fmt.Errorf("issue with getting user by phone %s: %s", phone, err)
	}

	return user, nil
}

func (repository *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	// Use QueryRow instead of Query for a single-row result
	row := repository.db.QueryRow(`
        SELECT id,login, first_name, last_name, email, phone, password,role_id
        FROM users
        WHERE email = $1
    `, email)

	// Scan the result directly
	err := row.Scan(&user.Id, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password, &user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user found, return nil and no error
			return nil, nil
		}
		return nil, fmt.Errorf("issue with getting user by email %s: %s", email, err)
	}

	return user, nil
}

func (repository *UserRepository) GetUserByLogin(login string) (*model.User, error) {
	const op = "repository.GetUserByLogin"
	user := &model.User{}

	// Use QueryRow instead of Query for a single-row result
	row := repository.db.QueryRow(`
        SELECT id,login, first_name, last_name, email, phone, password,role_id
        FROM users
        WHERE login = $1
    `, login)

	// Scan the result directly
	err := row.Scan(&user.Id, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password, &user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user found, return nil and no error
			return nil, nil
		}
		// Other errors during scanning
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return user, nil
}

func (repository *UserRepository) UpdateUser(id int, user *model.User) error {
	const op = "repository.UpdateUser"

	_, err := repository.db.Exec(`
            UPDATE users
            SET login = $1, first_name = $2, last_name = $3, email = $4, phone = $5, password = $6, role_id = $7
            WHERE id = $8
    `, user.Login, user.FirstName, user.LastName, user.Email, user.Phone, user.Password, &user.RoleID, id)

	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}

func (repository *UserRepository) DeleteUser(id int) error {
	const op = "repository.DeleteUser"
	_, err := repository.db.Exec(`
			DELETE FROM users WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (repository *UserRepository) SetPassword(id int, password string) error {
	const op = "repository.SetPassword"
	_, err := repository.db.Exec(`
			UPDATE users
			SET password = $1
			WHERE id = $2
	`, password, id)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (repository *UserRepository) AuthenticateByLogin(login, password string) error {
	const op = "repository.AuthenticateByLogin"
	var hashedPassword string
	row := repository.db.QueryRow(`
				SELECT password FROM users
				WHERE login = $1
	`, login)
	err := row.Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Errorf("No user with login %s found", login)
		}
		return fmt.Errorf("%s: %s", op, err)
	}
	err = utils.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (repository *UserRepository) AuthenticateByEmail(email, password string) error {
	const op = "repository.AuthenticateByEmail"
	var hashedPassword string
	row := repository.db.QueryRow(`
				SELECT password FROM users
				WHERE email = $1
	`, email)
	err := row.Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Errorf("No user with email %s found", email)
		}
		return fmt.Errorf("%s: %s", op, err)
	}
	err = utils.CheckPasswordHash(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*model.User, error) {
	user := new(model.User)

	err := rows.Scan(
		&user.Id,
		&user.Login,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.RoleID,
	)
	if err != nil {
		return nil, err
	}

	return user, err
}
