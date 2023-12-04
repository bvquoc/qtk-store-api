package supplierrepo

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type ListSupplierImportNoteStore interface {
	ListAllImportNoteBySupplier(
		supplierId string,
		filter *filter.SupplierImportFilter,
		ctx context.Context,
		paging *common.Paging) ([]importnotemodel.ImportNote, error)
}

type FindSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
}

type seeSupplierImportNoteRepo struct {
	importNoteStore ListSupplierImportNoteStore
	supplierStore   FindSupplierStore
}

func NewSeeSupplierImportNoteRepo(
	importNoteStore ListSupplierImportNoteStore,
	supplierStore FindSupplierStore) *seeSupplierImportNoteRepo {
	return &seeSupplierImportNoteRepo{
		importNoteStore: importNoteStore,
		supplierStore:   supplierStore,
	}
}

func (biz *seeSupplierImportNoteRepo) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) (*suppliermodel.ResImportNoteSupplier, error) {
	supplier, errSupplier := biz.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId})
	if errSupplier != nil {
		return nil, errSupplier
	}

	resSeeImportNoteSupplier := suppliermodel.GetResSeeImportNoteSupplierFromSupplier(supplier)

	importNotes, errImportNotes := biz.importNoteStore.ListAllImportNoteBySupplier(
		supplierId,
		filter,
		ctx,
		paging,
	)
	if errImportNotes != nil {
		return nil, errImportNotes
	}

	resSeeImportNoteSupplier.ImportHistory = importNotes

	return resSeeImportNoteSupplier, nil
}
