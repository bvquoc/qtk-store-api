const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const productSchema = mongoose.Schema(
  {
    active: {
      type: Boolean,
      default: true,
    },
    name: {
      type: String,
      unique: true,
      required: true,
      trim: true,
    },
    description: {
      type: String,
      trim: true,
      default: '',
    },
    price: {
      type: Number,
      required: true,
    },
    unit: {
      type: String,
      required: true,
      default: 'VND',
    },
    supplier: {
      type: mongoose.SchemaTypes.ObjectId,
      ref: 'Supplier',
    },
    category: {
      type: mongoose.SchemaTypes.ObjectId,
      ref: 'Category',
    },
    images: [
      {
        type: String,
      },
    ],
  },
  {
    timestamps: true,
  }
);

// add plugin that converts mongoose to json
productSchema.plugin(toJSON);
productSchema.plugin(paginate);

/**
 * @typedef Product
 */
const Product = mongoose.model('Product', productSchema);

module.exports = Product;
