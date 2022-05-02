package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostNotification(c *gin.Context) error {

	var item model.RawNotification
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)
	item.SentIn = time.Now().Format("2006-01-02 15:04:05")

	repository.InsertNotification(item)
	c.JSON(http.StatusOK, item)
	return nil

}

func GetNotifications(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	var notifications []model.RawNotification

	if err := repository.DownloadNotifications(user, &notifications); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
		}
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
	}

	c.JSON(http.StatusOK, notifications)
	return nil
}

func GetNotification(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)
	queried := c.Param("id")

	var notification model.Notification

	if user != notification.User.Username {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		return nil
	}

	if err := repository.DownloadNotification(queried, &notification); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return nil
	}

	c.JSON(http.StatusOK, notification)
	return nil
}
