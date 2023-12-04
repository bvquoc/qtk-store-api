package inventorychecknotemodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

type ReqCreateInventoryCheckNote struct {
	Id                  *string                                                        `json:"id" gorm:"column:id;" example:""`
	QuantityDifferent   int                                                            `json:"-" gorm:"column:qtyDifferent;"`
	QuantityAfterAdjust int                                                            `json:"-" gorm:"column:qtyAfterAdjust;"`
	CreateBy            string                                                         `json:"-" gorm:"column:createBy;"`
	Details             []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ReqCreateInventoryCheckNote) TableName() string {
	return constants.TblInventoryCheckNote
}

func (data *ReqCreateInventoryCheckNote) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrInventoryCheckNoteIdInvalid
	}

	mapExits := make(map[string]int)
	for _, detail := range data.Details {
		if err := detail.Validate(); err != nil {
			return err
		}
		mapExits[detail.ProductId]++
		if mapExits[detail.ProductId] >= 2 {
			return ErrInventoryCheckNoteExistDuplicateProduct
		}
	}
	return nil
}
