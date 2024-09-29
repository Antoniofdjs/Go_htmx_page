/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["*.{html,js,tmpl}", "./templates/**/*.templ", "./static/styles/output.css", "./static/styles/input.css"],
  theme: {
    extend: {
      colors: {
        customColor: "#485870",
        btnPrimary: "rgb(31 41 55 / 80%)",
        hoverBtnPrimary: "rgb(188 105 13 / 90%)",
        btnPrimaryText: "rgb(241 245 249)",
        btnAccent: "rgb(188 105 13 / 90%)",
        primary: "rgb(65 76 89)",
        secondary: "rgb(254 252 232)",
      },
      height: {
        "500px": "500px",
      },
    },
  },
  plugins: [],
};
