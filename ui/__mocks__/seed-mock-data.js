#!/usr/bin/env node
const request = require('superagent');
const Throttle = require('superagent-throttle')
const uuid = require('uuid/v4');
const meow = require('meow');
const _ = require('lodash');
const {lorem} = require('faker');
const cli = meow(`
    Usage
      $ seed-mock-data

    Options
      --number, -n  Number of mock items to create
      --offset, -o  Number to offset
      --rate,   -r  Number of items /1s to create
    Examples
      $ seed-mock-data --number=500
`,
  {
    alias: {
      n: 'number',
      o: 'offset',
      r: 'rate'
    }
  });

const count = cli.flags.number || 500;
const offset = cli.flags.offset || 1;
const rate = cli.flags.rate || 5;
const throttle = new Throttle({
  active: true,     // set false to pause queue
  rate: rate,          // how many requests can be sent every `ratePer`
  ratePer: 1000,   // number of ms in which `rate` requests may be sent
  concurrent: 10     // how many requests can be sent concurrently
});

_.each(_.range(offset, offset + count), function(iteration) {
  const UUID = uuid();
  request.put(`http://localhost:8080/methode/${UUID}`)
  .set('Content-Type', 'application/json')
  .set('X-Request-Id', `tid_${pad(iteration, 6)}`)
  .send({
    uuid: UUID,
    publishReference: `tid_${pad(iteration, 6)}`,
    title: lorem.words(Math.random() * (10 - 3) + 3)
  })
  .use(throttle.plugin())
  .end(err => console.log(err ? err : 'done ' + iteration)); //eslint-disable-line no-console
});

function pad(n, width, z) {
  z = z || '0';
  n = n + '';
  return n.length >= width ? n : new Array(width - n.length + 1).join(z) + n;
}
