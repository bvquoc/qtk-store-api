const httpStatus = require('http-status');
const { Invoice } = require('../models');
const ApiError = require('../utils/ApiError');

const createInvoice = async (body) => {
  return Invoice.create(body);
};

module.exports = {
  createInvoice,
};
