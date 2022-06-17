package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"starbuy/database"
	"starbuy/model"
	"starbuy/security"
)

func InsertUser(user model.User, password string) error {
	db := database.GrabDB()

	testQueries := map[string]string{
		fmt.Sprintf("SELECT * FROM users WHERE username='%s'", user.Username): "Username já está em uso",
		fmt.Sprintf("SELECT * FROM users WHERE email='%s'", user.Email):       "Email já está em uso",
	}

	var found model.User
	for key, value := range testQueries {
		err := db.Get(&found, key)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if err == nil {
			return errors.New(value)
		}
	}

	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO users (username, email, name, registration, birthdate, seller, profile_picture, city) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		user.Username, user.Email, user.Name, user.Registration, user.Birthdate, user.Seller, user.ProfilePicture, user.City)
	if err := tx.Commit(); err != nil {
		return err
	}

	crypt, _ := security.Hash(password)
	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO login VALUES ($1,$2)", user.Username, string(crypt))
	if err := tx2.Commit(); err != nil {
		return err
	}

	return nil
}

func DownloadUser(username string, user *model.User) error {
	db := database.GrabDB()

	var temp model.User
	if err := db.Get(&temp, "SELECT * FROM users WHERE username=$1 LIMIT 1", username); err != nil {
		return err
	}

	fmt.Println(temp)

	*user = temp

	return nil
}
