package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

//func (repository *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
//	rows, err := repository.db.Query("SELECT * FROM users WHERE phone = ?", phone)
//	if err != nil {
//		return nil, err
//	}
//
//	user := new(model.User)
//
//	for rows.Next() {
//		user, err = scanRowIntoUser(rows)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	if user.Id == 0 {
//		return nil, fmt.Errorf("User not found")
//	}
//
//	return user, nil
//}

func (repository *UserRepository) CreateUser(user *model.User) error {
	const op = "repository.CreateUser"
	hashedPassword, err := HashUserPassword(user.Password)
	if err != nil {
		return fmt.Errorf("%s : %s", op, err)
	}
	_, err = repository.db.Exec(`
            INSERT INTO users (login,first_name,last_name, email, phone, password) 
            VALUES ($1, $2, $3, $4,$5,$6);
    `, user.Login, user.FirstName, user.LastName, user.Email, user.Phone, hashedPassword)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}

func (repository *UserRepository) GetAllUsers() ([]*model.User, error) {
	const op = "repository.GetAllUsers"
	userList := []*model.User{}
	rows, err := repository.db.Query(`
            SELECT id,login,first_name,last_name,email,phone,password 
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

func (repository *UserRepository) GetUserByLogin(login string) (*model.User, error) {
	const op = "repository.GetUserByLogin"
	user := &model.User{}
	row, err := repository.db.Query(`
            SELECT * FROM users
            WHERE login = $1
    `, login)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer row.Close()
	if row.Next() {
		user, err = scanRowIntoUser(row)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}
	return user, nil
}

func (repository *UserRepository) UpdateUser(user *model.User) error {
	const op = "repository.UpdateUser"

	get_user, err := repository.GetUserById(user.Id)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	hashedPassword := get_user.Password

	if get_user.Password != "" {
		hashedPassword, err = HashUserPassword(get_user.Password)
		if err != nil {
			return fmt.Errorf("%s : %s", op, err)
		}
	}
	_, err = repository.db.Exec(`
            UPDATE users
            SET login = $1, first_name = $2, last_name = $3, email = $4, phone = $5,password = $6
            WHERE id = $7
    `, user.Login, user.FirstName, user.LastName, user.Email, user.Phone, hashedPassword, user.Id)
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

func (repository *UserRepository) GetUserById(id int) (*model.User, error) {
	const op = "repository.GetUserById"
	user := &model.User{}

	row := repository.db.QueryRow(`
		SELECT id, login, first_name, last_name, email, phone, password FROM users WHERE id = $1
	`, id)

	err := row.Scan(&user.Id, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: user with id %d not found", op, id)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return user, nil
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
	)
	if err != nil {
		return nil, err
	}

	return user, err
}
