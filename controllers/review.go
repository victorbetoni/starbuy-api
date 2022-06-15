package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) error {
	queried := c.Param("user")

	var reviews []model.Review
	var average int
	if loc, err := repository.QueryUserReviews(queried, &reviews); err != nil {
		if err == sql.ErrNoRows {
			average = loc
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average int            `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews, Average: average})
	return nil
}

func GetReview(c *gin.Context) error {
	user := c.Query("user")
	product := c.Query("product")

	var review model.Review
	if err := repository.DownloadReview(user, product, &review); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, review)
	return nil
}

func PostReview(c *gin.Context) error {

	type Request struct {
		Rating  int    `json:"rate"`
		Item    string `json:"item"`
		Message string `json:"message"`
	}

	req := Request{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	if req.Rating > 10 {
		req.Rating = 10
	}

	if req.Rating < 0 {
		req.Rating = 0
	}

	username, _ := authorization.ExtractUser(c)

	db := database.GrabDB()
	if err := db.Get(&model.RawOrder{}, "SELECT * FROM orders WHERE holder=$1 AND product=$2", username, req.Item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "no order found"})
			return nil
		}
		return err
	}

	final := model.RawReview{
		User:    username,
		Item:    req.Item,
		Message: req.Message,
		Rate:    req.Rating,
	}

	repository.InsertReview(final)

	c.JSON(http.StatusOK, final)
	return nil
}

func PutReview(c *gin.Context) error {
	type Request struct {
		User    string `json:"user"`
		Item    string `json:"item"`
		Rate    int    `json:"rate"`
		Message string `json:"message"`
	}

	req := Request{}
	username, _ := authorization.ExtractUser(c)

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	var review model.Review
	if err := repository.DownloadReview(req.User, req.Item, &review); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	if review.User.Username != username {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		return nil
	}

	final := model.RawReview{User: username, Message: req.Message, Rate: req.Rate, Item: review.Item.Item.Identifier}

	repository.UpdateReview(final)
	c.JSON(http.StatusOK, final)
	return nil
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {

}
