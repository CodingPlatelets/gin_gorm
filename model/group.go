package model

import "time"

type Group struct {
	GroupId   uint      `gorm:"primaryKey;" json:"groupId" uri:"groupId"`
	GroupName string    `gorm:"varchar(255);unique;omitempty" json:"groupName" uri:"groupName"`
	ItemCOUNT int       `json:"item" uri:"item" gorm:"omitempty"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    uint      `json:"userId" uri:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" uri:"updatedAt" gorm:"autoUpdateTime"`
}
