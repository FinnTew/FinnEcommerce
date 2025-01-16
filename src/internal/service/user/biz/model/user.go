package model

import (
	"context"
	"fmt"
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

func CreateUser(db *gorm.DB, ctx context.Context, user *User) (err error) {
	_, err = GetUserByEmail(db, ctx, user.Email)
	if err == nil {
		return fmt.Errorf("user %s already exists", user.Email)
	}

	err = db.WithContext(ctx).Model(&User{}).Create(&user).Error
	return
}

func DeleteUserByID(db *gorm.DB, ctx context.Context, userID uint32) (err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id=?", userID).Delete(&User{}).Error
	return
}

func UpdateUser(db *gorm.DB, ctx context.Context, user *User) (err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id=?", user.ID).Updates(&user).Error
	return
}

func GetUserByID(db *gorm.DB, ctx context.Context, id uint32) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("id=?", id).First(&user).Error
	return
}

func GetUserByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where("email=?", email).First(&user).Error
	return
}
