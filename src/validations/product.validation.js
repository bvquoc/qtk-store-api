const Joi = require('joi');
const { objectId } = require('./custom.validation');

const createProduct = {
  body: Joi.object().keys({
    name: Joi.string().required(),
    description: Joi.string().required(),
    price: Joi.number().required(),
    unit: Joi.string(),
    supplierId: Joi.string().custom(objectId).required(),
    categoryIds: Joi.array().items(Joi.string().custom(objectId).required()).unique().required(),
    images: Joi.array().items(Joi.string()),
  }),
};

const getProducts = {
  query: Joi.object().keys({
    name: Joi.string(),
    sortBy: Joi.string(),
    limit: Joi.number().integer(),
    page: Joi.number().integer(),
  }),
};

const getProductById = {
  params: Joi.object().keys({
    productId: Joi.string().custom(objectId),
  }),
};

const updateProduct = {
  params: Joi.object().keys({
    productId: Joi.string().required(),
  }),
  body: Joi.object()
    .keys({
      name: Joi.string(),
    })
    .min(1),
};

const deleteProduct = {
  params: Joi.object().keys({
    productId: Joi.string().custom(objectId),
  }),
};

module.exports = {
  createProduct,
  getProducts,
  getProductById,
  updateProduct,
  deleteProduct,
};
