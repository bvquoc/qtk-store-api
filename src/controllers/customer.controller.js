const httpStatus = require('http-status');
const pick = require('../utils/pick');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { customerService } = require('../services');

const createCustomer = catchAsync(async (req, res) => {
  const customer = await customerService.createCustomer(req.body);
  res.status(httpStatus.CREATED).send(customer);
});

const getCustomers = catchAsync(async (req, res) => {
  const filter = pick(req.query, ['name', 'email', 'phone']);
  const options = pick(req.query, ['limit', 'page']);
  const result = await customerService.queryCustomers(filter, options);
  res.send(result);
});

const getCustomer = catchAsync(async (req, res) => {
  const customer = await customerService.getCustomerById(req.params.customerId);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }
  res.send(customer);
});

const updateCustomer = catchAsync(async (req, res) => {
  const customer = await customerService.updateCustomerById(req.params.customerId, req.body);
  res.send(customer);
});

const deleteCustomer = catchAsync(async (req, res) => {
  await customerService.deleteCustomerById(req.params.customerId);
  res.status(httpStatus.NO_CONTENT).send();
});

const activateCustomer = catchAsync(async (req, res) => {
  await customerService.activateCustomerById(req.params.customerId);
  res.status(httpStatus.NO_CONTENT).send();
});

const deactivateCustomer = catchAsync(async (req, res) => {
  await customerService.deactivateCustomerById(req.params.customerId);
  res.status(httpStatus.NO_CONTENT).send();
});

module.exports = {
  createCustomer,
  getCustomers,
  getCustomer,
  updateCustomer,
  deleteCustomer,
  activateCustomer,
  deactivateCustomer,
};
