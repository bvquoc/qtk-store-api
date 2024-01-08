const httpStatus = require('http-status');
const pick = require('../utils/pick');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { productService, inventoryService } = require('../services');

const importProducts = catchAsync(async (req, res) => {
  const { products } = req.body;
  let totalImportPrice = 0;
  for (let i = 0; i < products.length; i++) {
    const productId = products[i].id;
    // eslint-disable-next-line no-await-in-loop
    const product = await productService.getProductById(productId);
    if (!product) {
      throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
    }

    totalImportPrice += products[i].importPrice * products[i].quantity;
  }

  req.body.createdBy = {
    id: req.user.id,
    name: req.user.name,
  };
  req.body.totalImportPrice = totalImportPrice;

  const importProductNote = await inventoryService.importProducts(req.body);
  res.status(httpStatus.CREATED).send(importProductNote);
});
const updateImportProductsStatus = catchAsync(async (req, res) => {
  const importProductNote = await inventoryService.updateImportStatus(req.params.importId, req.body.status);
  res.send(importProductNote);
});

const getImportProductsNotes = catchAsync(async (req, res) => {
  const filter = pick(req.query, ['status']);
  const options = pick(req.query, ['sortBy', 'limit', 'page']);
  const result = await inventoryService.queryImportNotes(filter, options);
  res.send(result);
});

const getInventoryItems = catchAsync(async (req, res) => {
  const filter = pick(req.query, []);
  const options = pick(req.query, ['sortBy', 'limit', 'page']);
  const result = await inventoryService.queryInventoryItems(filter, options);
  res.send(result);
});

module.exports = {
  importProducts,
  updateImportProductsStatus,
  getImportProductsNotes,
  getInventoryItems,
};
