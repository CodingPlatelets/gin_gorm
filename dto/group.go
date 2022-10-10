package dto

import "github.com/WenkanHuang/gin_gorm/model"

type GroupDto struct {
	GroupId   uint   `gorm:"primaryKey;" json:"groupId" uri:"groupId"`
	GroupName string `gorm:"varchar(255);unique" json:"groupName" uri:"groupName"`
	ItemCOUNT int    `json:"item" uri:"item"`
}

func ToGroupDto(group model.Group) GroupDto {
	return GroupDto{
		GroupId:   group.GroupId,
		GroupName: group.GroupName,
		ItemCOUNT: group.ItemCOUNT,
	}
}
