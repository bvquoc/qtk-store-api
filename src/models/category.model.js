const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const categorySchema = mongoose.Schema(
  {
    name: {
      type: String,
      unique: true,
      required: true,
      trim: true,
    },
    active: {
      type: Boolean,
      default: true,
    },
    productIds: [
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
categorySchema.plugin(toJSON);
categorySchema.plugin(paginate);

const Category = mongoose.model('Category', categorySchema);

module.exports = Category;
