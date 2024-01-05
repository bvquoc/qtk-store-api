const httpStatus = require('http-status');
const catchAsync = require('../utils/catchAsync');
const { uploadService } = require('../services');
const ApiError = require('../utils/ApiError');

const uploadFile = catchAsync((req, res) => {
  if (!req.file) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'File is required.');
  }

  fileInfo = uploadService.uploadFile(req.file);
  res.send(fileInfo);
});
module.exports = {
  uploadFile,
};
