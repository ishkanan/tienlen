/* eslint-disable @typescript-eslint/no-var-requires */

const url = require('url');
const CompressionPlugin = require('compression-webpack-plugin');
const plugins = [new CompressionPlugin({})];

const devTool = process.env.RELEASE === '1' ? 'hidden-source-map' : '#eval-source-map';

module.exports = {
  devServer: {
    host: '0.0.0.0',
    port: 26000,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:27000',
        changeOrigin: true,
      },
    },
    historyApiFallback: {
      app: url.parse('/assets/').pathname,
    },
  },
  configureWebpack: config => {
    config.devtool = devTool;
    config.resolve.alias['~'] = __dirname + '/src';
    config.plugins.push(...plugins);
  },
  chainWebpack: config => {
    config.plugins.delete('prefetch');

    config.plugin('fork-ts-checker').tap(args => {
      args[0].memoryLimit = process.env.TSCHK_MEMORY_LIMIT ? process.env.TSCHK_MEMORY_LIMIT : 2048;
      args[0].workers = process.env.TSCHK_WORKERS ? process.env.TSCHK_WORKERS : 2;
      return args;
    });

    config.plugin('define').tap(args => {
      args[0]['process.env'].GIT_COMMIT_SHA = JSON.stringify(process.env.GIT_COMMIT_SHA) || null;
      return args;
    });
  },
};
