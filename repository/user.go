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

func DownloadCart(username string, items *[]model.CartItem) error {
	db := database.GrabDB()

	var stored []model.RawCartItem
	if err := db.Select(&stored, "SELECT * FROM shopping_cart WHERE holder=$1", username); err != nil {
		return err
	}

	for _, item := range stored {
		var casted model.CartItem
		var downloadedItem model.ItemWithAssets
		if err := DownloadItem(item.Item, false, &downloadedItem); err != nil {
			return err
		}
		*items = append(*items, casted)
	}

	return nil
}
