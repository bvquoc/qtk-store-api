const express = require('express');
const uploadRoute = require('./upload.route');
const docsRoute = require('./docs.route');
const authRoute = require('./auth.route');
const userRoute = require('./user.route');
const categoryRoute = require('./category.route');
const customerRoute = require('./customer.route');
const supplierRoute = require('./supplier.route');
const productRoute = require('./product.route');
const config = require('../../config/config');

const router = express.Router();

const defaultRoutes = [
  {
    path: '/upload',
    route: uploadRoute,
  },
  {
    path: '/static',
    route: express.static('uploads'),
  },
  {
    path: '/auth',
    route: authRoute,
  },
  {
    path: '/users',
    route: userRoute,
  },
  {
    path: '/categories',
    route: categoryRoute,
  },
  {
    path: '/customers',
    route: customerRoute,
  },
  {
    path: '/suppliers',
    route: supplierRoute,
  },
  {
    path: '/products',
    route: productRoute,
  },
];

const devRoutes = [
  {
    path: '/docs',
    route: docsRoute,
  },
];

defaultRoutes.forEach((route) => {
  router.use(route.path, route.route);
});

/* istanbul ignore next */
if (config.env === 'development') {
  devRoutes.forEach((route) => {
    router.use(route.path, route.route);
  });
}

module.exports = router;
