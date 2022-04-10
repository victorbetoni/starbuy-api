package repository

import (
	"database/sql"
	"starbuy/database"
	"starbuy/model"
)

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

func InsertCartItem(item model.RawCartItem) error {
	db := database.GrabDB()

	tx := db.MustBegin()

	var recorded model.RawCartItem
	if err := db.Get(&recorded, "SELECT * FROM shopping_cart WHERE holder=$1 AND product=$2", item.Holder, item.Item); err != nil && err == sql.ErrNoRows {
		tx.MustExec("INSERT INTO shopping_cart VALUES ($1,$2,$3)", item.Holder, item.Item, item.Quantity)
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}

	tx2 := db.MustBegin()
	tx2.Exec("UPDATE shopping_cart SET quantity=$1 WHERE holder=$2 AND product=$3", item.Quantity+recorded.Quantity, item.Holder, item.Item)
	if err := tx2.Commit(); err != nil {
		return err
	}

	return nil
}
