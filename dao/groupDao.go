package dao

import (
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/model"
)

func GetGroupsByUserId(id uint) ([]*model.Group, error) {
	var groups []*model.Group
	err := db.DB.Where("user_id=?", id).Find(&groups).Error
	if err != nil {
		return nil, err
	} else {
		return groups, nil
	}
}

func GetGroupByGroupName(name string) (*model.Group, error) {
	group := new(model.Group)
	if err := db.DB.Where("group_name=?", name).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func GetGroupByGroupId(id uint) (*model.Group, error) {
	group := new(model.Group)
	if err := db.DB.Where("group_id=?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func GetGroupByTodoId(id uint) (*model.Group, error) {
	todo := new(model.Todo)
	if err := db.DB.Where("todo_id=?", id).First(&todo).Error; err != nil {
		return nil, err
	} else {
		group := new(model.Group)
		if err := db.DB.Where("group_id=?", todo.GroupId).First(&group).Error; err != nil {
			return nil, err
		}
		return group, nil
	}
}

func UpdateGroup(group *model.Group) (*model.Group, error) {
	err := db.DB.Omit("created_at", "user_id").Save(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func DeleteGroupByName(name string) error {
	group := model.Group{}
	err := db.DB.Where("groupName = ?", name).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteGroupById(id uint) error {
	group := model.Group{}
	err := db.DB.Where("group_id = ?", id).Delete(&group).Error
	if err != nil {
		return err
	}
	return nil
}

func AddGroup(g *model.Group) (*model.Group, error) {
	err := db.DB.Create(&g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}
