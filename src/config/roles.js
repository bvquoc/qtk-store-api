/* eslint-disable camelcase */

const roleUser_Users = [];
const roleStaff_Users = ['getUsers'];
const roleAdmin_Users = ['getUsers', 'manageUsers'];

const roleUser_Uploads = ['upload'];
const roleStaff_Uploads = ['upload'];
const roleAdmin_Uploads = ['upload'];

const roleUser_Categories = ['getCategories'];
const roleStaff_Categories = ['getCategories', 'manageCategories'];
const roleAdmin_Categories = ['getCategories', 'manageCategories'];

const roleUser_Customers = [];
const roleStaff_Customers = ['getCustomers', 'createCustomer', 'updateCustomer'];
const roleAdmin_Customers = ['getCustomers', 'createCustomer', 'updateCustomer', 'deleteCustomer'];

const roleUser_Suppliers = ['getSuppliers'];
const roleStaff_Suppliers = ['getSuppliers'];
const roleAdmin_Suppliers = ['getSuppliers', 'createSupplier', 'updateSupplier', 'deleteSupplier'];

const roleUser_Products = ['getProducts'];
const roleStaff_Products = ['getProducts'];
const roleAdmin_Products = ['getProducts', 'createProduct', 'updateProduct', 'deleteProduct'];

const roleUser_Inventories = [];
const roleStaff_Inventories = ['importProducts', 'getInventoryItems'];
const roleAdmin_Inventories = ['importProducts', 'getInventoryItems'];

const roleUser_Invoices = ['getOwnInvoices'];
const roleStaff_Invoices = ['getInvoices', 'createInvoice', 'updateInvoice'];
const roleAdmin_Invoices = ['getInvoices', 'createInvoice', 'updateInvoice', 'deleteInvoice'];

const allRoles = {
  user: [
    ...roleUser_Users,
    ...roleUser_Uploads,
    ...roleUser_Categories,
    ...roleUser_Customers,
    ...roleUser_Suppliers,
    ...roleUser_Products,
    ...roleUser_Inventories,
    ...roleUser_Invoices,
  ],
  staff: [
    ...roleStaff_Users,
    ...roleStaff_Uploads,
    ...roleStaff_Categories,
    ...roleStaff_Customers,
    ...roleStaff_Suppliers,
    ...roleStaff_Products,
    ...roleStaff_Inventories,
    ...roleStaff_Invoices,
  ],
  admin: [
    ...roleAdmin_Users,
    ...roleAdmin_Uploads,
    ...roleAdmin_Categories,
    ...roleAdmin_Customers,
    ...roleAdmin_Suppliers,
    ...roleAdmin_Products,
    ...roleAdmin_Inventories,
    ...roleAdmin_Invoices,
  ],
};

const roles = Object.keys(allRoles);
const roleRights = new Map(Object.entries(allRoles));

module.exports = {
  roles,
  roleRights,
};
