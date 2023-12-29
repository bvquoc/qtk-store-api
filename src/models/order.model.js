const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const orderSchema = mongoose.Schema(
  {
    active: {
      type: Boolean,
      default: true,
    },
    customer: {
      type: mongoose.SchemaTypes.ObjectId,
      ref: 'Customer',
    },
    products: [
      {
        type: mongoose.SchemaTypes.ObjectId,
        ref: 'Product',
        quantity: {
          type: Number,
          required: true,
          default: 1,
        },
      },
    ],
    totalPrice: {
      type: Number,
      required: true,
    },
    // status: {
    //   type: String,
    //   enum: ['Created', 'Pending', 'Delivering', 'Completed', 'Cancelled'],
    //   default: 'Pending',
    // },
    // paymentMethod: {
    //   type: String,
    //   enum: ['Money', 'Bank Transfer', 'Credit Card'],
    //   default: 'Money',
    // },
  },
  {
    timestamps: true,
  }
);

// add plugin that converts mongoose to json
orderSchema.plugin(toJSON);
orderSchema.plugin(paginate);

/**
 * @typedef Order
 */
const Order = mongoose.model('Order', orderSchema);

module.exports = Order;
