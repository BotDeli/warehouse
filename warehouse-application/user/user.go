package user

import (
	"database/sql"
	"errors"
)

type User struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

func (user *User) RegisterUser(db *sql.DB) error {
	action := `INSERT INTO users(login, hash_password) VALUES ($1,$2)`
	login := user.Login
	hashPassword := Hashing(user.Password)
	if _, err := db.Exec(action, login, hashPassword); err != nil {
		return err
	}
	return nil
}

func (user *User) LoginUser(db *sql.DB) error {
	var (
		correctLogin        string
		correctHashPassword string
	)

	err := db.QueryRow(`SELECT login, hash_password FROM users`).Scan(&correctLogin, &correctHashPassword)
	if err != nil {
		return err
	}

	if correctLogin == user.Login && correctHashPassword == Hashing(user.Password) {
		return nil
	}

	return errors.New("Wrong password or login!")
}
