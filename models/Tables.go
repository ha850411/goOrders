package models

// 紀錄桌號
type Tables struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"` // 0: 關閉, 1: 開啟
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
