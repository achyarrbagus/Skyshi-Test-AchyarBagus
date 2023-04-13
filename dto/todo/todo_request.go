package tododto

type CreateTodoRequest struct {
	ActivityGroupId int    `json:"activity_group_id" gorm:"type:varchar(225)"`
	Title           string `json:"title" gorm:"type:varchar(225)"`
	IsActive        bool   `json:"is_active" gorm:"type:varchar(225)"`
	Priority        string `json:"priority" gorm:"type:varchar(225)"`
}

type UpdateTodoRequest struct {
	Title    string `json:"title" gorm:"type:varchar(225)"`
	Priority string `json:"priority" gorm:"type:varchar(225)"`
	IsActive bool   `json:"is_active" gorm:"type:varchar(225)"`
	Status   string `json:"status" gorm:"type:varchar(225)"`
}
