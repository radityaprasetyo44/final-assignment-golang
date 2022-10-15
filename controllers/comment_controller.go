package controllers

import (
	"final-assignment/models"
	"final-assignment/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	body := models.Comment{}
	c.ShouldBindJSON(&body)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]

	body.UserID = int(userId.(float64))

	res, err := services.CreateComment(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func GetComment(c *gin.Context) {
	res, err := services.GetComment()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateComment(c *gin.Context) {
	var body models.Comment

	c.ShouldBindJSON(&body)

	conv, _ := strconv.Atoi(c.Param("commentId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	if body.Message == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "message is required"})
		return
	}

	res, err := services.UpdateComment(&body)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteComment(c *gin.Context) {
	var body models.Comment

	conv, _ := strconv.Atoi(c.Param("commentId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	res, err := services.DeleteComment(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
