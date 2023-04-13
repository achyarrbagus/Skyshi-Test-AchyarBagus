package tododto

import "time"

type TodoResponse struct {
	ID              int       `json:"id"`
	ActivityGroupId int       `json:"activity_group_id" gorm:"type:varchar(225)"`
	Title           string    `json:"title" gorm:"type:varchar(225)"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"-"`
	UpdateAt        time.Time `json:"-"`
}
