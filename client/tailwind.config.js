/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./App.{js,ts,tsx}', './app/**/*.{js,ts,tsx}'],

  presets: [require('nativewind/preset')],
  theme: {
    extend: {
      colors: {
        brand: '#FFF1CA',
        background: '#FFB823',
        surface: '#708A58',
        accent: '#2D4F2B',
      },
    },
  },
  plugins: [],
};
