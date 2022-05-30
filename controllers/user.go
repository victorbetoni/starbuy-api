package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"starbuy/model"
	"starbuy/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type IncomingUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Birthdate      string `json:"birthdate"`
	Seller         bool   `json:"seller"`
	ProfilePicture string `json:"profile_picture"`
	City           string `json:"city"`
	Password       string `json:"password"`
}

func Register(c *gin.Context) error {

	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	user := model.User{
		Username:       incoming.Username,
		Email:          incoming.Email,
		Name:           incoming.Name,
		Birthdate:      incoming.Birthdate,
		ProfilePicture: incoming.ProfilePicture,
		Seller:         incoming.Seller,
		City:           incoming.City,
		Registration:   time.Now().Format("02-01-2006"),
	}

	if err := user.Prepare(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	if err := repository.InsertUser(user, incoming.Password); err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)

	return nil
}

func GetUser(c *gin.Context) error {
	queried := c.Param("user")

	key, ok := c.GetQuery("includeItems")
	includeItems := ok && key == "true"

	var user model.User

	var items []model.ItemWithAssets

	if includeItems {
		var local []model.ItemWithAssets
		if err := repository.DownloadUserProducts(queried, &local); err != nil && err != sql.ErrNoRows {
			return err
		}

		//Removing seller (duplicated data)
		for _, item := range local {
			final := model.Item{
				Identifier:  item.Item.Identifier,
				Title:       item.Item.Title,
				Category:    item.Item.Category,
				Stock:       item.Item.Stock,
				Description: item.Item.Description,
				Price:       item.Item.Price,
			}
			items = append(items, model.ItemWithAssets{Item: final, Assets: item.Assets})
		}
	}

	if err := repository.DownloadUser(queried, &user); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	if includeItems {

		type UserWithItem struct {
			User  model.User             `json:"user"`
			Items []model.ItemWithAssets `json:"items"`
		}
		c.JSON(http.StatusOK, UserWithItem{user, items})
		return nil
	}

	fmt.Println(user)
	c.JSON(http.StatusOK, user)
	return nil
}
