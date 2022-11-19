package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func DeleteAddress(address string, username string) error {
	db := database.GrabDB()
	tx := db.MustBegin()

	tx.MustExec("DELETE FROM addresses WHERE identifier=$1 AND holder=$2", address, username)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func DownloadAddress(identifier string, address *model.Address) error {
	db := database.GrabDB()

	var stored model.RawAddress
	if err := db.Get(&stored, "SELECT * FROM addresses WHERE identifier=$1", identifier); err != nil {
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

func InsertAddress(address model.RawAddress) error {
	db := database.GrabDB()

	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO addresses VALUES ($1,$2,$3,$4,$5,$6)", address.Identifier, address.Name, address.Holder, address.CEP, address.Number, address.Complement)
	if err := tx2.Commit(); err != nil {
		return err
	}

	return nil
}
