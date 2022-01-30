package store

import (
	"go_project/src/users/models"
)

type UserRepository struct {

}

//var tableName = "users"

func (p *UserRepository) GetUser(e *models.User) error {
	return conn.db.QueryRow("SELECT users.id username, email, is_active, roles.name FROM users join roles on users.role_id=roles.id WHERE users.id=$1",
		e.ID).Scan(&e.ID, &e.Username, &e.Email, &e.IsActive, &e.Role)
}

func (p *UserRepository) UpdateUser(e *models.User) error {
	_, err :=
		conn.db.Exec("UPDATE users SET username=$1, email=$3 WHERE id=$3",
			e.Username, e.Email, e.ID)
	return err
}

func (p *UserRepository) DeleteUser(e *models.User) error {
	_, err := conn.db.Exec("DELETE FROM users WHERE id=$1", e.ID)
	return err
}

func (p *UserRepository) IsUserExistByEmail(e *models.User) (error, bool)  {
	var checkUserExist *bool

	err := conn.db.QueryRow("select exists(select email from users where email=$1)",
		e.Email).Scan(&checkUserExist)

	return err, *checkUserExist
}

func (p *UserRepository) GetUserPassword(e *models.User) (error, string)  {
	var password *string
	err := conn.db.QueryRow("select password from users where email=$1", e.Email).Scan(&password)

	return err, *password
}

func (p *UserRepository) GetUsernameAndEmail (e *models.User) error  {
	return conn.db.QueryRow("SELECT username, email FROM users WHERE email=$1",
		e.Email).Scan(&e.Username, &e.Email)
}

func (p *UserRepository) CreateUser(e *models.User) error {
	error1 := conn.db.QueryRow(
		"INSERT INTO users(username, email, password, role_id) VALUES($1, $2, $3, $4) RETURNING id, (select username from roles where roles.id = $4)", e.Username,
		e.Email, e.Password, e.RoleId).Scan(&e.ID, &e.Role)
	if error1 != nil {
		return error1
	}
	return nil
}


func (p *UserRepository) GetUsers (start, count int) ([]models.User, error) {
	rows, err := conn.db.Query(
		"SELECT users.id, username, email, is_active, roles.name FROM users JOIN roles on users.role_id = roles.id LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var p models.User
		if err := rows.Scan(&p.ID, &p.Username, &p.Email, &p.IsActive, &p.Role); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}


