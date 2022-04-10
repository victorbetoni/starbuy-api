package repository

import (
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
	tx.MustExec("INSERT INTO shopping_cart VALUES ($1,$2,$3)", item.Holder, item.Item, item.Quantity)
	if err := tx.Commit(); err != nil {
		tx2 := db.MustBegin()
		tx2.MustExec("UPDATE shopping_cart SET quantity=quantity+$1 WHERE product=$2 AND holder=$3", item.Quantity, item.Item, item.Holder)
		if erro := tx2.Commit(); erro != nil {
			return erro
		}
	}

	return nil
}
