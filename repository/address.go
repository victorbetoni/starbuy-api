package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func DownloadAddress(identifier string, address *model.Address) error {
	db := database.GrabDB()

	var stored model.RawAddress
	if err := db.Select(&stored, "SELECT * FROM addresses WHERE identifier=$1", identifier); err != nil {
		return err
	}

	var user model.User
	if err := DownloadUser(stored.Holder, &user); err != nil {
		return err
	}

	address.Holder = &user
	address.CEP = stored.CEP
	address.Complement = stored.Complement
	address.Identifier = stored.Identifier
	address.Number = stored.Number

	return nil
}

func DownloadAddresses(username string, address *[]model.RawAddress) error {
	db := database.GrabDB()

	if err := db.Select(address, "SELECT * FROM addresses WHERE holder=$1", username); err != nil {
		return err
	}

	return nil
}
