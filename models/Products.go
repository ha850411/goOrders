package models

type Products struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Amount        int    `json:"amount"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"DiscountPrice"`
	Status        int    `json:"status"`
	Content       string `json:"content"`
	Sort          int    `json:"sort"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
}