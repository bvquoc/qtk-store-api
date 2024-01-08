const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const productValidation = require('../../validations/product.validation');
const productController = require('../../controllers/product.controller');

const router = express.Router();

router
  .route('/')
  .post(auth('createProduct'), validate(productValidation.createProduct), productController.createProduct)
  .get(validate(productValidation.getProducts), productController.getProducts);

// router
//   .route('/:productId')
//   .get(auth('getProducts'), validate(productValidation.getProductById), userController.getUser)
//   .patch(auth('updateProduct'), validate(productValidation.updateProduct), userController.updateUser)
//   .delete(auth('deleteProduct'), validate(productValidation.deleteProduct), userController.deleteUser);

module.exports = router;
