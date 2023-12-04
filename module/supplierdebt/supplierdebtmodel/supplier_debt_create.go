package supplierdebtmodel

import (
	"qtk-store-api/common/enum"
	"qtk-store-api/constants"
)

type SupplierDebtCreate struct {
	Id           string         `json:"-" gorm:"column:id;"`
	SupplierId   string         `json:"supplierId" gorm:"column:supplierId;"`
	Quantity     float32        `json:"qty" gorm:"column:qty;"`
	QuantityLeft float32        `json:"-" gorm:"column:qtyLeft;"`
	DebtType     *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy     string         `json:"-" gorm:"column:createBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return constants.TblSupplierDebt
}
