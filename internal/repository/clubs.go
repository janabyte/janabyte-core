package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/internal/http/model"
)

type ClubRepository struct {
	db *sql.DB
}

func NewClubRepository(db *sql.DB) *ClubRepository {
	return &ClubRepository{db: db}
}

func (r *ClubRepository) CreateClub(club *model.Club) (id int, err error) {
	const op = "repository.CreateClub"
	err = r.db.QueryRow(`
			INSERT INTO clubs(name,description,address,work_time_start,work_time_end,x_size,y_size,user_id)
			VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8)
			RETURNING id
	`, club.Name, club.Description, club.Address, club.WorkTimeStart, club.WorkTimeEnd, club.XSize, club.YSize, club.UserId).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r *ClubRepository) GetClubById(id int) (*model.Club, error) {
	const op = "repository.GetClubById"
	club := &model.Club{}
	row := r.db.QueryRow(`
			SELECT 	name,description,address,work_time_start,work_time_end,x_size,y_size,user_id
			FROM clubs
			WHERe id = $1
	`, id)
	err := row.Scan(&club.Name, &club.Description, &club.Address, &club.WorkTimeStart, &club.WorkTimeEnd, &club.XSize, &club.YSize, &club.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, err)
		} else if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}
	return club, nil
}

func (r *ClubRepository) GetAllClubs() ([]*model.Club, error) {
	const op = "repository.GetAllClubs"
	clubList := []*model.Club{}
	rows, err := r.db.Query(`
			SELECT *
			FROM clubs
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		club := &model.Club{}
		err = rows.Scan(&club.Id, &club.Name, &club.Description, &club.Address, &club.WorkTimeStart, &club.WorkTimeEnd, &club.XSize, &club.YSize, &club.UserId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		clubList = append(clubList, club)
	}
	return clubList, nil
}

func (r *ClubRepository) DeleteClubById(id int) error {
	const op = "repository.DeleteClubById"
	_, err := r.db.Exec(`
			DELETE FROM clubs
			WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *ClubRepository) UpdateClub(id int, club *model.Club) error {
	const op = "repository.UpdateClub"
	_, err := r.db.Exec(`
			UPDATE clubs
			SET name = $1,description = $2,address = $3, work_time_start = $4,work_time_end = $5, x_size = $6, y_size = $7, user_id = $8
			WHERE id = $9
	`, club.Name, club.Description, club.Address, club.WorkTimeStart, club.WorkTimeEnd, club.XSize, club.YSize, club.UserId, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
