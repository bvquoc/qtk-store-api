const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const customerSchema = new mongoose.Schema({
  active: {
    type: Boolean,
    default: true,
  },
  name: {
    type: String,
    required: true,
  },
  userId: {
    type: mongoose.SchemaTypes.ObjectId,
    ref: 'User',
    required: true,
    unique: true,
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
  gender: {
    type: String,
    required: true,
    enum: ['Male', 'Female', 'Other'],
  },
  birthDate: {
    type: Date,
    required: true,
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
});

customerSchema.plugin(toJSON);
customerSchema.plugin(paginate);

const Customer = mongoose.model('Customer', customerSchema);
module.exports = Customer;
