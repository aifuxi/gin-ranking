package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func CreateUser(name string, password string) (User, error) {
	user := User{
		Name:     name,
		Password: password,
	}

	err := DB.Create(&user).Error

	return user, err
}

func GetUsers() ([]User, error) {
	var users []User

	err := DB.Find(&users).Error

	return users, err
}

func GetUser(id int) (User, error) {
	var user User
	err := DB.First(&user).Error

	return user, err
}

func UpdateUser(id int, name string, password string) (User, error) {
	user, err := GetUser(id)
	if err != nil {
		return User{}, err
	}
	if user.ID == 0 {
		return User{}, errors.New("user not found")
	}

	err = DB.Model(&user).Updates(User{
		Name:     name,
		Password: password,
	}).Error
	if err != nil {
		return User{}, err
	}

	return user, err
}

func DeleteUser(id int) error {
	user, err := GetUser(id)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("user not found")
	}

	err = DB.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
