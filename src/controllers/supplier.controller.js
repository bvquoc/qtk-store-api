const httpStatus = require('http-status');
const pick = require('../utils/pick');
// const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { supplierService } = require('../services');

const createSupplier = catchAsync(async (req, res) => {
  const category = await supplierService.createSupplier(req.body);
  res.status(httpStatus.CREATED).send(category);
});

const getSuppliers = catchAsync(async (req, res) => {
  const filter = pick(req.query, ['name', 'phone', 'email']);
  const options = pick(req.query, ['sortBy', 'limit', 'page']);
  const result = await supplierService.querySuppliers(filter, options);
  res.send(result);
});

const updateSupplier = catchAsync(async (req, res) => {
  const supplier = await supplierService.updateSupplierById(req.params.supplierId, req.body);
  res.send(supplier);
});

const deleteSupplier = catchAsync(async (req, res) => {
  const supplier = await supplierService.deleteSupplierById(req.params.supplierId);
  res.send(supplier);
});

module.exports = {
  createSupplier,
  getSuppliers,
  updateSupplier,
  deleteSupplier,
};
