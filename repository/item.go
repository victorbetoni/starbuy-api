package repository

import (
	"authentication-service/database"
	"authentication-service/model"
	"fmt"
)

func DownloadItem(id string, item *model.ItemWithAssets) error {
	db := database.GrabDB()

	if err := db.Get(&item.Item, "SELECT * FROM products WHERE identifier=$1 LIMIT 1", id); err != nil {
		fmt.Println("DEBUG 1")
		return err
	}

	if err := db.Select(&item.Assets, "SELECT url FROM product_images WHERE product=$1", id); err != nil {
		fmt.Println("DEBUG 2")
		return err
	}

	return nil
}
