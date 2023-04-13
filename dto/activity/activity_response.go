package activitydto

import "time"

type ActivityResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
