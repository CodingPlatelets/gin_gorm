package Model

import "time"

type Todo struct {
	TodoId      uint      `gorm:"primaryKey" json:"todoId" uri:"todoId"`
	TodoName    string    `gorm:"unique,omitempty" json:"todoName,omitempty" uri:"todoName"`
	TodoContent string    `gorm:"omitempty" json:"todoContent,omitempty" uri:"todoContent"`
	IsFinished  bool      `json:"isFinished,omitempty" uri:"isFinished" gorm:"force"`
	UserId      uint      `json:"userId,omitempty" uri:"userId" gorm:"omitempty"`
	GroupId     uint      `json:"groupId,omitempty" uri:"groupId" gorm:"omitempty"`
	User        User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Group       Group     `gorm:"foreignKey:GroupId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"createdAt,omitempty" uri:"createdAt" gorm:"autoCreateTime,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" uri:"updatedAt" gorm:"autoUpdateTime"`
}
