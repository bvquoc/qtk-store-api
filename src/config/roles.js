/* eslint-disable camelcase */

const roleUser_Users = [];
const roleStaff_Users = ['getUsers'];
const roleAdmin_Users = ['getUsers', 'manageUsers'];

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

const allRoles = {
  user: [...roleUser_Users, ...roleUser_Categories, ...roleUser_Customers, ...roleUser_Suppliers, ...roleUser_Products],
  staff: [
    ...roleStaff_Users,
    ...roleStaff_Categories,
    ...roleStaff_Customers,
    ...roleStaff_Suppliers,
    ...roleStaff_Products,
  ],
  admin: [
    ...roleAdmin_Users,
    ...roleAdmin_Categories,
    ...roleAdmin_Customers,
    ...roleAdmin_Suppliers,
    ...roleAdmin_Products,
  ],
};

const roles = Object.keys(allRoles);
const roleRights = new Map(Object.entries(allRoles));

module.exports = {
  roles,
  roleRights,
};
