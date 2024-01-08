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

module.exports = {
  createInvoice,
  // getInvoices,
};
