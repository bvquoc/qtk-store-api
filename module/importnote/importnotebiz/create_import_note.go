package importnotebiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotemodel"
)

type CreateImportNoteRepo interface {
	CheckProduct(
		ctx context.Context,
		productId string,
	) error
	CheckSupplier(
		ctx context.Context,
		supplierId string,
	) error
	HandleCreateImportNote(
		ctx context.Context,
		data *importnotemodel.ReqCreateImportNote,
	) error
	UpdatePriceProduct(
		ctx context.Context,
		productId string,
		price float32,
	) error
}

type createImportNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateImportNoteRepo
	requester middleware.Requester
}

func NewCreateImportNoteBiz(
	gen generator.IdGenerator,
	repo CreateImportNoteRepo,
	requester middleware.Requester) *createImportNoteBiz {
	return &createImportNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createImportNoteBiz) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ReqCreateImportNote) error {
	if !biz.requester.IsHasFeature(constants.ImportNoteCreateFeatureCode) {
		return importnotemodel.ErrImportNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	data.Round()

	for _, v := range data.ImportNoteDetails {
		if err := biz.repo.CheckProduct(ctx, v.ProductId); err != nil {
			return err
		}
	}

	if err := handleImportNoteCreateId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CheckSupplier(ctx, data.SupplierId); err != nil {
		return err
	}

	handleTotalPrice(data)

	if err := biz.repo.HandleCreateImportNote(ctx, data); err != nil {
		return err
	}

	for _, v := range data.ImportNoteDetails {
		if v.IsReplacePrice {
			if err := biz.repo.UpdatePriceProduct(
				ctx, v.ProductId, v.Price,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleImportNoteCreateId(
	gen generator.IdGenerator,
	data *importnotemodel.ReqCreateImportNote) error {
	idImportNote, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idImportNote

	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].ImportNoteId = *idImportNote
	}
	return nil
}

func handleTotalPrice(data *importnotemodel.ReqCreateImportNote) {
	var totalPrice float32 = 0
	for _, importNoteDetail := range data.ImportNoteDetails {
		totalPrice += importNoteDetail.Price * importNoteDetail.QuantityImport
	}
	data.TotalPrice = totalPrice
}
