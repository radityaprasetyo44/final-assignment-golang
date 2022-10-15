package models

import (
	"final-assignment/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Type of Migrate
type (
	User struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Username  string    `json:"username" valid:"required~username is required" gorm:"not null;unique"`
		Email     string    `json:"email" valid:"required~email is required,email~invalid email format" gorm:"not null;unique"`
		Password  string    `json:"password" valid:"required~password is required,minstringlength(6)~password has to have a minimum length of 6 characters" gorm:"not null"`
		Age       int       `json:"age" valid:"required~age is required" gorm:"not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Photo struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Title     string    `json:"title" valid:"required~title is required" gorm:"not null"`
		Caption   string    `json:"caption"`
		PhotoUrl  string    `json:"photo_url" valid:"required~photo_url is required" gorm:"not null"`
		UserID    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Comment struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		UserID    int       `json:"user_id"`
		PhotoID   int       `json:"photo_id"`
		Message   string    `json:"message" valid:"required~message is required" gorm:"not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	SocialMedia struct {
		ID             uint      `json:"id" gorm:"primary_key"`
		Name           string    `json:"name" valid:"required~name is required" gorm:"not null"`
		SocialMediaUrl string    `json:"social_media_url" valid:"required~social_media_url is required" gorm:"not null"`
		UserID         int       `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

// Type of Request
type (
	UserUpdate struct {
		ID       int    `json:"id"`
		Username string `json:"username" valid:"required~username is required" gorm:"not null;unique"`
		Email    string `json:"email" valid:"required~email is required,email~invalid email format" gorm:"not null;unique"`
	}
)

// Type of View
type (
	GetPhoto struct {
		Photo
		User GetUser `json:"user"`
	}

	GetUser struct {
		ID       int    `json:"id,omitempty"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	GetComment struct {
		Comment
		GetUser GetUser `json:"user"`
		Photo   Photo   `json:"photo"`
	}

	GetSocialMedia struct {
		SocialMedia
		GetUser GetUser `json:"user"`
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
