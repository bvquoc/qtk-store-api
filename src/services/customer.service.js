const httpStatus = require('http-status');
const { Customer } = require('../models');
const ApiError = require('../utils/ApiError');

/**
 * Create a user
 * @param {Object} customerBody
 * @returns {Promise<Customer>}
 */
const createCustomer = async (customerBody) => {
  return Customer.create(customerBody);
};

/**
 * Query for customers
 * @param {Object} filter - Mongo filter
 * @param {Object} options - Query options
 * @param {string} [options.sortBy] - Sort option in the format: sortField:(desc|asc)
 * @param {number} [options.limit] - Maximum number of results per page (default = 10)
 * @param {number} [options.page] - Current page (default = 1)
 * @returns {Promise<void>}
 */
const queryCustomers = async (filter, options) => {
  const customers = await Customer.paginate(filter, options);
  return customers;
};

/**
 * Get customer by id
 * @param customerId
 * @param updateBody
 * @returns {Promise<Customer>}
 */
const getCustomerById = async (id) => {
  return Customer.findById(id);
};

const updateCustomerById = async (customerId, updateBody) => {
  const customer = await getCustomerById(customerId);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }
  Object.assign(customer, updateBody);
  await customer.save();
  return customer;
};

const deleteCustomerById = async (id) => {
  const customer = await getCustomerById(id);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }
  await customer.remove();
  return customer;
};

const activateCustomerById = async (id) => {
  const customer = await getCustomerById(id);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }
  customer.active = true;
  await customer.save();
  return customer;
};

const deactivateCustomerById = async (id) => {
  const customer = await getCustomerById(id);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }
  customer.active = false;
  await customer.save();
  return customer;
};

module.exports = {
  createCustomer,
  queryCustomers,
  getCustomerById,
  updateCustomerById,
  deleteCustomerById,
  activateCustomerById,
  deactivateCustomerById,
};
