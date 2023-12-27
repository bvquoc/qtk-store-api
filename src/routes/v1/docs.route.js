const express = require('express');
const swaggerJsdoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');
const swaggerDefinition = require('../../docs/swaggerDef');

const router = express.Router();

const specs = swaggerJsdoc({
  swaggerDefinition,
  apis: ['src/docs/*.yml', 'src/routes/v1/*.js'],
});

router.use('/', swaggerUi.serve);
router.get(
  '/',
  swaggerUi.setup(specs, {
    customJs: 'https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.0.0/swagger-ui-bundle.min.js',
    customCssUrl: 'https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.0.0/swagger-ui.min.css',
  })
);

module.exports = router;
