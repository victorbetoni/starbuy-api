package repository

import (
	"fmt"
	"starbuy/database"
	"starbuy/model"
)

func QueryUserReviews(username string, reviews *[]model.Review) (int, error) {
	db := database.GrabDB()

	var count int
	var sum int
	var raw []model.RawReview
	if err := db.Select(&raw, "SELECT * FROM reviews WHERE username=$1", username); err != nil {
		return 0, err
	}

	for _, review := range raw {
		count++
		err, review := convertRawReview(review)
		if err != nil {
			return 0, err
		}
		sum += review.Rate
		*reviews = append(*reviews, review)
	}

	return (sum / count), nil

}

func QueryProductReviews(product string, reviews *[]model.Review) error {
	db := database.GrabDB()

	var raw []model.RawReview
	if err := db.Select(&raw, "SELECT * FROM reviews WHERE product=$1", product); err != nil {
		return err
	}

	for _, review := range raw {
		fmt.Println("Encontrou: " + review.User)
		err, review := convertRawReview(review)
		if err != nil {
			return err
		}
		*reviews = append(*reviews, review)
	}

	return nil
}

func DownloadReview(identifier string, review *model.Review) error {
	db := database.GrabDB()

	var raw model.RawReview
	if err := db.Get(&raw, "SELECT * FROM reviews WHERE identifier=$1", identifier); err != nil {
		return err
	}

	err, rev := convertRawReview(raw)
	if err != nil {
		return err
	}

	*review = rev

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

func DeleteReview(identifier string) error {
	db := database.GrabDB()

	tx2 := db.MustBegin()
	tx2.MustExec("DELETE FROM reviews WHERE identifier=$1", identifier)
	if err := tx2.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateReview(raw model.RawReview) error {
	db := database.GrabDB()

	tx := db.MustBegin()
	tx.MustExec("UPDATE reviews SET msg=$1, rate=$2 WHERE username=$3 AND product=$4", raw.Message, raw.Rate, raw.User, raw.Item)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func convertRawReview(raw model.RawReview) (error, model.Review) {
	var user model.User
	if err := DownloadUser(raw.User, &user); err != nil {
		return err, model.Review{}
	}

	var item model.ItemWithAssets
	if err := DownloadItem(raw.User, &item); err != nil {
		return err, model.Review{}
	}

	return nil, model.Review{User: user, Item: item, Message: raw.Message, Rate: raw.Rate}
}
