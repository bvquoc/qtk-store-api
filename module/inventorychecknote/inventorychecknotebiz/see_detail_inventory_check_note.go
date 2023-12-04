package inventorychecknotebiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"
)

type SeeDetailInventoryCheckNoteRepo interface {
	SeeDetailInventoryCheckNote(
		ctx context.Context,
		inventoryCheckNoteId string,
		paging *common.Paging,
	) (*inventorychecknotemodel.ResDetailInventoryCheckNote, error)
}

type seeDetailInventoryCheckNoteBiz struct {
	repo      SeeDetailInventoryCheckNoteRepo
	requester middleware.Requester
}

func NewSeeDetailImportNoteBiz(
	repo SeeDetailInventoryCheckNoteRepo,
	requester middleware.Requester) *seeDetailInventoryCheckNoteBiz {
	return &seeDetailInventoryCheckNoteBiz{repo: repo, requester: requester}
}

func (biz *seeDetailInventoryCheckNoteBiz) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string,
	paging *common.Paging) (*inventorychecknotemodel.ResDetailInventoryCheckNote, error) {
	if !biz.requester.IsHasFeature(constants.InventoryCheckNoteViewFeatureCode) {
		return nil, inventorychecknotemodel.ErrInventoryCheckNoteViewNoPermission
	}

	inventoryCheckNote, err := biz.repo.SeeDetailInventoryCheckNote(
		ctx,
		inventoryCheckNoteId,
		paging)

	if err != nil {
		return nil, err
	}

	return inventoryCheckNote, nil
}
