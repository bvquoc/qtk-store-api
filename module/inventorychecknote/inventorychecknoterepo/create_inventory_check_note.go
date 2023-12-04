package inventorychecknoterepo

import (
	"context"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"qtk-store-api/module/product/productmodel"
)

type CreateInventoryCheckNoteStore interface {
	CreateInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.ReqCreateInventoryCheckNote,
	) error
}

type CreateInventoryCheckNoteDetailStore interface {
	CreateListInventoryCheckNoteDetail(
		ctx context.Context,
		data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate,
	) error
}

type UpdateProductStore interface {
	UpdateQuantityProduct(
		ctx context.Context,
		id string,
		data *productmodel.ProductUpdateQuantity,
	) error
	FindProduct(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type createInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       CreateInventoryCheckNoteStore
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
	productStore                  UpdateProductStore
}

func NewCreateInventoryCheckNoteRepo(
	inventoryCheckNoteStore CreateInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore,
	productStore UpdateProductStore) *createInventoryCheckNoteRepo {
	return &createInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
		productStore:                  productStore,
	}
}

func (repo *createInventoryCheckNoteRepo) HandleInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	if err := repo.inventoryCheckNoteStore.CreateInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	if err := repo.inventoryCheckNoteDetailStore.CreateListInventoryCheckNoteDetail(
		ctx, data.Details,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInventoryCheckNoteRepo) HandleProductQuantity(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	qtyDiff := 0
	qtyAfter := 0
	for i, value := range data.Details {
		product, errGetProduct := repo.productStore.FindProduct(
			ctx, map[string]interface{}{"id": value.ProductId})
		if errGetProduct != nil {
			return errGetProduct
		}

		data.Details[i].Initial = product.Quantity
		data.Details[i].Final = product.Quantity + value.Difference
		qtyDiff += value.Difference
		qtyAfter += data.Details[i].Final

		if data.Details[i].Final < 0 {
			return inventorychecknotemodel.ErrInventoryCheckNoteModifyQuantityIsInvalid
		}

		productUpdate := productmodel.ProductUpdateQuantity{QuantityUpdate: value.Difference}

		if err := repo.productStore.UpdateQuantityProduct(
			ctx, value.ProductId, &productUpdate,
		); err != nil {
			return err
		}
	}

	data.QuantityDifferent = qtyDiff
	data.QuantityAfterAdjust = qtyAfter
	return nil
}
