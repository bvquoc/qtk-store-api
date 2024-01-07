const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const inventoryValidate = require('../../validations/inventory.validation');
// const productController = require('../../controllers/product.controller');

const router = express.Router();

router.route('/').get(auth('getProducts'), validate(inventoryValidate.getProducts), (req, res) => res.send('get products'));

router
  .route('/importProducts')
  .post(auth('getProducts', 'importProducts'), validate(inventoryValidate.importProducts), (req, res) =>
    res.send('import products')
  );

module.exports = router;
