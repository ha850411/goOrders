package models

type Users struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Role         int    `json:"role"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	LinebotToken string `json:"linebot_token"`
	UpdateTime   string `json:"update_time"`
}
