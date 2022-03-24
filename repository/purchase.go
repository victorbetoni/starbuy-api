package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func DownloadPurchases(username string, purchases *[]model.Purchase) error {
	db := database.GrabDB()

	if err := db.Select(purchases, "SELECT * FROM purchases WHERE seller = $1", username); err != nil {
		return err
	}

	return nil
}

func InsertPurchase(purchase model.Purchase) error {
	db := database.GrabDB()

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO purchase_log VALUES ($1,$2,$3,$4,$5)", purchase.Identifier, purchase.Customer.Username, purchase.Seller.Username, purchase.Item.Item.Identifier, purchase.Quantity, purchase.Price)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
