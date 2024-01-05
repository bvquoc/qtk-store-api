const express = require('express');
const auth = require('../../middlewares/auth');
const { uploadValidation } = require('../../validations');
const { uploadController } = require('../../controllers');

const router = express.Router();

router.post('/', auth('upload'), uploadValidation.uploadFile, uploadController.uploadFile);

module.exports = router;
