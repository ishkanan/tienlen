module.exports = {
  plugins: {
    autoprefixer: {},
    'postcss-import': {
      path: __dirname + '/src',
    },
    'postcss-preset-env': {
      stage: 1,
      features: {
        'nesting-rules': true,
      },
    },
    'postcss-custom-properties': {
      preserve: true,
    },
  },
};
