const httpStatus = require('http-status');
const pick = require('../utils/pick');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { invoiceService, productService, customerService } = require('../services');

const createInvoice = catchAsync(async (req, res) => {
  const customer = await customerService.getCustomerById(req.body.customerId);
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

const getInvoices = catchAsync(async (req, res) => {
  const filter = pick(req.query, ['status', 'customer']);
  const options = pick(req.query, ['sortBy', 'limit', 'page']);
  const result = await invoiceService.queryInvoices(filter, options);
  res.send(result);
});

const getInvoiceById = catchAsync(async (req, res) => {
  const invoice = await invoiceService.getInvoiceById(req.params.invoiceId);
  if (!invoice) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Invoice not found');
  }
  res.send(invoice);
});

const updateInvoiceStatus = catchAsync(async (req, res) => {
  const invoice = await invoiceService.getInvoiceById(req.params.invoiceId);

  if (!invoice) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Invoice not found');
  }

  if (invoice.status !== 'pending') {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Invoice status is not pending');
  }

  invoice.status = req.body.status;

  if (req.body.status === 'paid') {
    const customer = await customerService.getCustomerById(invoice.customer);
    if (!customer) {
      throw new ApiError(httpStatus.NOT_FOUND, 'Customer not found');
    }

    for (let i = 0; i < invoice.items.length; i++) {
      // eslint-disable-next-line no-await-in-loop
      const item = invoice.items[i];
      // eslint-disable-next-line no-await-in-loop
      const product = await productService.getProductById(item.productId);
      if (!product) {
        throw new ApiError(httpStatus.NOT_FOUND, 'Product not found');
      }

      if (product.quantity.inStock < item.quantity) {
        throw new ApiError(httpStatus.BAD_REQUEST, 'Product quantity is not enough');
      }
      product.quantity.inStock -= item.quantity;
      product.quantity.sold += item.quantity;
      // eslint-disable-next-line no-await-in-loop
      await product.save();
    }
  }

  await invoice.save();
  res.send(invoice);
});

const deleteInvoice = catchAsync(async (req, res) => {
  let invoice = await invoiceService.getInvoiceById(req.params.invoiceId);
  if (!invoice) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Invoice not found');
  }
  if (invoice.status === 'paid') {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Cannot delete paid invoice');
  }
  invoice = await invoiceService.deleteInvoiceById(req.params.invoiceId);
  if (!invoice) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Invoice not found');
  }
  res.status(httpStatus.NO_CONTENT).send();
});

module.exports = {
  createInvoice,
  getInvoices,
  getInvoiceById,
  updateInvoiceStatus,
  deleteInvoice,
};
