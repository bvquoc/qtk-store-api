package importnoterepo

import (
	"context"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
	"qtk-store-api/module/product/productmodel"
	"qtk-store-api/module/supplier/suppliermodel"
)

type CreateImportNoteStore interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ReqCreateImportNote,
	) error
}

type CreateImportNoteDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []importnotedetailmodel.ImportNoteDetailCreate,
	) error
}

type UpdatePriceProductStore interface {
	FindProduct(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
	UpdatePriceProduct(
		ctx context.Context,
		id string,
		data *productmodel.ProductUpdatePrice,
	) error
}

type CheckSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
}

type createImportNoteRepo struct {
	importNoteStore       CreateImportNoteStore
	importNoteDetailStore CreateImportNoteDetailStore
	productStore          UpdatePriceProductStore
	supplierStore         CheckSupplierStore
}

func NewCreateImportNoteRepo(
	importNoteStore CreateImportNoteStore,
	importNoteDetailStore CreateImportNoteDetailStore,
	productStore UpdatePriceProductStore,
	supplierStore CheckSupplierStore) *createImportNoteRepo {
	return &createImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		productStore:          productStore,
		supplierStore:         supplierStore,
	}
}

func (repo *createImportNoteRepo) CheckProduct(
	ctx context.Context,
	productId string) error {
	if _, err := repo.productStore.FindProduct(
		ctx, map[string]interface{}{"id": productId},
	); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) CheckSupplier(
	ctx context.Context,
	supplierId string) error {
	if _, err := repo.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId},
	); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ReqCreateImportNote) error {
	if err := repo.importNoteStore.CreateImportNote(ctx, data); err != nil {
		return err
	}
	if err := repo.importNoteDetailStore.CreateListImportNoteDetail(
		ctx,
		data.ImportNoteDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) UpdatePriceProduct(
	ctx context.Context,
	productId string,
	price float32) error {
	productUpdatePrice := productmodel.ProductUpdatePrice{
		Price: &price,
	}

	if err := repo.productStore.UpdatePriceProduct(
		ctx, productId, &productUpdatePrice,
	); err != nil {
		return err
	}
	return nil
}
