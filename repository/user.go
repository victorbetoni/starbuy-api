package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func DownloadUser(username string, user *model.User) error {
	db := database.GrabDB()

	if err := db.Get(user, "SELECT * FROM users WHERE username=$1 LIMIT 1", username); err != nil {
		return err
	}

	return nil
}
