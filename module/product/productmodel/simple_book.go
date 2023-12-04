package productmodel

import "qtk-store-api/constants"

type SimpleProduct struct {
	ID   string `json:"id" gorm:"column:id;primaryKey" example:"product id"`
	Name string `json:"name" gorm:"column:name" example:"Những câu chuyện hay"`
}

func (*SimpleProduct) TableName() string {
	return constants.TblProduct
}
