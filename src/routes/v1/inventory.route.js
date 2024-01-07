const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const inventoryValidate = require('../../validations/inventory.validation');
const inventoryController = require('../../controllers/inventory.controller');

const router = express.Router();

router.route('/').get(auth('getProducts'), validate(inventoryValidate.getProducts), (req, res) => res.send('get products'));

// import products
router
  .route('/import-products')
  .post(
    auth('getProducts', 'importProducts'),
    validate(inventoryValidate.importProducts),
    inventoryController.importProducts
  )
  .get(
    auth('getProducts', 'importProducts'),
    validate(inventoryValidate.getImportProductsNotes),
    inventoryController.getImportProductsNotes
  );

router
  .route('/import-products/:importId/status')
  .patch(
    auth('getProducts', 'importProducts'),
    validate(inventoryValidate.updateImportProductsStatus),
    inventoryController.updateImportProductsStatus
  );

module.exports = router;
