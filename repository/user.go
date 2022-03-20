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
		return errors.New(value)
	}

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users VALUES (:username,:email,:name,:gender,:registration,:birthdate,:seller,:profile_picture,:city)", &user)
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

	if err := db.Get(user, "SELECT * FROM users WHERE username=$1 LIMIT 1", username); err != nil {
		return err
	}

	return nil
}

func DownloadCart(username string, items *[]model.CartItem) error {
	db := database.GrabDB()

	var stored []model.RawCartItem
	if err := db.Select(&stored, "SELECT * FROM shopping_cart WHERE holder=$1", username); err != nil {
		return err
	}

	for _, item := range stored {
		var casted model.CartItem
		var downloadedItem model.ItemWithAssets
		if err := DownloadItem(item.Item, &downloadedItem); err != nil {
			return err
		}
		*items = append(*items, casted)
	}

	return nil
}
