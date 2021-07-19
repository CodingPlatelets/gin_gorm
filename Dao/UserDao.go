package Dao

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
)

func GetUserByName(name string) (*Model.User, error) {
	user := new(Model.User)
	if err := Db.DB.Where("Name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserById(id uint) (*Model.User, error) {
	user := new(Model.User)
	if err := Db.DB.Where("userId = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func DeleteUserById(id uint) error {
	user := Model.User{}
	err := Db.DB.Where("userId = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
