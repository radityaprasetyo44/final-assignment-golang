package services

import (
	"final-assignment/configs"
	"final-assignment/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(payload *models.Photo) (models.Photo, gin.H) {
	db := configs.DBInit()

	err := db.Table("photos").Create(&payload).Error
	if err != nil {
		return models.Photo{}, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return *payload, nil
}

func GetPhoto() ([]models.GetPhoto, gin.H) {
	db := configs.DBInit()
	var data []models.GetPhoto

	res := db.Table("photos").Find(&data)
	if res.Error != nil || res.RowsAffected == 0 {
		return []models.GetPhoto{}, gin.H{
			"error": "data not found",
		}
	}

	for i, value := range data {
		var dataUser models.GetUser
		err := db.Debug().Table("users").Where(map[string]interface{}{
			"id": value.UserID,
		}).Find(&dataUser).Error
		if err == nil {
			data[i].User = models.GetUser{
				Username: dataUser.Username,
				Email:    dataUser.Email,
			}
		}
	}

	return data, nil
}

func UpdatePhoto(payload *models.Photo) (models.Photo, gin.H) {
	db := configs.DBInit()
	var data models.Photo

	query := map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("photos").Where(query).Find(&data).Error
	if findUser != nil {
		return models.Photo{}, gin.H{
			"error": "photo not found",
		}
	}

	if payload.UserID != data.UserID {
		return models.Photo{}, gin.H{
			"error": "photo doesnt belong to this account",
		}
	}

	data.Title = payload.Title
	data.Caption = payload.Caption
	data.PhotoUrl = payload.PhotoUrl
	data.UpdatedAt = time.Now()

	err := db.Table("photos").Where(query).Updates(&data).Error
	if err != nil {
		return models.Photo{}, gin.H{
			"error":   "failed to update",
			"message": err.Error(),
		}
	}

	return data, nil
}

func DeletePhoto(payload *models.Photo) (gin.H, gin.H) {
	db := configs.DBInit()
	var (
		data  models.Photo
		query map[string]interface{}
	)

	query = map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("photos").Where(query).Find(&data).Error
	if findUser != nil {
		return nil, gin.H{
			"error": "photo not found",
		}
	}

	if payload.UserID != data.UserID {
		return nil, gin.H{
			"error": "photo doesnt belong to this account",
		}
	}

	err := db.Table("photos").Where(query).Delete(&data).Error
	if err != nil {
		return nil, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return gin.H{
		"message": "Your photo has been successfully deleted",
	}, nil
}
