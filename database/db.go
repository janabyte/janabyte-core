package database

import (
	"database/sql"
	"fmt"
	"github.com/aidosgal/janabyte/janabyte-core/config"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "database.New"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PublicHost,
		cfg.Port,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("%s %s", op, err)
	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users(
			    id SERIAL PRIMARY KEY,
			    login varchar(255) not null UNIQUE,
			    first_name varchar(255) not null,
			    last_name varchar(255) not null,
			    email varchar(255) not null UNIQUE,
			    phone varchar(255) not null UNIQUE,
			    password varchar(255) not null
			);
`)
	if err != nil {
		log.Fatalf("Error when creating Users: %s %s", op, err)

	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS owners(
			    id SERIAL PRIMARY KEY,
			    email varchar(255) not null UNIQUE,
			    first_name varchar(255) not null,
			    last_name varchar(255) not null,
			    password varchar(255) not null
			)
	`)
	if err != nil {
		log.Fatalf("Error when creating owners: %s %s", op, err)
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS clubs(
			    id SERIAL PRIMARY KEY,
			    name varchar(255) not null,
			    description text,
			    address varchar(255) not null,
			    work_time_start time not null,
			    work_time_end time not null,
			    x_size int,
			    y_size int,
			    owner_id int,
			    FOREIGN KEY(owner_id) REFERENCES owners(id)
			);
`)
	if err != nil {
		log.Fatalf("Error when creating Clubs: %s %s", op, err)
	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS admins(
			    id SERIAL PRIMARY KEY,
			    first_name varchar(255) not null,
			    last_name varchar(255) not null,
			    email varchar(255) not null UNIQUE,
			    phone varchar(255) not null UNIQUE,
			    password varchar(255) not null
			);
`)
	if err != nil {
		log.Fatalf("Error when creating Admins: %s %s", op, err)
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS instances(
			    id SERIAL PRIMARY KEY,
			    name varchar(255) not null,
			    icon_url varchar(255) not null
			)
	`)
	if err != nil {
		log.Fatalf("Error when creating instances: %s %s", op, err)
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS clubs_computers(
			    id SERIAL PRIMARY KEY,
			    computer_number int not null,
			    is_near_to_next bool not null,
			    is_near_to_prev bool not null,
			    gpu varchar(255) not null,
			    cpu varchar(255) not null,
			    ram varchar(255) not null,
			    y_pos int,
			    x_pos int,
			    instance_id int,
			    club_id int not null,
			    FOREIGN KEY (club_id) REFERENCES clubs(id),
			    FOREIGN KEY(instance_id) REFERENCES instances(id)
			);
`)
	if err != nil {
		log.Fatalf("Error when creating clubs_computers: %s %s", op, err)
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS appointments(
			    id SERIAL PRIMARY KEY,
			    startDate timestamp,
			    endDate timestamp,
			    isActive bool not null,
			    club_computer_id int not null,
			    user_id int not null,
			    FOREIGN KEY(club_computer_id) REFERENCES clubs_computers(id),
			    FOREIGN KEY(user_id) REFERENCES users(id)
			);
`)
	if err != nil {
		log.Fatalf("Error when creating appointments: %s %s", op, err)
	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS user_main_transaction(
			    id SERIAL PRIMARY KEY,
			    amount int not null,
			    transactionType varchar(255) not null,
			    user_id int not null,
			    club_id int not null,
			    FOREIGN KEY(user_id) REFERENCES users(id),
			    FOREIGN KEY(club_id) REFERENCES clubs(id)
			);
`)
	if err != nil {
		log.Fatalf("Error when creating user_main_transaction: %s %s", op, err)
	}
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS user_bonus_transaction(
			    id SERIAL PRIMARY KEY,
			    amount int not null,
			    transactionType varchar(255) not null,
			    user_id int not null,
			    club_id int not null,
			    FOREIGN KEY(user_id) REFERENCES users(id),
			    FOREIGN KEY(club_id) REFERENCES clubs(id)
			);
`)
	if err != nil {
		log.Fatalf("Error when creating user_bonus_transaction: %s %s", op, err)

	}
	return &Storage{DB: db}, nil
}
