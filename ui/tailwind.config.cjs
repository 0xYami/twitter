/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './components/**/*.vue',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.ts',
    './nuxt.config.ts',
  ],
  theme: {
    extend: {
      animation: {
        'spin-fast': 'spin 0.5s linear infinite',
      },
      colors: {
        bunker: {
          50: '#f3f3f4',
          100: '#e8e8e8',
          200: '#c5c5c6',
          300: '#a2a3a4',
          400: '#5c5d60',
          500: '#16181c',
          600: '#141619',
          700: '#111215',
          800: '#0d0e11',
          900: '#0b0c0e',
        },
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
};
