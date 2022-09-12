package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func DownloadPurchases(username string, orders *[]model.Order) error {
	db := database.GrabDB()

	var raw []model.RawOrder
	if err := db.Select(&raw, "SELECT * FROM orders WHERE holder = $1", username); err != nil {
		return err
	}

	for _, item := range raw {
		var order model.Order
		if err := DownloadPurchase(item.Identifier, &order); err != nil {
			return err
		}
		*orders = append(*orders, order)
	}

	return nil
}

func DownloadOrders(seller string, orders *[]model.OrderWithItem) error {
	db := database.GrabDB()

	var raw []model.RawOrder
	if err := db.Select(&raw, "SELECT * FROM orders WHERE seller = $1", seller); err != nil {
		return err
	}

	for _, item := range raw {
		var order model.Order
		var product model.ItemWithAssets
		if err := DownloadPurchase(item.Identifier, &order); err != nil {
			return err
		}
		if err := DownloadItem(item.Item, &product); err != nil {
			return err
		}
		*orders = append(*orders, model.OrderWithItem{Order: order, Item: product})
	}
	return nil
}

func DownloadPurchase(identifier string, order *model.Order) error {
	db := database.GrabDB()

	var raw model.RawOrder
	if err := db.Get(&raw, "SELECT * FROM orders WHERE identifier = $1", identifier); err != nil {
		return err
	}

	var customer model.User
	var item model.ItemWithAssets

	if err := DownloadUser(raw.Customer, &customer); err != nil {
		return err
	}

	if err := DownloadItem(raw.Item, &item); err != nil {
		return err
	}

	order.Customer = customer
	order.Identifier = identifier
	order.Quantity = raw.Quantity
	order.Seller = item.Item.Seller
	order.Price = raw.Price
	order.Item = item

	return nil
}

func InsertPurchase(purchase model.Order) error {
	db := database.GrabDB()

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6)", purchase.Identifier, purchase.Customer.Username, purchase.Seller.Username, purchase.Item.Item.Identifier, purchase.Quantity, purchase.Price)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
