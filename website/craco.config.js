// craco.config.js
// const CracoSwcPlugin = require('craco-swc');
const { ProvidePlugin } = require('webpack');

module.exports = {
  style: {
    postcss: {
      plugins: [require('tailwindcss'), require('autoprefixer')],
    },
  },
  // webpack: {
  //   plugins: [
  //     new ProvidePlugin({
  //       React: 'react',
  //     }),
  //   ],
  // },
  // plugins: [
  //   {
  //     plugin: CracoSwcPlugin,
  //     options: {
  //       swcLoaderOptions: {
  //         jsc: {
  //           externalHelpers: true,
  //           target: 'es2020',
  //           parser: {
  //             syntax: 'typescript',
  //             tsx: true,
  //             dynamicImport: true,
  //             decorators: false,
  //           },
  //         },
  //       },
  //     },
  //   },
  // ],
};
