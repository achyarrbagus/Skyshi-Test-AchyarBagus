package dto

type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResultCusSucces struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResultCusErr struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
