#!/usr/bin/env node
var HOMEPAGE_POLL_TIMEOUT = 60000,
    HOMEPAGE_POLL_INTERVAL = 5000,
    pollTime = 0,
    request = require('superagent');

function pollForTestServerReady() {
  request.get('http://localhost:' + process.env.PORT + '/').end(function(error, response) {
    if (error || response.status !== 200) {
      console.log('Test server not ready after ' + (pollTime / 1000) + 's');
      if (pollTime >= HOMEPAGE_POLL_TIMEOUT) {
        console.log('Not ready after ' + (HOMEPAGE_POLL_TIMEOUT / 1000) + 's. Aborting.');
        process.exit(1);
      }
      pollTime += HOMEPAGE_POLL_INTERVAL;
      setTimeout(pollForTestServerReady, HOMEPAGE_POLL_INTERVAL);
    }
    else {
      // Server is ready
      console.log('NODE_ENV=' + process.env.NODE_ENV + '. Test server is ready and listening. Running tests...');
      require('nightwatch/bin/runner.js');
    }
  });
}

pollForTestServerReady();
