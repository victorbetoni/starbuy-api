package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"net/http"
	"os"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostItem(c *gin.Context) error {

	var item model.PostedItem
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	item.Item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return nil
	}

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(c, item.Assets[0], uploader.UploadParams{
		PublicID: "assets/" + item.Item.Identifier})

	item.Assets[0] = resp.URL

	item.Item.Seller = user
	if err := repository.InsertItem(item); err != nil {
		return err
	}
	c.JSON(http.StatusOK, item)
	return nil
}

func QueryItems(c *gin.Context) error {
	query := c.Param("query")

	var items []model.ItemWithAssets
	if err := repository.QueryItemsByName(query, &items); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}
	c.JSON(http.StatusOK, items)

	return nil
}

func GetItem(c *gin.Context) error {
	queried := c.Param("id")
	key, ok := c.GetQuery("reviews")
	includeReviews := ok && key == "true"

	fmt.Println("1 ", queried)

	type ItemReview struct {
		User    model.User `json:"user,omitempty"`
		Message string     `json:"message"`
		Rate    int        `json:"rate"`
	}

	average := float64(-1)

	var incoming []model.Review
	var reviews []ItemReview
	if includeReviews {
		var err error
		average, err = repository.QueryProductReviews(queried, &incoming)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		for _, review := range incoming {
			reviews = append(reviews, ItemReview{User: review.User, Message: review.Message, Rate: review.Rate})
		}
	}

	fmt.Println("2")

	var item model.ItemWithAssets
	if err := repository.DownloadItem(queried, &item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	fmt.Println("3")

	type Response struct {
		Item    model.ItemWithAssets `json:"item,omitempty"`
		Reviews []ItemReview         `json:"reviews,omitempty"`
		Average float64              `json:"average"`
	}

	var response Response
	if includeReviews {
		response.Reviews = reviews
	}
	response.Item = item
	response.Average = average

	fmt.Println(response)

	c.JSON(http.StatusOK, response)
	return nil
}

func GetAllItems(c *gin.Context) error {
	var items []model.ItemWithAssets
	if err := repository.DownloadAllItems(&items); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}
	c.JSON(http.StatusOK, items)

	return nil
}

func GetCategory(c *gin.Context) error {
	queried, _ := strconv.Atoi(c.Param("id"))
	var items []model.ItemWithAssets

	if err := repository.DownloadItemByCategory(queried, &items); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, items)
	return nil
}
