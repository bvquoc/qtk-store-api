package importnoterepo

import (
	"context"
	"qtk-store-api/common/enum"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
	"qtk-store-api/module/product/productmodel"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplierdebt/supplierdebtmodel"
)

type ChangeStatusImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		id string,
		data *importnotemodel.ReqUpdateImportNote,
	) error
}

type GetImportNoteDetailStore interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateQuantityProductStore interface {
	UpdateQuantityProduct(
		ctx context.Context,
		id string,
		data *productmodel.ProductUpdateQuantity,
	) error
}

type UpdateDebtOfSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.ReqUpdateDebtSupplier,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore       ChangeStatusImportNoteStore
	importNoteDetailStore GetImportNoteDetailStore
	productStore          UpdateQuantityProductStore
	supplierStore         UpdateDebtOfSupplierStore
	supplierDebtStore     CreateSupplierDebtStore
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStore,
	importNoteDetailStore GetImportNoteDetailStore,
	productStore UpdateQuantityProductStore,
	supplierStore UpdateDebtOfSupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		productStore:          productStore,
		supplierStore:         supplierStore,
		supplierDebtStore:     supplierDebtStore,
	}
}

func (repo *changeStatusImportNoteRepo) FindImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, err := repo.importNoteStore.FindImportNote(
		ctx, map[string]interface{}{"id": importNoteId},
	)
	if err != nil {
		return nil, err
	}
	return importNote, nil
}

func (repo *changeStatusImportNoteRepo) UpdateImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ReqUpdateImportNote) error {
	if err := repo.importNoteStore.UpdateImportNote(
		ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ReqUpdateImportNote) error {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx,
		map[string]interface{}{"id": importNote.SupplierId})
	if err != nil {
		return err
	}

	qtyBorrow := -importNote.TotalPrice
	qtyLeft := supplier.Debt + qtyBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:           supplierDebtId,
		SupplierId:   importNote.SupplierId,
		Quantity:     qtyBorrow,
		QuantityLeft: qtyLeft,
		DebtType:     &debtType,
		CreateBy:     importNote.CloseBy,
	}

	if err := repo.supplierDebtStore.CreateSupplierDebt(
		ctx, &supplierDebtCreate,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) UpdateDebtSupplier(
	ctx context.Context,
	importNote *importnotemodel.ReqUpdateImportNote) error {
	qtyUpdate := -importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.ReqUpdateDebtSupplier{
		QuantityUpdate: &qtyUpdate,
	}
	if err := repo.supplierStore.UpdateSupplierDebt(
		ctx, importNote.SupplierId, &supplierUpdateDebt,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) FindListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	importNoteDetails, errGetImportNoteDetails :=
		repo.importNoteDetailStore.FindListImportNoteDetail(
			ctx,
			map[string]interface{}{"importNoteId": importNoteId},
		)
	if errGetImportNoteDetails != nil {
		return nil, errGetImportNoteDetails
	}
	return importNoteDetails, nil
}

func (repo *changeStatusImportNoteRepo) HandleProductQuantity(
	ctx context.Context,
	productTotalQuantityNeedUpdate map[string]int) error {
	for key, value := range productTotalQuantityNeedUpdate {
		productUpdate := productmodel.ProductUpdateQuantity{QuantityUpdate: value}
		if err := repo.productStore.UpdateQuantityProduct(
			ctx, key, &productUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}
