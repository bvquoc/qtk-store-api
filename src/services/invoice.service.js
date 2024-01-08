const httpStatus = require('http-status');
const { Invoice } = require('../models');
const ApiError = require('../utils/ApiError');

const createInvoice = async (body) => {
  return Invoice.create(body);
};

const queryInvoices = async (filter, options) => {
  const invoices = await Invoice.paginate(filter, options);
  return invoices;
};

const getInvoiceById = async (id) => {
  return Invoice.findById(id);
};

const getInvoicesByStatus = async (status) => {
  return Invoice.find({ status });
};

const getInvoicesByCustomerId = async (customerId) => {
  return Invoice.find({ customerId });
};

const deleteInvoiceById = async (id) => {
  const invoice = await getInvoiceById(id);
  if (!invoice) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Invoice not found');
  }
  await invoice.remove();
  return invoice;
};

module.exports = {
  createInvoice,
  queryInvoices,
  getInvoiceById,
  getInvoicesByStatus,
  getInvoicesByCustomerId,
  deleteInvoiceById,
};
