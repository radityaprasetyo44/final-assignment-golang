package controllers

import (
	"final-assignment/models"
	"final-assignment/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	body := models.SocialMedia{}
	c.ShouldBindJSON(&body)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]

	body.UserID = int(userId.(float64))

	res, err := services.CreateSocialMedia(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func GetSocialMedia(c *gin.Context) {
	res, err := services.GetSocialMedia()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateSocialMedia(c *gin.Context) {
	body := models.SocialMedia{}
	c.ShouldBindJSON(&body)

	conv, _ := strconv.Atoi(c.Param("socialMediaId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	if body.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "name is required"})
		return
	}

	if body.SocialMediaUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "social media url is required"})
		return
	}

	res, err := services.UpdateSocialMedia(&body)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteSocialMedia(c *gin.Context) {
	var body models.SocialMedia

	conv, _ := strconv.Atoi(c.Param("socialMediaId"))
	body.ID = uint(conv)

	userData := c.MustGet("user_data").(jwt.MapClaims)
	userId := userData["id"]
	body.UserID = int(userId.(float64))

	res, err := services.DeleteSocialMedia(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
