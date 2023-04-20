/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./src/template/**/*.html",
    "./src/asset/*.js"
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Work Sans', ...defaultTheme.fontFamily.sans],
      },
    }
  },
  plugins: [],
}

