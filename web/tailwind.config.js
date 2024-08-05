/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontSize: {
        '4xl': '2.25rem', // Sem line-height
      },
      fontFamily: {
        'sans': ['Poppins', 'sans-serif'],        
      },
    },
  },

  plugins: [],
}