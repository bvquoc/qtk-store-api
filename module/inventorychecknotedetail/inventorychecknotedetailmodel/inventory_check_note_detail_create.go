package inventorychecknotedetailmodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type InventoryCheckNoteDetailCreate struct {
	InventoryCheckNoteId string `json:"-" gorm:"column:inventoryCheckNoteId;"`
	ProductId            string `json:"productId" gorm:"column:productId;" example:"product id"`
	Initial              int    `json:"-" gorm:"column:initial;"`
	Difference           int    `json:"difference" gorm:"column:difference;" example:"100"`
	Final                int    `json:"-" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetailCreate) TableName() string {
	return constants.TblInventoryCheckNoteDetail
}

func (data *InventoryCheckNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.ProductId) {
		return ErrInventoryCheckDetailProductIdInvalid
	}
	if data.Difference == 0 {
		return ErrInventoryCheckDifferenceIsInvalid
	}
	return nil
}
