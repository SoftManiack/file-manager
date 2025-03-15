module.exports = {
    root: true,
    env: {
      node: true,
    },
    extends: [
      'eslint:recommended',
      'plugin:vue/essential', // или 'plugin:vue/strongly-recommended' для более строгих правил
    ],
    rules: {
        'vue/no-v-model-argument': 'off',
    },
  };