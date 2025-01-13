import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
      },
    },
    colors: {
      primary: '#DBC500',
      red: '#dc2626',
      green: '#84cc16',
      white: '#ffffff',
      textColor: '#863C2D',
      black: '#000000',
      borderprimary: '#707070'
    }
  },
  plugins: [],
} satisfies Config;
