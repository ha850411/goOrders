package models

type ProductsType struct {
	Id         int    `json:"id"`
	Pid        int    `json:"pid"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	Sort       int    `json:"sort"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
