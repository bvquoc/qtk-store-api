package importnotemodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
)

type ReqCreateImportNote struct {
	Id                *string                                        `json:"id" gorm:"column:id;" example:""`
	TotalPrice        float32                                        `json:"-" gorm:"column:totalPrice;"`
	SupplierId        string                                         `json:"supplierId" gorm:"column:supplierId" example:"supplier id"`
	CreateBy          string                                         `json:"-" gorm:"column:createBy;"`
	ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ReqCreateImportNote) TableName() string {
	return constants.TblImportNote
}

func (data *ReqCreateImportNote) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrImportNoteIdInvalid
	}
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrImportNoteSupplierIdInvalid
	}
	if data.ImportNoteDetails == nil || len(data.ImportNoteDetails) == 0 {
		return ErrImportNoteDetailsEmpty
	}

	mapIngredientUpdatePriceTimes := make(map[string]int)
	for _, importNoteDetail := range data.ImportNoteDetails {
		if err := importNoteDetail.Validate(); err != nil {
			return err
		}
		mapIngredientUpdatePriceTimes[importNoteDetail.ProductId]++
		if mapIngredientUpdatePriceTimes[importNoteDetail.ProductId] > 1 {
			return ErrImportNoteHasSameProduct
		}
	}
	return nil
}

func (data *ReqCreateImportNote) Round() {
	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].Round()
	}
}
