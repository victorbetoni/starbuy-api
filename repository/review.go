package repository

import (
	"starbuy/database"
	"starbuy/model"
)

func QueryReviews(product string, reviews *[]model.RawReview) error {
	db := database.GrabDB()

	if err := db.Select(reviews, "SELECT * FROM reviews WHERE product=$1", product); err != nil {
		return err
	}

	return nil
}

func InsertReview(review model.RawReview) error {
	db := database.GrabDB()

	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO reviews VALUES ($1,$2,$3,$4)", review.Item, review.User, review.Message, review.Rate)
	if err := tx2.Commit(); err != nil {
		return err
	}

	return nil
}
