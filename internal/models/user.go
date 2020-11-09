package models

import (
	"database/sql"
)

type UsersRepository interface {
	AllUsers() ([]*User, error)
	CreateUser(User) (int, error)
	UserById(int) (*User, error)
	UpdateUser(User) error
	DeleteUser(int) error
}

type User struct {
	Id   int    `json:"user_id,omitempty"`
	Name string `json:"name"`
}

type userDB struct {
	*sql.DB
}

func (db *userDB) AllUsers() ([]*User, error) {
	rows, err := db.Query("select user_id, name from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (db *userDB) CreateUser(u User) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO users (name)
	VALUES ($1)
	RETURNING user_id`

	row := db.QueryRow(sqlInsert, u.Name)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *userDB) UserById(id int) (*User, error) {
	query := `select user_id, name from users where user.user_id = $1`

	row := db.QueryRow(query, id)
	user := new(User)
	err := row.Scan(&user.Id, &user.Name)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *userDB) UpdateUser(u User) error {
	sqlUpdate := `UPDATE users SET name=$1 WHERE id=$2`

	_, err :=
		db.Exec(sqlUpdate,
			u.Name, u.Id)

	return err
}

func (db *userDB) DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM user WHERE id=$1", id)

	return err
}
