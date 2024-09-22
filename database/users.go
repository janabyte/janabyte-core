package database

//
//type User struct {
//	id        int
//	login     string
//	firstName string
//	lastName  string
//	email     string
//	phone     string
//	password  string
//}
//
//func (s *Storage) GetAllUsers() ([]*User, error) {
//	const op = "database.GetAllUsers"
//	userList := []*User{}
//	rows, err := s.db.Query(`
//			SELECT id,login,first_name,last_name,email
//			FROM users
//`)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %s", op, err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var user User
//		err = rows.Scan(&user.id, &user.login, &user.firstName, &user.lastName, &user.email)
//		if err != nil {
//			return nil, fmt.Errorf("%s: %s", op, err)
//		}
//		userList = append(userList, &user)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, fmt.Errorf("%s - %s", op, err)
//	}
//	return userList, nil
//}
//
//func (s *Storage) NewUser(login, firstName, lastName, email, phone, password string) error {
//	const op = "database.NewUser"
//	hashedPassword, err := repository.HashUserPassword(password)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	_, err = s.db.Exec(`
//			INSERT INTO users (login,first_name,last_name,email,phone,password)
//			VALUES ($1,$2,$3,$4,$5,$6);
//	`, login, firstName, lastName, email, phone, hashedPassword)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//
//	return nil
//}
//
//func (s *Storage) GetUserByLogin(login string) (*User, error) {
//	const op = "database.GetUserByLogin"
//	var user User
//	row, err := s.db.Query(`
//			SELECT id,login,first_name,last_name,email,phone
//			FROM users
//			WHERE login=$1;
//	`, login)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %s", op, err)
//	}
//	defer row.Close()
//	if row.Next() {
//		err = row.Scan(&user.id, &user.login, &user.firstName, &user.lastName, &user.email, &user.phone)
//		if err != nil {
//			return nil, fmt.Errorf("%s: %s", op, err)
//		}
//	}
//	return &user, nil
//}
//
//func (s *Storage) UpdateUser(id, login, firstName, lastName, email, phone, password string) error {
//	const op = "database.UpdateUser"
//	var currLogin, currFirstName, currLastName, currEmail, currPhone, currPassword string
//
//	rows, err := s.db.Query(`
//			SELECT login,first_name,last_name,email,phone,password
//			FROM users
//			WHERE login = $1
//	`, login)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	defer rows.Close()
//	if rows.Next() {
//		err = rows.Scan(&currLogin, &currFirstName, &currLastName, &currEmail, &currPhone, &currPassword)
//		if err != nil {
//			return fmt.Errorf("%s: %s", op, err)
//		}
//	}
//
//	hashedPassword := currPassword
//
//	if login == "" {
//		login = currLogin
//	}
//	if firstName == "" {
//		firstName = currFirstName
//	}
//	if lastName == "" {
//		lastName = currLastName
//	}
//	if email == "" {
//		email = currEmail
//	}
//	if phone == "" {
//		phone = currPhone
//	}
//
//	if password != "" {
//		hashedPassword, err = repository.HashUserPassword(password)
//		if err != nil {
//			return fmt.Errorf("%s: %s", op, err)
//		}
//	}
//	_, err = s.db.Exec(`
//			UPDATE users
//			SET = (?,?,?,?,?,?)
//			WHERE id = ?
//	`, login, firstName, lastName, email, phone, hashedPassword, id)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	return nil
//}
//
//func (s *Storage) DeleteUser(id int) error {
//	const op = "database.DeleteUser"
//	_, err := s.db.Exec(`
//			DELETE FROM users
//			WHERE id = $1
//	`, id)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	return nil
//}
//
//func (s *Storage) AuthenticateUser(login, password string) error {
//	const op = "database.AuthenticateUser"
//	var hashedPassword string
//	rows, err := s.db.Query(`
//			SELECT password from users
//			WHERE login = $1
//	`, login)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	defer rows.Close()
//	if rows.Next() {
//		err = rows.Scan(&hashedPassword)
//		if err != nil {
//			return fmt.Errorf("%s: %s", op, err)
//		}
//	}
//	err = repository.CheckPasswordHash(hashedPassword, password)
//	if err != nil {
//		return fmt.Errorf("%s: %s", op, err)
//	}
//	return nil
//}
