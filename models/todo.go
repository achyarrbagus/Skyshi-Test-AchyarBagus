package models

import "time"

type Todo struct {
	ID              int                       `json:"id"`
	Title           string                    `json:"title"`
	ActivityGroupId int                       `json:"-"`
	ActivityGroup   ActivityGroupTodoResponse `json:"activity_groups"`
	IsActive        bool                      `json:"is_active"`
	Priority        string                    `json:"priority"`
	Status          string                    `json:"status"`
	CreatedAt       time.Time                 `json:"-"`
	UpdateAt        time.Time                 `json:"-"`
}
