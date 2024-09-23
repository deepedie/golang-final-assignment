package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string    `json:"title" form:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoURL  string    `json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID    uint      `json:"user_id"`
	User      *User     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
