package Dao

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
)

func GetGroupsByUserId(id uint) ([]*Model.Group, error) {
	var groups []*Model.Group
	err := Db.DB.Where("userId=?", id).Find(&groups).Error
	if err != nil {
		return nil, err
	} else {
		return groups, nil
	}
}
func GetGroupByGroupName(name string) (*Model.Group, error) {
	group := new(Model.Group)
	if err := Db.DB.Where("groupName=?", name).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
func GetGroupByGroupId(id uint) (*Model.Group, error) {
	group := new(Model.Group)
	if err := Db.DB.Where("groupId=?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func UpdateGroup(group *Model.Group) (*Model.Group, error) {
	err := Db.DB.Save(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func DeleteGroupByName(name string) error {
	group := Model.Group{}
	err := Db.DB.Where("groupName = ?", name).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteGroupById(id uint) error {
	group := Model.Group{}
	err := Db.DB.Where("groupId = ?", id).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}
