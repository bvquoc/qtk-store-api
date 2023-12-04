package productmodel

import "qtk-store-api/constants"

type ProductUpdatePrice struct {
	Price *float32 `json:"price" gorm:"column:price;"`
}

func (*ProductUpdatePrice) TableName() string {
	return constants.TblProduct
}
