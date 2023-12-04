package productstore

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/product/productmodel"

	"golang.org/x/net/context"
)

type ProductDBModel struct {
	Id            *string `json:"id" gorm:"column:id;"`
	ProductInfoId string  `json:"productInfoId" gorm:"column:productInfoId;"`
	Quantity      int     `json:"quantity" gorm:"column:qty;"`
	Edition       int     `json:"edition" gorm:"column:edition;"`
	Price         float64 `json:"price" gorm:"column:price;"`
	SalePrice     float64 `json:"salePrice" gorm:"column:salePrice;"`
	IsActive      bool    `json:"isActive" gorm:"column:isActive;"`
}

func (*ProductDBModel) TableName() string {
	return constants.TblProduct
}

func (s *sqlStore) CreateProduct(ctx context.Context, data *productmodel.ReqCreateProduct) error {
	db := s.db

	var tmpData ProductDBModel = ProductDBModel{
		ProductInfoId: "JFK",
		Quantity:      data.Quantity,
		Edition:       data.Edition,
		Price:         data.Price,
		SalePrice:     data.SalePrice,
	}
	if err := db.Create(tmpData).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return productmodel.ErrProductIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	//if err := db.Create(data).Error; err != nil {
	//	if gormErr := common.GetGormErr(err); gormErr != nil {
	//		switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
	//		case "PRIMARY":
	//			return productmodel.ErrProductIdDuplicate
	//		}
	//	}
	//	return common.ErrDB(err)
	//}

	return nil
}
