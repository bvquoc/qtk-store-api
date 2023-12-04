package importnoterepo

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
)

type SeeDetailImportNoteStore interface {
	ListImportNoteDetail(
		ctx context.Context,
		importNoteId string,
		paging *common.Paging) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type FindImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*importnotemodel.ImportNote, error)
}

type seeDetailImportNoteRepo struct {
	importNoteStore       FindImportNoteStore
	importNoteDetailStore SeeDetailImportNoteStore
}

func NewSeeDetailImportNoteRepo(
	importNoteStore FindImportNoteStore,
	importNoteDetailStore SeeDetailImportNoteStore) *seeDetailImportNoteRepo {
	return &seeDetailImportNoteRepo{
		importNoteDetailStore: importNoteDetailStore,
		importNoteStore:       importNoteStore,
	}
}

func (repo *seeDetailImportNoteRepo) SeeDetailImportNote(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) (*importnotemodel.ResDetailImportNote, error) {
	importNote, errImportNote := repo.importNoteStore.FindImportNote(
		ctx,
		map[string]interface{}{
			"id": importNoteId,
		},
		"Supplier", "CreateByUser", "CloseByUser")
	if errImportNote != nil {
		return nil, errImportNote
	}

	resDetailImportNote := importnotemodel.GetResDetailImportNoteFromImportNote(importNote)

	details, errImportNoteDetail := repo.importNoteDetailStore.ListImportNoteDetail(
		ctx,
		importNoteId,
		paging,
	)
	if errImportNoteDetail != nil {
		return nil, errImportNoteDetail
	}

	resDetailImportNote.Details = details

	return resDetailImportNote, nil
}
