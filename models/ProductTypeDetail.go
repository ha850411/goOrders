package models

type ProductTypeDetail struct {
	Tid int `json:"tid"`
	Pid int `json:"pid"`
}

func (ProductTypeDetail) GetTableName() string {
	return "product_type_detail"
}
