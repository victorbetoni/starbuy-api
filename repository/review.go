package repository

import (
	"fmt"
	"starbuy/database"
	"starbuy/model"
)

func QueryUserReceivedReviews(username string, reviews *[]model.Review) (float64, error) {
	db := database.GrabDB()

	count, sum := 0, 0
	var raw []model.RawReview
	if err := db.Select(&raw, "SELECT R.product, R.username, R.msg, R.rate  FROM reviews R INNER JOIN products P ON R.product = P.identifier AND P.seller=$1", username); err != nil {
		return 0, err
	}

	for _, review := range raw {
		count++
		var rev model.Review
		err := convertRawReview(review, &rev)
		if err != nil {
			return 0, err
		}
		sum += review.Rate
		*reviews = append(*reviews, rev)
	}

	return (float64(sum) / float64(count)), nil

}

func QueryUserReviews(username string, reviews *[]model.Review) (float64, error) {
	db := database.GrabDB()

	count, sum := 0, 0
	var raw []model.RawReview
	if err := db.Select(&raw, "SELECT * FROM reviews WHERE username=$1", username); err != nil {
		return 0, err
	}

	for _, review := range raw {
		count++
		var rev model.Review
		err := convertRawReview(review, &rev)
		if err != nil {
			return 0, err
		}
		sum += review.Rate
		*reviews = append(*reviews, rev)
	}

	return (float64(sum) / float64(count)), nil

}

func QueryProductReviews(product string, reviews *[]model.Review) (float64, error) {
	db := database.GrabDB()

	count, sum := 0, 0
	var raw []model.RawReview
	if err := db.Select(&raw, "SELECT * FROM reviews WHERE product=$1", product); err != nil {
		return 0, err
	}

	for _, review := range raw {
		count++
		var rev model.Review
		err := convertRawReview(review, &rev)
		if err != nil {
			return 0, err
		}
		*reviews = append(*reviews, rev)
		sum += review.Rate
	}

	return (float64(count) / float64(sum)), nil
}

func DownloadReview(user string, item string, review *model.Review) error {
	db := database.GrabDB()

	var raw model.RawReview
	if err := db.Get(&raw, "SELECT * FROM reviews WHERE username=$1 AND product=$2", user, item); err != nil {
		return err
	}

	var holder model.User
	if err := DownloadUser(user, &holder); err != nil {
		return err
	}

	var reviewItem model.ItemWithAssets
	if err := DownloadItem(item, &reviewItem); err != nil {
		return err
	}

	*review = model.Review{User: holder, Item: &reviewItem, Message: raw.Message, Rate: raw.Rate}

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

func convertRawReview(raw model.RawReview, review *model.Review) error {
	var rev model.Review

	if err := DownloadReview(raw.User, raw.Item, &rev); err != nil {
		return err
	}

	fmt.Println(rev)

	*review = rev

	return nil
}
