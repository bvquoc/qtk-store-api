const Joi = require('joi');
const { objectId, phoneNumber } = require('./custom.validation');

const createSupplier = {
  body: Joi.object().keys({
    name: Joi.string().required(),
    email: Joi.string().required().email(),
    phone: Joi.string().required().custom(phoneNumber),
    taxIdentificationNumber: Joi.string().required(),
    address: Joi.object({
      province: Joi.string().required(),
      district: Joi.string().required(),
      ward: Joi.string().required(),
    }).required(),
  }),
};

const getSuppliers = {
  query: Joi.object().keys({
    name: Joi.string(),
    email: Joi.string(),
    phone: Joi.string().custom(phoneNumber),
    taxIdentificationNumber: Joi.string(),
    sortBy: Joi.string(),
    limit: Joi.number().integer(),
    page: Joi.number().integer(),
  }),
};

const updateSuppliers = {
  params: Joi.object().keys({
    supplierId: Joi.string().required(),
  }),
  body: Joi.object()
    .keys({
      name: Joi.string(),
      phone: Joi.string().custom(phoneNumber),
      address: Joi.object({
        province: Joi.string().required(),
        district: Joi.string().required(),
        ward: Joi.string().required(),
      }),
    })
    .min(1),
};

const deleteSupplier = {
  params: Joi.object().keys({
    supplierId: Joi.string().custom(objectId),
  }),
};

module.exports = {
  createSupplier,
  getSuppliers,
  updateSuppliers,
  deleteSupplier,
};
