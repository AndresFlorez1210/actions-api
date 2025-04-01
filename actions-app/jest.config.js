module.exports = {
  testEnvironment: 'jsdom',
  moduleFileExtensions: ['js', 'jsx', 'json', 'vue'],
  transform: {
    '^.+\\.vue$': '@vue/vue3-jest',
    '^.+\\.js$': 'babel-jest'
  },
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/src/$1'
  },
  setupFiles: ['<rootDir>/jest.setup.js'],
  testMatch: [
    '**/tests/unit/**/*.spec.[jt]s?(x)',
    '**/tests/unit/**/*.test.[jt]s?(x)'
  ],
  transformIgnorePatterns: ['/node_modules/(?!vue)'],
  testEnvironmentOptions: {
    customExportConditions: ['node', 'node-addons']
  }
} 