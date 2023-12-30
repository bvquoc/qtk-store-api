const httpStatus = require('http-status');
const { Supplier } = require('../models');
const ApiError = require('../utils/ApiError');

const createSupplier = async (supplierBody) => {
  return Supplier.create(supplierBody);
};

const querySuppliers = async (filter, options) => {
  const suppliers = await Supplier.paginate(filter, options);
  return suppliers;
};

const getSupplierById = async (id) => {
  return Supplier.findById(id);
};

const updateSupplierById = async (id, updateBody) => {
  const supplier = await getSupplierById(id);
  if (!supplier) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Supplier not found');
  }
  Object.assign(supplier, updateBody);
  await supplier.save();
  return supplier;
};

const deleteSupplierById = async (id) => {
  const supplier = await getSupplierById(id);
  if (!supplier) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Supplier not found');
  }
  await supplier.remove();
  return supplier;
};

const addProductToSupplier = async (supplierId, productId) => {
  const supplier = await getSupplierById(supplierId);
  supplier.productIds.push(productId);
  await supplier.save();
};

const removeProductFromSupplier = async (supplierId, productId) => {
  const supplier = await getSupplierById(supplierId);
  supplier.productIds.pull(productId);
  await supplier.save();
};

module.exports = {
  createSupplier,
  querySuppliers,
  getSupplierById,
  updateSupplierById,
  deleteSupplierById,

  addProductToSupplier,
  removeProductFromSupplier,
};
