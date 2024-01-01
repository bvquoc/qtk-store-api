const httpStatus = require('http-status');
const { Category } = require('../models');
const ApiError = require('../utils/ApiError');

const createCategory = async (categoryBody) => {
  return Category.create(categoryBody);
};

const queryCategories = async (filter, options) => {
  const categories = await Category.paginate(filter, options);
  return categories;
};

const getCategoryById = async (id) => {
  return Category.findById(id);
};

const getSimpleCategoryById = async (id) => {
  return Category.findById(id).select('_id name');
};

const updateCategoryById = async (categoryId, updateBody) => {
  const category = await getCategoryById(categoryId);
  if (!category) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Category not found');
  }
  Object.assign(category, updateBody);
  await category.save();
  return category;
};

const deleteCategoryById = async (categoryId) => {
  const category = await getCategoryById(categoryId);
  if (!category) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Category not found');
  }
  await category.remove();
  return category;
};

const addProductToCategory = async (categoryId, productId) => {
  const category = await getCategoryById(categoryId);
  category.productIds.push(productId);
  await category.save();
};

const removeProductFromCategory = async (categoryId, productId) => {
  const category = await getCategoryById(categoryId);
  category.productIds.pull(productId);
  await category.save();
};

module.exports = {
  createCategory,
  queryCategories,
  getCategoryById,
  getSimpleCategoryById,
  updateCategoryById,
  deleteCategoryById,

  addProductToCategory,
  removeProductFromCategory,
};
