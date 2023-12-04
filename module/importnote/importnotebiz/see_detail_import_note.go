package importnotebiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotemodel"
)

type SeeDetailImportNoteRepo interface {
	SeeDetailImportNote(
		ctx context.Context,
		importNoteId string,
		paging *common.Paging,
	) (*importnotemodel.ResDetailImportNote, error)
}

type seeDetailImportNoteBiz struct {
	repo      SeeDetailImportNoteRepo
	requester middleware.Requester
}

func NewSeeDetailImportNoteBiz(
	repo SeeDetailImportNoteRepo,
	requester middleware.Requester) *seeDetailImportNoteBiz {
	return &seeDetailImportNoteBiz{repo: repo, requester: requester}
}

func (biz *seeDetailImportNoteBiz) SeeDetailImportNote(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) (*importnotemodel.ResDetailImportNote, error) {
	if !biz.requester.IsHasFeature(constants.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	importNote, err := biz.repo.SeeDetailImportNote(
		ctx,
		importNoteId,
		paging)

	if err != nil {
		return nil, err
	}

	return importNote, nil
}
