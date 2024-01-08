const httpStatus = require('http-status');
const pick = require('../utils/pick');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { invoiceService, productService, customerService, userService } = require('../services');

const createInvoice = catchAsync(async (req, res) => {
  const customer = await userService.getUserById(req.body.customerId);
  if (!customer) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
  }

  const data = {
    customer: req.body.customerId,
    totalAmount: 0,
    dateIssued: new Date(),
    status: 'pending',
    items: [],
  };

  for (let i = 0; i < req.body.items.length; i++) {
    // eslint-disable-next-line no-await-in-loop
    const item = req.body.items[i];
    // eslint-disable-next-line no-await-in-loop
    const product = await productService.getProductById(item.productId);
    if (!product) {
      throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
    }
    if (product.quantity.inStock < req.body.items[i].quantity) {
      throw new ApiError(httpStatus.BAD_REQUEST, 'Product quantity is not enough');
    }

    const dbItem = {
      productId: item.productId,
      quantity: item.quantity,
      sumPrice: product.price * item.quantity,
      price: product.price,
    };
    data.totalAmount += dbItem.sumPrice;
    data.items.push(dbItem);
  }

  const invoice = await invoiceService.createInvoice(data);
  res.status(httpStatus.CREATED).send(invoice);
});

module.exports = {
  createInvoice,
};
