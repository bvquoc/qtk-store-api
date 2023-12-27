const Joi = require('joi');
const { objectId, phoneNumber } = require('./custom.validation');

const createCustomer = {
  body: Joi.object().keys({
    name: Joi.string().required(),
    email: Joi.string().required().email(),
    phone: Joi.string().required().custom(phoneNumber),
    gender: Joi.string().required().valid('Male', 'Female', 'Other'),
    birthDate: Joi.date().required(),
    address: Joi.object({
      province: Joi.string().required(),
      district: Joi.string().required(),
      ward: Joi.string().required(),
    }).required(),
  }),
};

const getCustomers = {
  query: Joi.object().keys({
    email: Joi.string(),
    phone: Joi.string().custom(phoneNumber),
    limit: Joi.number().integer(),
    page: Joi.number().integer(),
  }),
};

const getCustomer = {
  params: Joi.object().keys({
    customerId: Joi.string().custom(objectId),
  }),
};

const updateCustomer = {
  params: Joi.object().keys({
    customerId: Joi.string().required(),
  }),
  body: Joi.object()
    .keys({
      name: Joi.string(),
      email: Joi.string().email(),
      phone: Joi.string().custom(phoneNumber),
      gender: Joi.string().valid('Male', 'Female', 'Other'),
      birthDate: Joi.date(),
      address: Joi.object({
        province: Joi.string().required(),
        district: Joi.string().required(),
        ward: Joi.string().required(),
      }),
    })
    .min(1),
};

const deleteCustomer = {
  params: Joi.object().keys({
    customerId: Joi.string().custom(objectId),
  }),
};

const activateCustomer = {
  params: Joi.object().keys({
    customerId: Joi.string().custom(objectId),
  }),
};

const deactivateCustomer = {
  params: Joi.object().keys({
    customerId: Joi.string().custom(objectId),
  }),
};

module.exports = {
  createCustomer,
  getCustomers,
  getCustomer,
  updateCustomer,
  deleteCustomer,
  activateCustomer,
  deactivateCustomer,
};
