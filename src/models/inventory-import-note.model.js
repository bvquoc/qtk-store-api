const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const productImportInfoSchema = new mongoose.Schema({
  _id: false,
  id: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'Product',
    required: true,
  },
  quantity: {
    type: Number,
    required: true,
    min: 1,
  },
  importPrice: {
    type: Number,
    required: true,
    min: 0,
  },
  receivedDate: {
    type: Date,
    default: Date.now,
  },
  expiryDate: {
    type: Date,
    required: true,
  },
});

const inventoryImportProductSchema = new mongoose.Schema(
  {
    createdBy: {
      type: Object,
      required: true,
    },
    status: {
      type: String,
      default: 'pending',
      enum: ['pending', 'completed', 'cancelled'],
    },
    products: [productImportInfoSchema],
    totalImportPrice: {
      type: Number,
      required: true,
      min: 0,
    },
  },
  {
    timestamps: true,
  }
);

inventoryImportProductSchema.plugin(toJSON);
inventoryImportProductSchema.plugin(paginate);

const InventoryImportNote = mongoose.model('InventoryImportNote', inventoryImportProductSchema);

module.exports = InventoryImportNote;
