package supplierrepo

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
	"qtk-store-api/module/supplierdebt/supplierdebtmodel"
)

type ListSupplierDebtStore interface {
	ListSupplierDebt(
		ctx context.Context,
		supplierId string,
		filterSupplierDebt *filter.SupplierDebtFilter,
		paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error)
}

type seeSupplierDebtRepo struct {
	debtStore     ListSupplierDebtStore
	supplierStore FindSupplierStore
}

func NewSeeSupplierDebtRepo(
	debtStore ListSupplierDebtStore,
	supplierStore FindSupplierStore) *seeSupplierDebtRepo {
	return &seeSupplierDebtRepo{
		debtStore:     debtStore,
		supplierStore: supplierStore,
	}
}

func (biz *seeSupplierDebtRepo) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) (*suppliermodel.ResDebtSupplier, error) {
	supplier, errSupplier := biz.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId})
	if errSupplier != nil {
		return nil, errSupplier
	}

	resSeeDebtSupplier := suppliermodel.GetResSeeDebtSupplierFromSupplier(supplier)

	debts, errDebts := biz.debtStore.ListSupplierDebt(
		ctx,
		supplierId,
		filterSupplierDebt,
		paging,
	)
	if errDebts != nil {
		return nil, errDebts
	}

	resSeeDebtSupplier.DebtHistory = debts

	return resSeeDebtSupplier, nil
}
