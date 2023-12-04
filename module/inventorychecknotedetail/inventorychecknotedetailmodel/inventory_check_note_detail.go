package inventorychecknotedetailmodel

import (
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/product/productmodel"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string                     `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;" example:"inventory check note id"`
	ProductId            string                     `json:"-" gorm:"column:productId;"`
	Product              productmodel.SimpleProduct `json:"product"`
	Initial              int                        `json:"initial" gorm:"column:initial;" example:"100"`
	Difference           int                        `json:"difference" gorm:"column:difference;" example:"100"`
	Final                int                        `json:"final" gorm:"column:final;" example:"200"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return constants.TblInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailProductIdInvalid = common.NewCustomError(
		errors.New("id of product is invalid"),
		"id of product is is invalid",
		"ErrInventoryCheckDetailProductIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"difference is is invalid",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
