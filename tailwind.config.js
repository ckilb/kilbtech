/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./asset/*.*",
    "./tpl/**/*.tmpl",
    "./tpl/**/**/*.tmpl"
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

