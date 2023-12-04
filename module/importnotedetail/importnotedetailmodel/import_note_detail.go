package importnotedetailmodel

import (
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/product/productmodel"
)

type ImportNoteDetail struct {
	ImportNoteId   string                     `json:"importNoteId" gorm:"column:importNoteId;" example:"import note id"`
	ProductId      string                     `json:"-" gorm:"column:productId;"`
	Product        productmodel.SimpleProduct `json:"product"`
	QuantityImport int                        `json:"qtyImport" gorm:"column:qtyImport;"`
	Price          float32                    `json:"price" gorm:"column:price;"`
}

func (*ImportNoteDetail) TableName() string {
	return constants.TblImportNoteDetail
}

var (
	ErrImportDetailProductIdInvalid = common.NewCustomError(
		errors.New("id of product is invalid"),
		"id of product is invalid",
		"ErrImportDetailProductIdInvalid",
	)
	ErrImportDetailPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"price of ingredient is negative number",
		"ErrImportDetailPriceIsNegativeNumber",
	)
	ErrImportDetailQuantityImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("quantity import is not positive number"),
		"quantity import is not positive number",
		"ErrImportDetailQuantityImportIsNotPositiveNumber",
	)
)
