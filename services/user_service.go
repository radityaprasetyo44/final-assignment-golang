package services

import (
	"final-assignment/configs"
	"final-assignment/helpers"
	"final-assignment/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(payload *models.User) (gin.H, gin.H) {
	db := configs.DBInit()

	err := db.Table("users").Create(&payload).Error
	if err != nil {
		return nil, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}
	}

	return gin.H{
		"age":      payload.Age,
		"email":    payload.Email,
		"id":       payload.ID,
		"username": payload.Username,
	}, nil
}

func Login(payload *models.User) (gin.H, gin.H) {
	db := configs.DBInit()
	password := ""

	password = payload.Password

	query := map[string]interface{}{
		"email": payload.Email,
	}

	var data models.User
	err := db.Table("users").Where(query).Find(&data).Error
	if err != nil {
		return nil, gin.H{
			"error":   "account not found",
			"message": err.Error(),
		}
	}

	comparePass := helpers.ComparePass([]byte(data.Password), []byte(password))
	if !comparePass {
		return nil, gin.H{
			"error":   "Unauthorized",
			"message": "invalid password",
		}
	}

	token := helpers.GenerateToken(data.ID, data.Email)

	return gin.H{
		"token": token,
	}, nil
}

func UpdateUser(payload *models.UserUpdate) (gin.H, gin.H) {
	db := configs.DBInit()
	var (
		err  error
		data models.User
	)

	query := map[string]interface{}{
		"id": payload.ID,
	}
	findUser := db.Table("users").Where(query).Find(&data).Error
	if findUser != nil {
		return nil, gin.H{
			"error": "account not found",
		}
	}

	data.Email = payload.Email
	data.Username = payload.Username
	data.UpdatedAt = time.Now()

	err = db.Table("users").Where("id = ?", payload.ID).Updates(&data).Error
	if err != nil {
		return nil, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		}

	}

	return gin.H{
		"age":        data.Age,
		"email":      data.Email,
		"id":         data.ID,
		"username":   data.Username,
		"updated_at": data.UpdatedAt,
	}, nil
}

func DeleteUser(id int) (gin.H, gin.H) {
	db := configs.DBInit()

	res := db.Table("users").Where("id = ?", id).Delete(&models.User{})
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, gin.H{
			"error": "can not delete user",
		}
	}

	return gin.H{
		"message": "Your account has been successfully deleted",
	}, nil
}
