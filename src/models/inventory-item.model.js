const mongoose = require('mongoose');
const { toJSON, paginate } = require('./plugins');

const inventoryItemSchema = new mongoose.Schema(
  {
    _id: {
      type: mongoose.SchemaTypes.ObjectId,
      required: true,
      trim: true,
    },
    imports: [
      {
        type: mongoose.SchemaTypes.ObjectId,
        ref: 'InventoryImportNote',
      },
    ],

    quantity: [
      {
        _id: false,
        quantity: {
          type: Number,
          required: true,
          trim: true,
        },
        expiryDate: {
          type: Date,
          required: true,
          trim: true,
        },
      },
    ],
    totalQuantity: {
      type: Number,
      required: true,
      trim: true,
    },
  },

  {
    timestamps: true,
  }
);

inventoryItemSchema.plugin(toJSON);
inventoryItemSchema.plugin(paginate);

const InventoryItem = mongoose.model('InventoryItem', inventoryItemSchema);

module.exports = InventoryItem;
