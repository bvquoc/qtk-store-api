/* eslint-disable camelcase */
// const userWithCustomerAccess = ['getCustomers', 'createCustomer', 'updateCustomer'];
// const adminWithCustomerAccess = ['getCustomers', 'createCustomer', 'updateCustomer', "deleteCustomer"];

const roleUser_Users = [];
const roleAdmin_Users = ['getUsers', 'manageUsers'];

const roleUser_Categories = ['getCategories'];
const roleAdmin_Categories = ['getCategories', 'manageCategories'];

const roleUser_Customers = ['getCustomers', 'createCustomer', 'updateCustomer'];
const roleAdmin_Customers = ['getCustomers', 'createCustomer', 'updateCustomer', 'deleteCustomer'];

const roleUser_Suppliers = ['getSuppliers'];
const roleAdmin_Suppliers = ['getSuppliers', 'createSupplier', 'updateSupplier', 'deleteSupplier'];

const roleUser_Products = ['getProducts'];
const roleAdmin_Products = ['getProducts', 'createProduct', 'updateProduct', 'deleteProduct'];

const allRoles = {
  user: [...roleUser_Users, ...roleUser_Categories, ...roleUser_Customers, ...roleUser_Suppliers, ...roleUser_Products],
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
