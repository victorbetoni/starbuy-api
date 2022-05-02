package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func InsertNotification(notification model.RawNotification) error {
	db := database.GrabDB()

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO orders VALUES ($1,$2,$3,$4)", notification.Identifier, notification.User, notification.Text, notification.SentIn)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func DownloadNotifications(user string, notifications *[]model.RawNotification) error {
	db := database.GrabDB()

	if err := db.Select(notifications, "SELECT * FROM notification WHERE holder = $1", user); err != nil {
		return err
	}

	return nil
}

func DownloadNotification(id string, notification *model.Notification) error {
	db := database.GrabDB()

	var raw model.RawNotification
	if err := db.Get(&raw, "SELECT * FROM notification WHERE identifier = $1", id); err != nil {
		return err
	}

	var holder model.User
	if err := DownloadUser(raw.User, &holder); err != nil {
		return err
	}

	*notification = model.Notification{
		Identifier: raw.Identifier,
		User:       &holder,
		Text:       raw.Text,
		SentIn:     raw.SentIn,
	}

	return nil
}
