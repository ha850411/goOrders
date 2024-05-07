package models

type Users struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Role         int    `json:"role"`
	LinebotToken string `json:"linebot_token"`
	UpdateTime   string `json:"update_time"`
}
