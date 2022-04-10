package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetReviews(c *gin.Context) error {
	queried := c.Param("user")

	var reviews []model.Review
	if err := repository.QueryUserReviews(queried, &reviews); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, reviews)
	return nil
}

func GetReview(c *gin.Context) error {
	queried := c.Param("id")

	var review model.Review
	if err := repository.DownloadReview(queried, &review); err != nil {
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
	if err := db.Get(nil, "SELECT * FROM purchase_log WHERE holder=$1 AND product=$2", username, req.Item); err != nil && err == sql.ErrNoRows {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "no order found"})
			return nil
		}
		return err
	}

	final := model.RawReview{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		User:       username,
		Item:       req.Item,
		Message:    req.Message,
		Rate:       req.Rating,
	}

	repository.InsertReview(final)

	c.JSON(http.StatusOK, final)
	return nil
}

func PutReview(c *gin.Context) error {
	type Request struct {
		Review  string `json:"id"`
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
	if err := repository.DownloadReview(req.Review, &review); err != nil {
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

	final := model.RawReview{Identifier: review.Identifier, User: username, Message: req.Message, Rate: req.Rate, Item: review.Item.Item.Identifier}

	repository.UpdateReview(final)
	c.JSON(http.StatusOK, final)
	return nil
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {

}
