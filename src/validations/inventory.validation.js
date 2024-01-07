const Joi = require('joi');
const { objectId } = require('./custom.validation');

const importProducts = {
  body: Joi.object().keys({
    products: Joi.array()
      .items(
        Joi.object().keys({
          id: Joi.string().custom(objectId).required(),
          quantity: Joi.number().integer().greater(0).required(),
          importPrice: Joi.number().integer().greater(0).required(),
          receivedDate: Joi.date().default(Date.now()),
          expiryDate: Joi.date().greater(Date.now()).required(),
        })
      )
      .min(1)
      .required(),
  }),
};

const getInventoryItems = {
  query: Joi.object().keys({
    name: Joi.string(),
    sortBy: Joi.string(),
    limit: Joi.number().integer(),
    page: Joi.number().integer(),
  }),
};

module.exports = {
  importProducts,
  getInventoryItems,
};
