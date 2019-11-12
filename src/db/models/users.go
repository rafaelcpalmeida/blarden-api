package models

import (
	"blarden-api/src/db"
	"errors"
	"github.com/gofrs/uuid"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	Id             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	PhoneNumber    string    `json:"phone_number" gorm:"type:varchar(9);"`
	HasPermission  bool      `json:"has_permission" gorm:"type:bool"`
	UserType       uint8     `json:"user_type" gorm:"default:3"`
	UserUniqueKey  string    `json:"user_unique_key" gorm:"type:varchar(128);"`
	AccountBlocked bool      `json:"account_blocked" gorm:"type:bool"`
	Timestamps
}

func AllUsers(queryParameters map[string]interface{}) ([]User, error) {
	var users []User

	err := db.DatabaseHandler().Where(queryParameters).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SpecificUser(id uuid.UUID) (User, error) {
	user := specificUser(id)

	if id != user.Id {
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

func UpdateUser(id uuid.UUID, user User) (User, error) {
	retrievedUser := specificUser(id)

	if retrievedUser.Id != id {
		return User{}, errors.New("unable to find user")
	}

	err := db.DatabaseHandler().Model(&user).Updates(map[string]interface{}{"has_permission": user.HasPermission,
		"account_blocked": user.AccountBlocked}).Error

	if err != nil {
		return User{}, err
	}

	return specificUser(user.Id), nil
}

func DeleteUser(id uuid.UUID) (bool, error) {
	user := specificUser(id)

	if id != user.Id {
		return false, errors.New("unable to find user")
	}

	if err := db.DatabaseHandler().Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func specificUser(id uuid.UUID) User {
	var user User

	db.DatabaseHandler().Where(&User{Id: id}).First(&user)

	return user
}
