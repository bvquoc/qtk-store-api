const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const categoryController = require('../../controllers/category.controller');
const categoryValidation = require('../../validations/category.validation');

const router = express.Router();

router
  .route('/')
  .post(auth('createSupplier'), validate(categoryValidation.createCategory), categoryController.createCategory)
  .get(auth('getSuppliers'), validate(categoryValidation.getCategories), categoryController.getCategories);

// router
//   .route('/:categoryId')
//   .patch(auth('manageCategories'), validate(categoryValidation.updateCategory), categoryController.updateCategory)
//   .delete(auth('manageCategories'), validate(categoryValidation.deleteCategory), categoryController.deleteCategory);

module.exports = router;
