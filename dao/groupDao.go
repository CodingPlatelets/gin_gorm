package dao

import (
	"gorm.io/gorm"

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

func UpdateGroupCount(ori, cur uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// find ori group
		oriTodo := &model.Todo{}
		if err := tx.Model(oriTodo).Where("todo_id=?", ori).First(oriTodo).Error; err != nil {
			return err
		}
		oriGroup := &model.Group{}
		if err := tx.Model(oriGroup).Where("group_id=?", oriTodo.GroupId).First(oriGroup).Error; err != nil {
			return err
		}

		// find current group
		curGroup := &model.Group{}
		if err := tx.Model(curGroup).Where("group_id=?", cur).First(curGroup).Error; err != nil {
			return err
		}
		if err := tx.Model(&oriGroup).Update("item_count", gorm.Expr("item_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.Model(&curGroup).Update("item_count", gorm.Expr("item_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
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
	}
	group := new(model.Group)
	if err := db.DB.Where("group_id=?", todo.GroupId).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil

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
