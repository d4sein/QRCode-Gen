module.exports = {
  root: true,
  env: {
    node: true
  },
  'extends': [
    'plugin:vue/essential',
    'eslint:recommended',
    '@vue/typescript/recommended'
  ],
  parserOptions: {
    ecmaVersion: 2020
  },
  rules: {
    'semi': ['error', 'never'],
    'max-len': ['error', { 'code': 90, 'ignoreStrings': true, 'ignoreTemplateLiterals': true }],
    '@typescript-eslint/no-non-null-assertion': 'off'
  }
}
