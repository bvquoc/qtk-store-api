package importnotedetailmodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type ImportNoteDetailCreate struct {
	ImportNoteId   string  `json:"-" gorm:"column:importNoteId;"`
	ProductId      string  `json:"productId" gorm:"column:productId;" example:"product id"`
	QuantityImport float32 `json:"qtyImport" gorm:"column:qtyImport;" example:"100"`
	Price          float32 `json:"price" gorm:"column:price;" example:"60000"`
	IsReplacePrice bool    `json:"isReplacePrice" gorm:"-" example:"true"`
}

func (*ImportNoteDetailCreate) TableName() string {
	return constants.TblImportNoteDetail
}

func (data *ImportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.ProductId) {
		return ErrImportDetailProductIdInvalid
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrImportDetailPriceIsNegativeNumber
	}
	if common.ValidateNotPositiveNumber(data.QuantityImport) {
		return ErrImportDetailQuantityImportIsNotPositiveNumber
	}
	return nil
}

func (data *ImportNoteDetailCreate) Round() {
	common.CustomRound(&data.Price)
	common.CustomRound(&data.QuantityImport)
}
