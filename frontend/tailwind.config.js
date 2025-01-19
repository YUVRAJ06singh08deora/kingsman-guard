/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    fontFamily: {
      inter: ['Inter', 'serif'],
    },
    extend: {
      colors: {
        primary: '#F9F9FB',
        secondary: '#EAEBEF',
      },
    },
  },
  plugins: [],
}
