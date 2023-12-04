package productstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/product/productmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateQuantityProduct(
	ctx context.Context,
	id string,
	data *productmodel.ProductUpdateQuantity) error {
	db := s.db

	if err := db.Table(constants.TblProduct).
		Where("id = ?", id).
		Update("qty", gorm.Expr("qty + ?", data.QuantityUpdate)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
