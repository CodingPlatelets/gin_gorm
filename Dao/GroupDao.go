package Dao

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
)

func GetGroupsByUserId(id uint) ([]*Model.Group, error) {
	var groups []*Model.Group
	err := Db.DB.Where("user_id=?", id).Find(&groups).Error
	if err != nil {
		return nil, err
	} else {
		return groups, nil
	}
}
func GetGroupByGroupName(name string) (*Model.Group, error) {
	group := new(Model.Group)
	if err := Db.DB.Where("group_name=?", name).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
func GetGroupByGroupId(id uint) (*Model.Group, error) {
	group := new(Model.Group)
	if err := Db.DB.Where("group_id=?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
func GetGroupByTodoId(id uint) (*Model.Group, error) {
	todo := new(Model.Todo)
	if err := Db.DB.Where("todo_id=?", id).First(&todo).Error; err != nil {
		return nil, err
	} else {
		group := new(Model.Group)
		if err := Db.DB.Where("group_id=?", todo.GroupId).First(&group).Error; err != nil {
			return nil, err
		}
		return group, nil
	}
}

func UpdateGroup(group *Model.Group) (*Model.Group, error) {
	err := Db.DB.Omit("created_at", "user_id").Save(&group).Error
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
	err := Db.DB.Where("group_id = ?", id).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}

func AddGroup(g *Model.Group) (*Model.Group, error) {
	err := Db.DB.Create(&g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}
