const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const invoiceValidate = require('../../validations/invoice.validation');
const invoiceController = require('../../controllers/invoice.controller');

const router = express.Router();

router
  .route('/')
  .get(auth('getInvoices'), validate(invoiceValidate.getInvoices), invoiceController.getInvoices)
  .post(auth('createInvoice'), validate(invoiceValidate.createInvoice), invoiceController.createInvoice);

router
  .route('/:invoiceId')
  .get(auth('getInvoices'), validate(invoiceValidate.getInvoice), invoiceController.getInvoiceById)
  .delete(auth('deleteInvoice'), validate(invoiceValidate.deleteInvoice), invoiceController.deleteInvoice);

router
  .route('/:invoiceId/update-status')
  .patch(auth('updateInvoice'), validate(invoiceValidate.updateInvoiceStatus), invoiceController.updateInvoiceStatus);

module.exports = router;
