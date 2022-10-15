package services

import (
	"final-assignment/configs"
	"final-assignment/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(payload *models.SocialMedia) (models.SocialMedia, gin.H) {
	db := configs.DBInit()

	err := db.Table("social_media").Create(&payload).Error
	if err != nil {
		return models.SocialMedia{}, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return *payload, nil
}

func GetSocialMedia() (gin.H, gin.H) {
	db := configs.DBInit()
	var data []models.GetSocialMedia

	res := db.Table("social_media").Find(&data)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, gin.H{
			"error": "social media not found",
		}
	}

	for i, value := range data {
		var dataUser models.GetUser
		errUser := db.Table("users").Where(map[string]interface{}{
			"id": value.UserID,
		}).Find(&dataUser).Error
		if errUser == nil {
			data[i].GetUser = models.GetUser{
				ID:       dataUser.ID,
				Username: dataUser.Username,
				Email:    dataUser.Email,
			}
		}
	}

	return gin.H{
		"social_medias": data,
	}, nil
}

func UpdateSocialMedia(payload *models.SocialMedia) (models.SocialMedia, gin.H) {
	db := configs.DBInit()
	var data models.SocialMedia

	query := map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("social_media").Where(query).Find(&data).Error
	if findUser != nil {
		return models.SocialMedia{}, gin.H{
			"error": "social media not found",
		}
	}

	fmt.Println(payload.UserID, data.UserID)
	if payload.UserID != data.UserID {
		return models.SocialMedia{}, gin.H{
			"error": "social media doesnt belong to this account",
		}
	}

	data.Name = payload.Name
	data.SocialMediaUrl = payload.SocialMediaUrl
	data.UpdatedAt = time.Now()

	err := db.Table("social_media").Where(query).Updates(&data).Error
	if err != nil {
		return models.SocialMedia{}, gin.H{
			"error":   "failed to update",
			"message": err.Error(),
		}
	}

	return data, nil
}

func DeleteSocialMedia(payload *models.SocialMedia) (gin.H, gin.H) {
	db := configs.DBInit()
	var (
		data  models.SocialMedia
		query map[string]interface{}
	)

	query = map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("social_media").Where(query).Find(&data).Error
	if findUser != nil {
		return nil, gin.H{
			"error": "photo not found",
		}
	}

	if payload.UserID != data.UserID {
		return nil, gin.H{
			"error": "social media doesnt belong to this account",
		}
	}

	query = map[string]interface{}{
		"id": payload.ID,
	}
	err := db.Table("social_media").Where(query).Delete(&data).Error
	if err != nil {
		return nil, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return gin.H{
		"message": "Your social media has been successfully deleted",
	}, nil
}
