package models

import (
	"blarden-api/src/db"
	"errors"
	"github.com/gofrs/uuid"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	PhoneNumber   string    `json:"phone_number" gorm:"type:varchar(9);"`
	HasPermission bool      `json:"has_permission" gorm:"type:bool"`
	Timestamps
}

func AllUsers() ([]User, error) {
	var users []User

	err := db.DatabaseHandler().Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SpecificUser(id uuid.UUID) (User, error) {
	user := specificUser(id)

	if id != user.ID {
		return User{}, errors.New("unable to find user")
	}

	return user, nil
}

func NewUser(user User) (User, error) {
	var userCount int64
	db.DatabaseHandler().Model(User{}).Where("phone_number = ?", user.PhoneNumber).Count(&userCount)

	if userCount >= 1 {
		return User{}, errors.New("user is already registered")
	}

	err := db.DatabaseHandler().Save(&user).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) (User, error) {
	err := db.DatabaseHandler().Save(&user).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUser(id uuid.UUID) (bool, error) {
	user := specificUser(id)

	if id != user.ID {
		return false, errors.New("unable to find user")
	}

	if err := db.DatabaseHandler().Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func specificUser(id uuid.UUID) User {
	var user User

	db.DatabaseHandler().Where(&User{ID: id}).First(&user)

	return user
}
