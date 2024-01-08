const httpStatus = require('http-status');
const { InventoryImportNote, InventoryItem } = require('../models');
const ApiError = require('../utils/ApiError');
const { productService } = require('./index');

const importProducts = async (body) => {
  return InventoryImportNote.create(body);
};

const getImportNoteById = async (id) => {
  return InventoryImportNote.findById(id);
};

const updateImportStatus = async (id, status) => {
  const inventoryImportNote = await getImportNoteById(id);
  if (!inventoryImportNote) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Inventory Import Note not found');
  }
  if (inventoryImportNote.status === 'completed') {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Inventory Import Note already imported');
  }
  if (inventoryImportNote.status === 'cancelled') {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Inventory Import Note already cancelled');
  }
  Object.assign(inventoryImportNote, {
    status,
    updatedAt: Date.now(),
  });
  await inventoryImportNote.save();

  if (status === 'completed') {
    await productService.importProductsFromImportNote(inventoryImportNote.products);

    // eslint-disable-next-line no-restricted-syntax
    for (const product of inventoryImportNote.products) {
      // eslint-disable-next-line no-await-in-loop
      const inventoryItem = await InventoryItem.findById(product.id);
      if (inventoryItem) {
        Object.assign(inventoryItem, {
          imports: [...inventoryItem.imports, inventoryImportNote.id],
          quantity: [...inventoryItem.quantity, { quantity: product.quantity, expiryDate: product.expiryDate }],
          totalQuantity: inventoryItem.totalQuantity + product.quantity,
          updatedAt: Date.now(),
        });
        // eslint-disable-next-line no-await-in-loop
        await inventoryItem.save();
      } else {
        // eslint-disable-next-line no-await-in-loop
        await InventoryItem.create({
          _id: product.id,
          imports: [inventoryImportNote.id],
          quantity: [{ quantity: product.quantity, expiryDate: product.expiryDate }],
          totalQuantity: product.quantity,
        });
      }
    }
  }
  return inventoryImportNote;
};

const queryImportNotes = async (filter, options) => {
  const inventoryImportNotes = await InventoryImportNote.paginate(filter, options);
  return inventoryImportNotes;
};

const queryInventoryItems = async (filter, options) => {
  const items = await InventoryItem.paginate(filter, options);
  return items;
};

module.exports = {
  importProducts,
  updateImportStatus,
  queryImportNotes,
  queryInventoryItems,
};
