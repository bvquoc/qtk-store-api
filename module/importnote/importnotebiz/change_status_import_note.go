package importnotebiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
)

type ChangeStatusImportNoteRepo interface {
	FindImportNote(
		ctx context.Context,
		importNoteId string) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		importNoteId string,
		data *importnotemodel.ReqUpdateImportNote) error
	CreateSupplierDebt(
		ctx context.Context,
		supplierDebtId string,
		importNote *importnotemodel.ReqUpdateImportNote) error
	UpdateDebtSupplier(
		ctx context.Context,
		importNote *importnotemodel.ReqUpdateImportNote) error
	FindListImportNoteDetail(
		ctx context.Context,
		importNoteId string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
	HandleProductQuantity(
		ctx context.Context,
		productTotalQuantityNeedUpdate map[string]int,
	) error
}

type changeStatusImportNoteRepo struct {
	gen       generator.IdGenerator
	repo      ChangeStatusImportNoteRepo
	requester middleware.Requester
}

func NewChangeStatusImportNoteBiz(
	gen generator.IdGenerator,
	repo ChangeStatusImportNoteRepo,
	requester middleware.Requester) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusImportNoteRepo) ChangeStatusImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ReqUpdateImportNote) error {
	if !biz.requester.IsHasFeature(constants.ImportNoteChangeStatusFeatureCode) {
		return importnotemodel.ErrImportNoteChangeStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	importNote, errGetImportNote := biz.repo.FindImportNote(ctx, importNoteId)
	if errGetImportNote != nil {
		return errGetImportNote
	}
	data.Id = importNoteId
	data.TotalPrice = importNote.TotalPrice
	data.SupplierId = importNote.SupplierId

	if *importNote.Status != importnotemodel.InProgress {
		return importnotemodel.ErrImportNoteClosed
	}

	if *data.Status == importnotemodel.Done {
		supplierDebtId, errGenerateId := biz.gen.GenerateId()
		if errGenerateId != nil {
			return errGenerateId
		}

		if err := biz.repo.CreateSupplierDebt(ctx, supplierDebtId, data); err != nil {
			return err
		}

		if err := biz.repo.UpdateDebtSupplier(ctx, data); err != nil {
			return err
		}

		importNoteDetails, errGetImportNoteDetails := biz.repo.FindListImportNoteDetail(
			ctx,
			importNoteId)
		if errGetImportNoteDetails != nil {
			return errGetImportNoteDetails
		}

		mapProductQuantity := getMapProductTotalQuantityNeedUpdated(importNoteDetails)
		if err := biz.repo.HandleProductQuantity(ctx, mapProductQuantity); err != nil {
			return err
		}
	}
	if err := biz.repo.UpdateImportNote(ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func getMapProductTotalQuantityNeedUpdated(
	importNoteDetails []importnotedetailmodel.ImportNoteDetail) map[string]int {
	result := make(map[string]int)
	for _, v := range importNoteDetails {
		result[v.ProductId] += v.QuantityImport
	}
	return result
}
