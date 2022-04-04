package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetReviews(c *gin.Context) {
	queried := c.Param("user")

	var reviews []model.Review
	if err := repository.QueryUserReviews(queried, &reviews); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetReview(c *gin.Context) {
	queried := c.Param("id")

	var review model.Review
	if err := repository.DownloadReview(queried, &review); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, review)
}

func PostReview(c *gin.Context) {

	rate, err := strconv.Atoi(c.PostForm("rate"))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid rating"))
		return
	}

	item, message := c.PostForm("item"), c.PostForm("message")

	if rate > 10 {
		rate = 10
	}

	if rate < 0 {
		rate = 0
	}

	username, erro := authorization.ExtractUser(c)
	if erro != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	db := database.GrabDB()
	if err := db.Get(nil, "SELECT * FROM purchase_log WHERE holder=$1 AND product=$2", username, item); err != nil && err == sql.ErrNoRows {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	final := model.RawReview{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		User:       username,
		Item:       item,
		Message:    message,
		Rate:       rate,
	}

	repository.InsertReview(final)

	c.JSON(http.StatusOK, final)
}

func PutReview(c *gin.Context) {
	type Request struct {
		Review  string `json:"id"`
		Rate    int    `json:"rate"`
		Message string `json:"message"`
	}

	req := Request{}
	username, erro := authorization.ExtractUser(c)
	if erro != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var review model.Review
	if err := repository.DownloadReview(req.Review, &review); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if review.User.Username != username {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	final := model.RawReview{Identifier: review.Identifier, User: username, Message: req.Message, Rate: req.Rate, Item: review.Item.Item.Identifier}

	repository.UpdateReview(final)
	c.JSON(http.StatusOK, final)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {

}
