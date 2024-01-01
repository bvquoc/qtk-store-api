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
    supplierId: {
      type: mongoose.SchemaTypes.ObjectId,
      ref: 'Supplier',
    },
    supplier: {
      type: Object,
      default: undefined,
    },
    categoryIds: [
      {
        type: mongoose.SchemaTypes.ObjectId,
        ref: 'Category',
      },
    ],
    categories: [
      {
        type: Object,
        default: undefined,
      },
    ],
    images: [
      {
        type: String,
      },
    ],
    quantity: {
      type: Object,
      default: {
        imported: 0,
        inStock: 0,
        sold: 0,
      },
    },
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
