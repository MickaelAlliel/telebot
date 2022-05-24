#!/usr/bin/env -S npx ts-node
import 'dotenv/config';
import express = require('express');
import { postgraphile } from 'postgraphile';
import { getDatabaseConnectionConfig, schemas, options, port } from './common';
import { expressjwt } from 'express-jwt';
import { expressJwtSecret, GetVerificationKey } from 'jwks-rsa';

const middleware = postgraphile(
  getDatabaseConnectionConfig(),
  schemas,
  options
);

const issuer = process.env.ISSUER;
const checkJwt = expressjwt({
  secret: expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: process.env.JWKS_URI,
  }) as GetVerificationKey,
  audience: process.env.AUDIENCE,
  issuer,
  algorithms: ['RS256'],
});

const app = express();
app.use(checkJwt);
app.use(middleware);

const server = app.listen(port, () => {
  const address = server.address();
  if (typeof address !== 'string') {
    const href = `http://127.0.0.1:${address.port}${
      options.graphiqlRoute || '/graphiql'
    }`;
    console.log(`PostGraphiQL available at ${href} 🚀`);
  } else {
    console.log(`PostGraphile listening on ${address} 🚀`);
  }
});
