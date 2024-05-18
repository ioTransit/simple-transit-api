/** @type {import('tailwindcss').Config} */

import colors from 'tailwindcss/colors'
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    extend: {
      ...colors
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
