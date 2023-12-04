package supplierdebtstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/supplierdebt/supplierdebtmodel"
)

func (s *sqlStore) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
