package models

import (
	"errors"

	"dkds.com/rest-api/db"
	"dkds.com/rest-api/security"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type userData struct {
	id           int64
	email        string
	passwordHash string
	passwordSalt string
}

func (user *User) Save() error {
	passwordSalt := security.GenerateSalt()
	passwordHash, err := security.HashPassword(user.Password, passwordSalt)
	if err != nil {
		return errors.New("Could not save user, " + err.Error())
	}

	query := `
	INSERT INTO user
	(email, password_hash, password_salt)
	VALUES
	(?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Could not save user, " + err.Error())
	}

	result, err := statement.Exec(user.Email, passwordHash, passwordSalt)
	if err != nil {
		return errors.New("Could not save user, " + err.Error())
	}
	defer statement.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return errors.New("Could not retrieve last saved ID, " + err.Error())
	}

	user.ID = id

	return nil
}

func (user *User) ValidateCredentials() error {
	query := `
	SELECT id, email, password_hash, password_salt
	FROM user
	WHERE email = ?
	`
	row := db.DB.QueryRow(query, user.Email)
	var userData userData
	err := row.Scan(
		&userData.id,
		&userData.email,
		&userData.passwordHash,
		&userData.passwordSalt,
	)
	if err != nil {
		return errors.New("Could not retrieve the user" + err.Error())
	}

	valid := security.CheckPasswordHash(user.Password, userData.passwordHash, userData.passwordSalt)
	if !valid {
		return errors.New("Could not validate the credentials")
	}
	return nil
}
