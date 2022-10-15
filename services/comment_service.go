package services

import (
	"final-assignment/configs"
	"final-assignment/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(payload *models.Comment) (models.Comment, gin.H) {
	db := configs.DBInit()

	err := db.Table("comments").Create(&payload).Error
	if err != nil {
		return models.Comment{}, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return *payload, nil
}

func GetComment() ([]models.GetComment, gin.H) {
	db := configs.DBInit()

	data := []models.GetComment{}
	res := db.Table("comments").Find(&data)
	if res.Error != nil || res.RowsAffected == 0 {
		return []models.GetComment{}, gin.H{
			"error": "data not found",
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

		var dataPhoto models.Photo
		errPhoto := db.Table("photos").Where(map[string]interface{}{
			"id": value.PhotoID,
		}).Find(&dataPhoto).Error
		if errPhoto == nil {
			data[i].Photo = dataPhoto
		}
	}

	return data, nil
}

func UpdateComment(payload *models.Comment) (models.Comment, gin.H) {
	db := configs.DBInit()
	var data models.Comment

	query := map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("comments").Where(query).Find(&data).Error
	if findUser != nil {
		return models.Comment{}, gin.H{
			"error": "comment not found",
		}
	}

	if payload.UserID != data.UserID {
		return models.Comment{}, gin.H{
			"error": "photo doesnt belong to this account",
		}
	}

	data.Message = payload.Message
	data.UpdatedAt = time.Now()

	err := db.Table("comments").Where(query).Updates(&data).Error
	if err != nil {
		return models.Comment{}, gin.H{
			"error":   "failed to update",
			"message": err.Error(),
		}
	}

	return data, nil
}

func DeleteComment(payload *models.Comment) (gin.H, gin.H) {
	db := configs.DBInit()

	var data models.Comment

	query := map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("comments").Where(query).Find(&data).Error
	if findUser != nil {
		return nil, gin.H{
			"error": "comment not found",
		}
	}

	if payload.UserID != data.UserID {
		return nil, gin.H{
			"error": "photo doesnt belong to this account",
		}
	}

	res := db.Table("comments").Where(query).Delete(&models.Comment{})
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, gin.H{
			"error": "can not delete comment",
		}
	}

	return gin.H{
		"message": "Your comment has been successfully deleted",
	}, nil
}
