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

func GetItemReviews(c *gin.Context) error {
	queried := c.Param("item")

	var reviews []model.Review
	average, err := repository.QueryProductReviews(queried, &reviews)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews, Average: average})
	return nil
}

func GetUserReceivedReviews(c *gin.Context) error {
	username := c.Param("user")

	var reviews []model.Review
	average, err := repository.QueryUserReceivedReviews(username, &reviews)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews, Average: average})
	return nil
}

func GetUserReviews(c *gin.Context) error {
	username, _ := authorization.ExtractUser(c)

	var reviews []model.Review
	if err := repository.QueryUserReviews(username, &reviews); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "opaaa"})
		return nil
	}

	if req.Rating > 10 {
		req.Rating = 10
	}

	if req.Rating < 0 {
		req.Rating = 0
	}

	db := database.GrabDB()
	username, _ := authorization.ExtractUser(c)

	var foo model.RawReview
	if err := db.Get(&foo, "SELECT * FROM reviews WHERE username=$1 AND product=$2", username, req.Item); err == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Você já avaliou este produto."})
		return nil
	}

	if err := db.Get(&model.RawOrder{}, "SELECT * FROM orders WHERE holder=$1 AND product=$2", username, req.Item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Você não pode avaliar esse produto."})
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

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Avaliação adicionada com sucesso"})
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

func DeleteReview(c *gin.Context) error {
	item := c.Param("item")
	username, _ := authorization.ExtractUser(c)

	db := database.GrabDB()
	if err := db.Get(nil, "SELECT * FROM reviews WHERE username=$1 AND product=$2", username, item); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Nenhuma avaliação encontrada"})
			return nil
		}
	}

	repository.DeleteReview(username, item)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Avaliação excluída com sucesso"})
	return nil
}
