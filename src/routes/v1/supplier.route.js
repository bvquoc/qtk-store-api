const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const supplierContoller = require('../../controllers/supplier.controller');
const supplierValidation = require('../../validations/supplier.validation');

const router = express.Router();

router
  .route('/')
  .post(auth('createSupplier'), validate(supplierValidation.createSupplier), supplierContoller.createSupplier)
  .get(auth('getSuppliers'), validate(supplierValidation.getSuppliers), supplierContoller.getSuppliers);

router
  .route('/:supplierId')
  .patch(auth('updateSupplier'), validate(supplierValidation.updateSuppliers), supplierContoller.updateSupplier)
  .delete(auth('deleteSupplier'), validate(supplierValidation.deleteSupplier), supplierContoller.deleteSupplier);

module.exports = router;
