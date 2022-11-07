package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func InsertItem(item model.PostedItem) error {
	db := database.GrabDB()

	var transaction = db.MustBegin()
	transaction.MustExec("INSERT INTO products VALUES ($1,$2,$3,$4,$5,$6,$7)", item.Item.Identifier, item.Item.Title, item.Item.Seller, item.Item.Price, item.Item.Stock, item.Item.Category, item.Item.Description)

	if err := transaction.Commit(); err != nil {
		return err
	}

	for _, url := range item.Assets {
		transaction = db.MustBegin()
		transaction.MustExec("INSERT INTO product_images VALUES ($1, $2)", item.Item.Identifier, url)
		transaction.Commit()
	}
	return nil
}

func DownloadItem(id string, item *model.ItemWithAssets) error {
	db := database.GrabDB()

	var raw model.RawItem
	if err := db.Get(&raw, "SELECT * FROM products WHERE identifier=$1 LIMIT 1", id); err != nil {
		return err
	}

	var user model.User
	if err := DownloadUser(raw.Seller, &user); err != nil {
		return err
	}

	var assets []string
	if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", id); err != nil {
		return err
	}

	literalItem := model.Item{
		Seller:      user,
		Description: raw.Description,
		Title:       raw.Title,
		Identifier:  raw.Identifier,
		Price:       raw.Price,
		Stock:       raw.Stock,
		Category:    raw.Category,
	}

	*item = model.ItemWithAssets{Item: literalItem, Assets: assets}
	return nil
}

func QueryItemsByName(query string, items *[]model.ItemWithAssets) error {
	db := database.GrabDB()

	var raws []model.RawItem
	if err := db.Select(&raws, "SELECT * FROM products WHERE title iLIKE '%"+query+"%'"); err != nil {
		return err
	}

	for _, item := range raws {
		var itemWithAssets model.ItemWithAssets
		convertRawItem(item, &itemWithAssets)
		*items = append(*items, itemWithAssets)
	}
	return nil
}

func DownloadAllItems(items *[]model.ItemWithAssets) error {
	db := database.GrabDB()

	var raws []model.RawItem
	if err := db.Select(&raws, "SELECT * FROM products"); err != nil {
		return err
	}

	for _, item := range raws {
		var itemWithAssets model.ItemWithAssets
		convertRawItem(item, &itemWithAssets)
		*items = append(*items, itemWithAssets)
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
		var itemWithAssets model.ItemWithAssets
		convertRawItem(item, &itemWithAssets)
		*items = append(*items, itemWithAssets)
	}

	return nil
}

func DownloadUserProducts(username string, items *[]model.ItemWithAssets) error {
	db := database.GrabDB()
	var raws []model.RawItem
	if err := db.Select(&raws, "SELECT * FROM products WHERE seller=$1", username); err != nil {
		return err
	}

	for _, item := range raws {
		var itemWithAssets model.ItemWithAssets
		convertRawItem(item, &itemWithAssets)
		*items = append(*items, itemWithAssets)
	}

	return nil
}

func convertRawItem(raw model.RawItem, itemWithAssets *model.ItemWithAssets) error {
	var assets []string
	db := database.GrabDB()

	if err := db.Select(&assets, "SELECT url FROM product_images WHERE product=$1", raw.Identifier); err != nil {
		return err
	}

	var user model.User
	if err := db.Get(&user, "SELECT * FROM users WHERE username=$1", raw.Seller); err != nil {
		return err
	}

	item := model.Item{
		Description: raw.Description,
		Title:       raw.Title,
		Identifier:  raw.Identifier,
		Price:       raw.Price,
		Stock:       raw.Stock,
		Category:    raw.Category,
		Seller:      user,
	}

	*itemWithAssets = model.ItemWithAssets{Item: item, Assets: assets}
	return nil
}
