package productmodel

import (
	"qtk-store-api/common"
)

type ReqCreateProduct struct {
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"desc" gorm:"column:desc"`
	Edition     int      `json:"edition" gorm:"column:edition"`
	Quantity    int      `json:"quantity" gorm:"column:qty"`
	Price       float64  `json:"price" gorm:"column:price"`
	SalePrice   float64  `json:"salePrice" gorm:"column:salePrice"`
	PublisherID string   `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds"`
}

//func (*ReqCreateProduct) TableName() string {
//	return constants.TableProduct
//}

func (data *ReqCreateProduct) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrProductNameEmpty
	}

	if data.Price <= 0 {
		return ErrProductPriceIsLessThanZero
	}

	if data.SalePrice <= 0 {
		return ErrProductSalePriceIsLessThanZero
	}

	if data.Quantity < 0 {
		return ErrProductQuantityIsNegativeNumber
	}

	if data.Edition < 0 {
		return ErrProductEditionIsNegativeNumber
	}

	if common.ValidateEmptyString(data.PublisherID) {
		return ErrProductPublisherIdEmpty
	}

	if len(data.AuthorIDs) == 0 {
		return ErrProductAuthorIdsEmpty
	}

	if len(data.CategoryIDs) == 0 {
		return ErrProductCategoryIdsEmpty
	}

	return nil
}
