package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func InsertItem(item model.PostedItem) {
	db := database.GrabDB()

	var transaction = db.MustBegin()
	transaction.NamedExec("INSERT INTO products VALUES (:identifier, :title, :seller, :price, :stock, :category, :description)", &item.Item)
	transaction.Commit()

	for _, url := range item.Assets {
		transaction = db.MustBegin()
		transaction.MustExec("INSERT INTO product_images VALUES ($1, $2)", item.Item.Identifier, url)
		transaction.Commit()
	}
}

func DownloadItem(id string, includeUser bool, item *model.ItemWithAssets) error {
	db := database.GrabDB()

	var raw model.RawItem
	if err := db.Get(&raw, "SELECT * FROM products WHERE identifier=$1 LIMIT 1", id); err != nil {
		return err
	}

	var user model.User
	if includeUser {
		if err := DownloadUser(raw.Seller, false, &user); err != nil {
			return err
		}
	}

	var assets []string
	if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", id); err != nil {
		return err
	}

	literalItem := model.Item{
		Description: raw.Description,
		Title:       raw.Title,
		Identifier:  raw.Identifier,
		Price:       raw.Price,
		Stock:       raw.Stock,
		Category:    raw.Category,
	}

	if includeUser {
		literalItem.Seller = user
	}

	*item = model.ItemWithAssets{Item: literalItem, Assets: assets}
	return nil
}

func DownloadAllItems(items *[]model.ItemWithAssets, includeUser bool) error {
	db := database.GrabDB()

	var raws []model.RawItem
	if err := db.Select(&raws, "SELECT * FROM products"); err != nil {
		return err
	}

	for _, item := range raws {
		var assets []string

		if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", item.Identifier); err != nil {
			return err
		}

		var user model.User
		if includeUser {
			if err := db.Get(&user, "SELECT * FROM users WHERE username=$1", item.Seller); err != nil {
				return err
			}
		}

		item := model.Item{
			Description: item.Description,
			Title:       item.Title,
			Identifier:  item.Identifier,
			Price:       item.Price,
			Stock:       item.Stock,
			Seller:      user,
			Category:    item.Category,
		}

		if includeUser {
			item.Seller = user
		}

		*items = append(*items, model.ItemWithAssets{Item: item, Assets: assets})
	}
	return nil
}

func DownloadItemByCategory(category int, items *[]model.ItemWithAssets) error {
	db := database.GrabDB()

	var raws []model.RawItem
	if err := db.Select(&raws, "SELECT * FROM products WHERE category=$1", category); err != nil {
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
