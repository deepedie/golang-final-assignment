package models

import (
	"assignment-4/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"full_name" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `json:"age" form:"age" valid:"required~Age is required,range(9|120)~Age must be greater than 8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}
