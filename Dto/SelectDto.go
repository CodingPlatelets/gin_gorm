package Dto

import "time"

type Condition struct {
	CreatedAt  time.Time `time_format:"2006-01-02 15:04:05" json:"created_at" uri:"created_at" form:"created_at"`
	Keyword    string    `json:"keyword" uri:"keyword" form:"keyword"`
	GroupId    int       `json:"group_id" uri:"group_id" form:"group_id"`
	IsFinished bool      `json:"is_finished" uri:"is_finished" form:"is_finished"`
}
