package importnotebiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotemodel"
)

type ListImportNoteRepo interface {
	ListImportNote(
		ctx context.Context,
		filter *importnotemodel.Filter,
		paging *common.Paging,
	) ([]importnotemodel.ImportNote, error)
}

type listImportNoteBiz struct {
	repo      ListImportNoteRepo
	requester middleware.Requester
}

func NewListImportNoteBiz(
	repo ListImportNoteRepo,
	requester middleware.Requester) *listImportNoteBiz {
	return &listImportNoteBiz{repo: repo, requester: requester}
}

func (biz *listImportNoteBiz) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(constants.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	result, err := biz.repo.ListImportNote(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
