// tailwind.config.js
module.exports = {
  purge: [
    '/static/cms/assets/templates/*.html',
    '/static/cms/assets/templates/*.js',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {},
  plugins: [],
}