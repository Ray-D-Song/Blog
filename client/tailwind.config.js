/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: [],
  theme: {
    extend: {
      backgroundColor: {
        'code-dark': '#24283B',
        'code-light': '#F6F8FA',
        'curtain-light': '#F4F6F9',
        'curtain-dark': '#23283C',
        'select-dark': '#353535',
        'select-light': '#AFC4D5'
      },
      width: {
        '9/10': '90%',
        '7/10': '70%'
      }
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
  future: {
    hoverOnlyWhenSupported: true
  }
}

