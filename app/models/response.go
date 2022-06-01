package models

type ResponseModel struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Timestamp  string      `json:"timestamp"`
}

type LoginResponseModel struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}
