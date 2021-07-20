package Model

import "time"

type User struct {
	UserId    uint      `gorm:"primaryKey;" json:"id" uri:"id"`
	Name      string    `gorm:"varchar(20);not null;unique" json:"name" uri:"name" binding:"required"`
	Password  string    `gorm:"size:255;not null" json:"password" uri:"password" binding:"required"`
	CreatedAt time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime,omitempty" json:"created_at" binding:"-"`
}
