/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["*.{html,js}", "./htmlTemplates", "./static/styles/output.css", "./static/styles/input.css"],
  theme: {
    extend: {
      colors: {
        customColor: "#485870",
      },
    },
  },
  plugins: [],
};
