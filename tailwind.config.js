/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: ["./src/template/**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Public Sans', ...defaultTheme.fontFamily.sans],
      },
    }
  },
  plugins: [],
}

