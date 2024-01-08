const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const invoiceValidate = require('../../validations/invoice.validation');
const invoiceController = require('../../controllers/invoice.controller');

const router = express.Router();

router
  .route('/')
  .get(auth('getInvoices'), validate({}), (req, res, next) => {
    try {
      res.status(200).json({ message: 'getInvoices' });
    } catch (error) {
      next(error);
    }
  })
  .post(auth('createInvoice'), validate(invoiceValidate.createInvoice), invoiceController.createInvoice);

router
  .route('/:invoiceId')
  .get(auth('getInvoices'), validate({}), (req, res, next) => {
    try {
      res.status(200).json({ message: 'getInvoice' });
    } catch (error) {
      next(error);
    }
  })
  .patch(auth('updateInvoices'), validate({}), (req, res, next) => {
    try {
      res.status(200).json({ message: 'updateInvoice' });
    } catch (error) {
      next(error);
    }
  })
  .delete(auth('deleteInvoices'), validate({}), (req, res, next) => {
    try {
      res.status(200).json({ message: 'deleteInvoice' });
    } catch (error) {
      next(error);
    }
  });

module.exports = router;
