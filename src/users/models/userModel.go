package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//var tableName = "users"

func (p *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT users.id username, email, is_active, roles.name FROM users join roles on users.role_id=roles.id WHERE users.id=$1",
		p.ID).Scan(&p.ID, &p.Username, &p.Email, &p.IsActive, &p.Role)
}

func (p *User) UpdateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET username=$1, email=$3 WHERE id=$3",
			p.Username, p.Email, p.ID)
	return err
}

func (p *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", p.ID)
	return err
}

func (p *User) LoginUser(db *sql.DB) error {
	var checkUserExist *bool
	var password *string
	err := db.QueryRow("select exists(select email from users where email=$1)",
		p.Email).Scan(&checkUserExist)
	if *checkUserExist {
		err := db.QueryRow("select password from users where email=$1", p.Email).Scan(&password)
		if err != nil {
			return err
		}
		byteHash := []byte(*password)
		bytePass := []byte(p.Password)
		result := bcrypt.CompareHashAndPassword(byteHash, bytePass)
		if result != nil {
			return errors.New("Login or password is not correct")
		}
		return db.QueryRow("SELECT username, email FROM users WHERE email=$1",
			p.Email).Scan(&p.Username, &p.Email)
	}
	if err != nil {
		return err
	}
	return errors.New("Login or password is not correct")
}

func (p *User) UserRegister(db *sql.DB) error {
	var checkUserExist *bool
	err := db.QueryRow("select exists(select email from users where email=$1)",
		p.Email).Scan(&checkUserExist)
	if *checkUserExist {
		return errors.New("A user is already registered to this mail")
	}
	if err != nil {
		return err
	}
	bytePassword := []byte(p.Password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	password := string(hash)
	p.Password = password
	error1 := db.QueryRow(
		"INSERT INTO users(username, email, password, role_id) VALUES($1, $2, $3, $4) RETURNING id, (select username from roles where roles.id = $4)", p.Username,
		p.Email, p.Password, p.RoleId).Scan(&p.ID, &p.Role)
	p.Password = ""
	if error1 != nil {
		return error1
	}
	return nil
}

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	rows, err := db.Query(
		"SELECT users.id, username, email, is_active, roles.name FROM users JOIN roles on users.role_id = roles.id LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var p User
		if err := rows.Scan(&p.ID, &p.Username, &p.Email, &p.IsActive, &p.Role); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}
