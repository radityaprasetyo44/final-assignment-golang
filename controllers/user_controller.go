package controllers

import (
	"final-assignment/models"
	"final-assignment/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	body := models.User{}
	c.ShouldBindJSON(&body)

	if body.Age <= 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "age must be greater than 8",
		})
		return
	}

	res, err := services.Register(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func Login(c *gin.Context) {
	body := models.User{}
	c.ShouldBindJSON(&body)

	res, err := services.Login(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {
	body := models.UserUpdate{}
	c.ShouldBindJSON(&body)

	conv, _ := strconv.Atoi(c.Param("userId"))
	body.ID = conv

	if body.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "email is required"})
		return
	}

	if body.Username == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "username is required"})
		return
	}

	res, err := services.UpdateUser(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	userData := c.MustGet("user_data").(jwt.MapClaims)
	id := userData["id"]

	res, err := services.DeleteUser(int(id.(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
