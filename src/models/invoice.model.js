const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');
const { Schema } = mongoose;

const invoiceSchema = new Schema(
  {
    customer: {
      type: Schema.Types.ObjectId,
      ref: 'Customer',
      required: true,
    },
    items: [
      {
        _id: false,
        productId: {
          type: Schema.Types.ObjectId,
          ref: 'Product',
          required: true,
        },
        quantity: {
          type: Number,
          required: true,
          min: 1,
        },
        price: {
          type: Number,
          required: true,
        },
        sumPrice: {
          type: Number,
          required: true,
        },
      },
    ],
    totalAmount: {
      type: Number,
      required: true,
    },
    dateIssued: {
      type: Date,
      default: Date.now,
    },
    status: {
      type: String,
      enum: ['pending', 'paid', 'cancelled'],
      default: 'pending',
    },
  },
  {
    timestamps: true, // Adds createdAt and updatedAt timestamps
  }
);

invoiceSchema.plugin(toJSON);
invoiceSchema.plugin(paginate);

const Invoice = mongoose.model('Invoice', invoiceSchema);

module.exports = Invoice;
