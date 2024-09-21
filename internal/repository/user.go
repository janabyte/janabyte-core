package repository

import (
	"database/sql"
	"fmt"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (repository *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
    rows, err := repository.db.Query("SELECT * FROM users WHERE phone = ?", phone)
    if err != nil {
        return nil, err
    }

    user := new(model.User)

    for rows.Next() {
        user, err = scanRowIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if user.Id == 0 {
        return nil, fmt.Errorf("User not found")
    }

    return user, nil
}

func (repository *UserRepository) CreateUser(user model.User) error {
    _, err := repository.db.Exec("INSERT INTO users (name, email, phone, password) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Phone, user.Password)
    if err != nil {
        return err
    }

    return nil
}

func scanRowIntoUser(rows *sql.Rows) (*model.User, error) {
    user := new(model.User)

    err := rows.Scan(
        &user.Id,
        &user.Name,
        &user.Email,
        &user.Phone,
        &user.Password,
    )
    if err != nil {
        return nil, err
    }

    return user, err
}
