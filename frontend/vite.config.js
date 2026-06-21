import path from 'path'
import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
import tailwindcss from '@tailwindcss/vite'
import { router } from 'sv-router/vite-plugin'; 

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), svelte(), router()],

  resolve: {
    alias: {
      $lib: path.resolve("./src/lib"),
    },
  },
})
