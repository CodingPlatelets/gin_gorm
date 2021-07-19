package Model

import "time"

type Group struct {
	GroupId   uint      `gorm:"primaryKey;" json:"groupId" uri:"groupId"`
	GroupName string    `gorm:"varchar(255);unique" json:"groupName" uri:"groupName"`
	ItemCOUNT int       `json:"item" uri:"item"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    uint      `json:"userId" uri:"userId"`
	CreatedAt time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" uri:"updatedAt" gorm:"autoUpdateTime"`
}
