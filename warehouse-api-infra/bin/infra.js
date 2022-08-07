#!/usr/bin/env node

const { App } = require('aws-cdk-lib');
const { WarehouseStack } = require('../lib/infra-stack');

const app = new App();
new WarehouseStack(app, 'WarehouseStack', {
  env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },
});
