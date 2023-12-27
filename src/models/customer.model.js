const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const customerSchema = mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
      trim: true,
    },
  },
  {
    timestamps: true,
  }
);

customerSchema.plugin(toJSON);
customerSchema.plugin(paginate);

const Customer = mongoose.model('Customer', customerSchema);
module.exports = Customer;
