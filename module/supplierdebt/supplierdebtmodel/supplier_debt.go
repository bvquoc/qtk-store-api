package supplierdebtmodel

import (
	"qtk-store-api/common/enum"
	"qtk-store-api/constants"
	"time"
)

type SupplierDebt struct {
	Id           string         `json:"id" gorm:"column:id;" example:"debt001"`
	SupplierId   string         `json:"supplierId" gorm:"column:supplierId;" example:"123"`
	Quantity     float32        `json:"qty" gorm:"column:qty;" example:"-70000"`
	QuantityLeft float32        `json:"qtyLeft" gorm:"column:qtyLeft;" example:"-100000"`
	DebtType     *enum.DebtType `json:"type" gorm:"column:type;" example:"Debt"`
	CreateBy     string         `json:"createBy" gorm:"column:createBy;" example:"user_id"`
	CreateAt     *time.Time     `json:"createAt" gorm:"column:createAt;" example:"1709500431"`
}

func (*SupplierDebt) TableName() string {
	return constants.TblSupplierDebt
}
