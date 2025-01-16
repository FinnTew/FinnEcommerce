package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email           string `gorm:"type:varchar(255);unique_index"`
	Password        string `gorm:"type:varchar(255);not null"`
	IsEmailVerified bool   `gorm:"type:tinyint(1);default:false"`
	Role            string `gorm:"type:varchar(50);default:'user'"`
	Permissions     string `gorm:"type:json"`
}

func (User) TableName() string {
	return "users"
}

func GetUserByID(db *gorm.DB, ctx context.Context, id uint32) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id=?", id).First(&user).Error
	return
}
