package models

import "time"

type ActivityGroup struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email" gorm:"type: varchar(255);unique;not null"`
	Todo      []Todo    `json:"todos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ActivityGroupTodoResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Email string `json:"email"`
}

func (ActivityGroupTodoResponse) TableName() string {
	return "activity_groups"
}
