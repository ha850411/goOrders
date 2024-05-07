package models

type ProductType struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (ProductType) GetTableName() string {
	return "product_type"
}
