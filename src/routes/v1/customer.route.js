const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const customerController = require('../../controllers/customer.controller');
const customerValidation = require('../../validations/customer.validation');

const router = express.Router();

router
  .route('/')
  .post(auth('createCustomer'), validate(customerValidation.createCustomer), customerController.createCustomer)
  .get(auth('getCustomers'), validate(customerValidation.getCustomers), customerController.getCustomers);

router
  .route('/:customerId')
  .get(auth('getCustomers'), validate(customerValidation.getCustomer), customerController.getCustomer)
  .patch(auth('updateCustomer'), validate(customerValidation.updateCustomer), customerController.updateCustomer)
  .delete(auth('deleteCustomer'), validate(customerValidation.deleteCustomer), customerController.deleteCustomer);

module.exports = router;
