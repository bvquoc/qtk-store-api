package inventorychecknotebiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"
)

type CreateInventoryCheckNoteRepo interface {
	HandleInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error
	HandleProductQuantity(
		ctx context.Context,
		data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error
}

type createInventoryCheckNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateInventoryCheckNoteRepo
	requester middleware.Requester
}

func NewCreateInventoryCheckNoteBiz(
	gen generator.IdGenerator,
	repo CreateInventoryCheckNoteRepo,
	requester middleware.Requester) *createInventoryCheckNoteBiz {
	return &createInventoryCheckNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInventoryCheckNoteBiz) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	if !biz.requester.IsHasFeature(constants.InventoryCheckNoteCreateFeatureCode) {
		return inventorychecknotemodel.ErrInventoryCheckNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInventoryCheckNoteId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleProductQuantity(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleInventoryCheckNoteId(
	gen generator.IdGenerator,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	id, errGenerateId := gen.IdProcess(data.Id)
	if errGenerateId != nil {
		return errGenerateId
	}
	data.Id = id

	for i := range data.Details {
		data.Details[i].InventoryCheckNoteId = *id
	}

	return nil
}
