/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "#262626",
        secondary: "#32e0b8",
        success: "#4BB543",
        danger: "#FF4136",
        warning: "#FFC300",
        info: "#7FDBFF",
      },
    },
  },
  plugins: [],
};
