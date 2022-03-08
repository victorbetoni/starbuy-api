package repository

import (
	"authentication-service/database"
	"authentication-service/model"
)

func DownloadItem(id string, item *model.ItemWithAssets) error {
	db := database.GrabDB()

	if err := db.Get(&item.Item, "SELECT * FROM products WHERE identifier=$1 LIMIT 1", id); err != nil {
		return err
	}

	if err := db.Select(&item.Assets, "SELECT url FROM product_images WHERE product=$1", id); err != nil {
		return err
	}

	return nil
}

func DownloadAllItems(items *[]model.ItemWithAssets) error {
	db := database.GrabDB()

	var ids []string
	if err := db.Select(&ids, "SELECT identifier FROM products"); err != nil {
		return err
	}

	for _, id := range ids {
		var item model.ItemWithAssets
		if err := db.Get(&item.Item, "SELECT * FROM products WHERE identifier=$1 LIMIT 1", id); err != nil {
			return err
		}

		if err := db.Select(&item.Assets, "SELECT * FROM product_images WHERE product=$1", id); err != nil {
			return err
		}
		*items = append(*items, item)
	}

	return nil
}
