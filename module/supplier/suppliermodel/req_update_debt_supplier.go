package suppliermodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type ReqUpdateDebtSupplier struct {
	QuantityUpdate *float32 `json:"qtyUpdate" gorm:"-" example:"10000"`
	CreateBy       string   `json:"-" gorm:"-"`
}

func (*ReqUpdateDebtSupplier) TableName() string {
	return constants.TblSupplier
}

func (data *ReqUpdateDebtSupplier) Validate() *common.AppError {
	if data.QuantityUpdate == nil {
		return ErrSupplierDebtPayNotExist
	}
	if *data.QuantityUpdate == 0 {
		return ErrSupplierDebtPayIsInvalid
	}
	return nil
}

func (data *ReqUpdateDebtSupplier) Round() {
	common.CustomRound(data.QuantityUpdate)
}
