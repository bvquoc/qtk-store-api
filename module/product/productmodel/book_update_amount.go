package productmodel

import "qtk-store-api/constants"

type ProductUpdateQuantity struct {
	QuantityUpdate int `json:"qtyUpdate" gorm:"-"`
}

func (*ProductUpdateQuantity) TableName() string {
	return constants.TblProduct
}
