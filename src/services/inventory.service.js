const httpStatus = require('http-status');
const { InventoryImportNote } = require('../models');
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
  }
  return inventoryImportNote;
};

module.exports = {
  importProducts,
  updateImportStatus,
};
