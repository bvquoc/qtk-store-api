package productmodel

import (
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type Product struct {
	ID          string  `json:"id" gorm:"column:id;primaryKey"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"desc" gorm:"column:desc"`
	Edition     int     `json:"edition" gorm:"column:edition"`
	Quantity    int     `json:"quantity" gorm:"column:qty"`
	Price       float64 `json:"price" gorm:"column:price"`
	SalePrice   float64 `json:"salePrice" gorm:"column:salePrice"`
	PublisherID string  `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   string  `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs string  `json:"categoryIds" gorm:"column:categoryIds"`
	IsActive    bool    `json:"isActive" gorm:"column:isActive"`
}

func (*Product) TableName() string {
	return constants.TblProduct
}

var (
	ErrProductIdInvalid = common.NewCustomError(
		errors.New("id of Product is invalid"),
		"id of Product is invalid",
		"ErrProductIdInvalid",
	)

	ErrProductNameEmpty = common.NewCustomError(
		errors.New("name of Product is empty"),
		"name of Product is empty",
		"ErrProductNameEmpty",
	)

	ErrProductPriceIsLessThanZero = common.NewCustomError(
		errors.New("price of Product is a less than zero"),
		"price of Product must be greater than 0",
		"ErrProductPriceIsLessThanZero",
	)

	ErrProductSalePriceIsLessThanZero = common.NewCustomError(
		errors.New("sale price of Product is less than zero"),
		"sale price of Product must be greater than 0",
		"ErrProductSalePriceIsLessThanZero",
	)

	ErrProductQuantityIsNegativeNumber = common.NewCustomError(
		errors.New("quantity of Product is a negative number"),
		"quantity of Product is a negative number",
		"ErrProductQuantityIsNegativeNumber",
	)

	ErrProductEditionIsNegativeNumber = common.NewCustomError(
		errors.New("edition number of Product is a negative number"),
		"edition number of Product is a negative number",
		"ErrProductEditionIsNegativeNumber",
	)

	ErrProductPublisherIdEmpty = common.NewCustomError(
		errors.New("publisher ID of Product is empty"),
		"publisher ID of Product is empty",
		"ErrProductPublisherIdEmpty",
	)

	ErrProductAuthorIdsEmpty = common.NewCustomError(
		errors.New("author IDs of Product are empty"),
		"author IDs of Product are empty",
		"ErrProductAuthorIdsEmpty",
	)

	ErrProductCategoryIdsEmpty = common.NewCustomError(
		errors.New("category IDs of Product are empty"),
		"category IDs of Product are empty",
		"ErrProductCategoryIdsEmpty",
	)

	ErrProductQtyUpdateInvalid = common.NewCustomError(
		errors.New("quantity need to update for the Product is invalid"),
		"quantity need to update for the Product is invalid",
		"ErrProductQtyUpdateInvalid",
	)
	ErrProductIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Product is duplicate"),
	)
	ErrProductCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Product"),
	)
	ErrProductViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Product"),
	)
)
