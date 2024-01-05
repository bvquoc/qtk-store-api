// eslint-disable-next-line import/no-extraneous-dependencies
const ip = require('ip');
const config = require('../config/config');

const uploadFile = (file) => {
  let ipAddr = `localhost`;
  if (config.env === 'production') {
    ipAddr = ip.address();
  }

  const res = {
    ...file,
    url: `http://${ipAddr}:3000/v1/static/${file.filename}`,
  };
  return res;
};

module.exports = {
  uploadFile,
};
