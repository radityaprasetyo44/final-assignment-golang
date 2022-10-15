package controllers

import (
	"final-assignment/models"
	"final-assignment/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	body := models.Photo{}
	c.ShouldBindJSON(&body)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]

	body.UserID = int(userId.(float64))

	res, err := services.CreatePhoto(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func GetPhoto(c *gin.Context) {
	res, err := services.GetPhoto()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdatePhoto(c *gin.Context) {
	body := models.Photo{}
	c.ShouldBindJSON(&body)

	conv, _ := strconv.Atoi(c.Param("photoId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	if body.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "title is required"})
		return
	}

	if body.PhotoUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "photo url is required"})
		return
	}

	res, err := services.UpdatePhoto(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeletePhoto(c *gin.Context) {
	var body models.Photo

	conv, _ := strconv.Atoi(c.Param("photoId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	res, err := services.DeletePhoto(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
