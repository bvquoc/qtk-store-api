const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const categorySchema = mongoose.Schema(
  {
    name: {
      type: String,
      unique: true, // This enforces uniqueness on the 'name' field
      required: true, // This makes sure the field is mandatory
      trim: true, // This trims whitespace from the beginning and end of the value
    },
  },
  {
    timestamps: true, // This adds 'createdAt' and 'updatedAt' timestamps
  }
);

// add plugin that converts mongoose to json
categorySchema.plugin(toJSON);
categorySchema.plugin(paginate);

/**
 * @typedef Category
 */
const Category = mongoose.model('Category', categorySchema);

module.exports = Category;
