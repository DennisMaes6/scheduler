module.exports = {
  purge: ["./src/**/*.svelte", "./src/**/*.html"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    maxHeight: {
      '0': '0vh',
      '1/4': '25vh',
      '1/2': '50vh',
      '3/4': '75vh',
      'full': '100vh',
    },
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
