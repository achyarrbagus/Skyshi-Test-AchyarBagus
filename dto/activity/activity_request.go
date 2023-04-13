package activitydto

type CreateActivityRequest struct {
	Title string `json:"title" form:"title"`
	Email string `json:"email" form:"email" `
}

type UpdateActivityRequest struct {
	Title string `json:"title" form:"title" `
}
