package dao

import (
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/model"
)

func GetUserByName(name string) (*model.User, error) {
	user := new(model.User)
	if err := db.DB.Where("Name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserById(id uint) (*model.User, error) {
	user := new(model.User)
	if err := db.DB.Where("userId = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func DeleteUserById(id uint) error {
	user := model.User{}
	err := db.DB.Where("userId = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
