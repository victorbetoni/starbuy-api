package controllers

import (
	"database/sql"
	"errors"
	"math"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"

	"github.com/gin-gonic/gin"
)

func GetItemReviews(c *gin.Context) (int, error) {
	queried := c.Param("item")

	var reviews []model.Review
	average, err := repository.QueryProductReviews(queried, &reviews)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews, Average: average})
	return 0, nil
}

func GetUserReceivedReviews(c *gin.Context) (int, error) {
	username := c.Param("user")

	var reviews []model.Review
	average, err := repository.QueryUserReceivedReviews(username, &reviews)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews, Average: average})
	return 0, nil
}

func GetUserReviews(c *gin.Context) (int, error) {
	username, _ := authorization.ExtractUser(c)

	var reviews []model.Review
	if err := repository.QueryUserReviews(username, &reviews); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	type ItemReviews struct {
		Reviews []model.Review `json:"reviews"`
		Average float64        `json:"average"`
	}

	c.JSON(http.StatusOK, ItemReviews{Reviews: reviews})
	return 0, nil
}

func GetReview(c *gin.Context) (int, error) {
	user := c.Query("user")
	product := c.Query("product")

	var review model.Review
	if err := repository.DownloadReview(user, product, &review); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, review)
	return 0, nil
}

func PostReview(c *gin.Context) (int, error) {

	type Request struct {
		Rating  int    `json:"rate"`
		Item    string `json:"item"`
		Message string `json:"message"`
	}

	req := Request{}
	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	req.Rating = int(math.Min(10, float64(req.Rating)))
	req.Rating = int(math.Max(0, float64(req.Rating)))

	db := database.GrabDB()
	username, _ := authorization.ExtractUser(c)

	var foo model.RawReview
	if err := db.Get(&foo, "SELECT * FROM reviews WHERE username=$1 AND product=$2", username, req.Item); err == nil {
		return http.StatusUnauthorized, errors.New("Você já avaliou este produto.")
	}

	if err := db.Get(&model.RawOrder{}, "SELECT * FROM orders WHERE holder=$1 AND product=$2", username, req.Item); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusUnauthorized, errors.New("Você não pode avaliar este produto.")
		}
		return http.StatusInternalServerError, err
	}

	final := model.RawReview{
		User:    username,
		Item:    req.Item,
		Message: req.Message,
		Rate:    req.Rating,
	}

	if err := repository.InsertReview(final); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Avaliação adicionada com sucesso"})
	return 0, nil
}

func PutReview(c *gin.Context) (int, error) {
	type Request struct {
		User    string `json:"user"`
		Item    string `json:"item"`
		Rate    int    `json:"rate"`
		Message string `json:"message"`
	}

	req := Request{}
	username, _ := authorization.ExtractUser(c)

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	var review model.Review
	if err := repository.DownloadReview(req.User, req.Item, &review); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("not found")
		}
		return http.StatusInternalServerError, err
	}

	if review.User.Username != username {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	final := model.RawReview{User: username, Message: req.Message, Rate: req.Rate, Item: review.Item.Item.Identifier}

	if err := repository.UpdateReview(final); err != nil {
		return 0, err
	}

	c.JSON(http.StatusOK, final)
	return 0, nil
}

func DeleteReview(c *gin.Context) (int, error) {
	item := c.Param("item")
	username, _ := authorization.ExtractUser(c)

	db := database.GrabDB()
	if err := db.Get(nil, "SELECT * FROM reviews WHERE username=$1 AND product=$2", username, item); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Nenhuma avaliação encontrada")
		}
	}

	if err := repository.DeleteReview(username, item); err != nil {
		return 0, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Avaliação excluída com sucesso"})
	return 0, nil
}
