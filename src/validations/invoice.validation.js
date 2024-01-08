const Joi = require('joi');
const { objectId, objectIds } = require('./custom.validation');

const createInvoice = {
  body: Joi.object().keys({
    items: Joi.array()
      .items(
        Joi.object().keys({
          productId: Joi.string().custom(objectId).required(),
          quantity: Joi.number().integer().min(1).required(),
        })
      )
      .min(1)
      .required(),
    customerId: Joi.string().custom(objectId).required(),
    totalAmount: Joi.number().default(0),
    status: Joi.string().valid('pending', 'paid', 'cancelled').default('pending'),
  }),
};

const updateInvoiceStatus = {
  body: Joi.object().keys({
    status: Joi.string().valid('pending', 'paid', 'cancelled').required(),
  }),
  params: Joi.object().keys({
    invoiceId: Joi.string().custom(objectId),
  }),
};

const getInvoices = {
  query: Joi.object().keys({
    customer: Joi.string().custom(objectId),
    status: Joi.string().valid('pending', 'paid', 'cancelled'),
    sortBy: Joi.string(),
    limit: Joi.number().integer(),
    page: Joi.number().integer(),
  }),
};

const getInvoice = {
  params: Joi.object().keys({
    invoiceId: Joi.string().custom(objectId),
  }),
};

const deleteInvoice = {
  params: Joi.object().keys({
    invoiceId: Joi.string().custom(objectId),
  }),
};

module.exports = {
  createInvoice,
  getInvoice,
  getInvoices,
  updateInvoiceStatus,
  deleteInvoice,
};
