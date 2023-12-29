const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const supplierSchema = mongoose.Schema(
  {
    active: {
      type: Boolean,
      default: true,
    },
    name: {
      type: String,
      required: true,
    },
    email: {
      type: String,
      required: true,
      unique: true,
    },
    phone: {
      type: String,
      required: true,
      unique: true,
    },
    address: {
      province: {
        type: String,
        required: true,
      },
      district: {
        type: String,
        required: true,
      },
      ward: {
        type: String,
        required: true,
      },
    },
    importedProducts: [
      {
        type: mongoose.SchemaTypes.ObjectId,
        ref: 'Product',
      },
    ],
  },
  {
    timestamps: true,
  }
);

// add plugin that converts mongoose to json
supplierSchema.plugin(toJSON);
supplierSchema.plugin(paginate);

/**
 * @typedef Supplier
 */
const Supplier = mongoose.model('Supplier', supplierSchema);

module.exports = Supplier;
