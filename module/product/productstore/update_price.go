package productstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/product/productmodel"
)

func (s *sqlStore) UpdatePriceProduct(
	ctx context.Context,
	id string,
	data *productmodel.ProductUpdatePrice) error {
	db := s.db

	if err := db.Table(constants.TblProduct).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
