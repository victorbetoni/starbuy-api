package repository

import (
	"starbuy/database"
	"starbuy/model"
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

	var raws []model.DatabaseItem
	if err := db.Select(&raws, "SELECT * FROM products"); err != nil {
		return err
	}

	for _, item := range raws {
		var assets []string

		if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", item.Identifier); err != nil {
			return err
		}

		var user model.User
		if err := db.Get(&user, "SELECT * FROM users WHERE username=$1", item.Seller); err != nil {
			return err
		}

		item := model.Item{
			Description: item.Description,
			Title:       item.Title,
			Identifier:  item.Identifier,
			Seller:      user,
			Price:       item.Price,
			Stock:       item.Stock,
			Category:    item.Category,
		}
		*items = append(*items, model.ItemWithAssets{Item: item, Assets: assets})
	}
	return nil
}

func DownloadItemByCategory(category int, items *[]model.ItemWithAssets) error {
	db := database.GrabDB()

	var raws []model.DatabaseItem
	if err := db.Select(&raws, "SELECT identifier FROM products WHERE category=$1"); err != nil {
		return err
	}

	for _, item := range raws {
		var assets []string

		if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", item.Identifier); err != nil {
			return err
		}

		var user model.User
		if err := db.Get(&user, "SELECT * FROM users WHERE username=$1", item.Seller); err != nil {
			return err
		}

		item := model.Item{
			Description: item.Description,
			Title:       item.Title,
			Identifier:  item.Identifier,
			Seller:      user,
			Price:       item.Price,
			Stock:       item.Stock,
			Category:    item.Category,
		}
		*items = append(*items, model.ItemWithAssets{Item: item, Assets: assets})
	}

	return nil
}
