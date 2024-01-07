const httpStatus = require('http-status');
const { Product } = require('../models');
const ApiError = require('../utils/ApiError');

const createProduct = async (productBody) => {
  return Product.create(productBody);
};

const queryProducts = async (filter, options) => {
  const products = await Product.paginate(filter, options);
  return products;
};

const getProductById = async (id) => {
  return Product.findById(id);
};

const updateProductById = async (productId, updateBody) => {
  const product = await getProductById(productId);
  if (!product) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
  }
  Object.assign(product, updateBody);
  await product.save();
  return product;
};

const deleteProductById = async (productId) => {
  const product = await getProductById(productId);
  if (!product) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
  }
  await product.remove();
  return product;
};

const importProductsFromImportNote = async (productList) => {
  for (let i = 0; i < productList.length; i++) {
    const product = productList[i];
    // eslint-disable-next-line no-await-in-loop
    const productInDb = await getProductById(product.id);
    if (productInDb) {
      const tmp = productInDb.quantity;
      tmp.imported += product.quantity;
      tmp.inStock += product.quantity;
      productInDb.quantity = tmp;
      // eslint-disable-next-line no-await-in-loop
      await productInDb.save();
    } else {
      throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
    }
  }
};

module.exports = {
  createProduct,
  queryProducts,
  getProductById,
  updateProductById,
  deleteProductById,

  importProductsFromImportNote,
};
