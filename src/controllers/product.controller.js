const httpStatus = require('http-status');
const pick = require('../utils/pick');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { productService, supplierService, categoryService } = require('../services');

const createProduct = catchAsync(async (req, res) => {
  const { supplierId } = req.body;
  const supplier = await supplierService.getSupplierById(supplierId);
  if (!supplier) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Supplier not found');
  }

  const { categoryIds } = req.body;
  for (let i = 0; i < categoryIds.length; i++) {
    const id = categoryIds[i];
    // eslint-disable-next-line no-await-in-loop
    const category = await categoryService.getCategoryById(id);
    if (!category) {
      throw new ApiError(httpStatus.NOT_FOUND, 'Category not found');
    }
  }

  const product = await productService.createProduct(req.body);

  await supplierService.addProductToSupplier(supplier, product._id);
  for (let i = 0; i < categoryIds.length; i++) {
    const id = categoryIds[i];
    // eslint-disable-next-line no-await-in-loop
    await categoryService.addProductToCategory(id, product._id);
  }

  res.status(httpStatus.CREATED).send(product);
});

const getProducts = catchAsync(async (req, res) => {
  const filter = pick(req.query, ['name']);
  const options = pick(req.query, ['sortBy', 'limit', 'page']);
  const result = await productService.queryProducts(filter, options);

  const products = result.results;
  for (let i = 0; i < products.length; i++) {
    const product = products[i];
    // eslint-disable-next-line no-await-in-loop
    const supplier = await supplierService.getSimpleSupplierById(product.supplierId);
    product.supplier = supplier;
    product.supplierId = undefined;

    const categories = [];
    for (let j = 0; j < product.categoryIds.length; j++) {
      const id = product.categoryIds[j];
      // eslint-disable-next-line no-await-in-loop
      const category = await categoryService.getSimpleCategoryById(id);
      categories.push(category);
    }
    product.categories = categories;
    product.categoryIds = undefined;
  }

  res.send(result);
});

const updateProduct = catchAsync(async (req, res) => {
  const product = await productService.updateProductById(req.params.productId, req.body);
  res.send(product);
});

const deleteProduct = catchAsync(async (req, res) => {
  await productService.deleteProductById(req.params.productId);
  res.status(httpStatus.NO_CONTENT).send();
});

module.exports = {
  createProduct,
  getProducts,
  updateProduct,
  deleteProduct,
};
