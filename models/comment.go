package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	User    *User  `json:"user" gorm:"foreignKey:UserID"`
	Photo   *Photo `json:"photo" gorm:"foreignKey:PhotoID"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	return
}
