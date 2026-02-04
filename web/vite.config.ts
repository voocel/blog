import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      }
    }
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks(id: string) {
          // Core React (required for all pages)
          if (id.includes('node_modules/react-dom') ||
              id.includes('node_modules/react/') ||
              id.includes('node_modules/react-router')) {
            return 'react-vendor'
          }
          // Markdown editor (only for AIChat and PostPage)
          if (id.includes('node_modules/@uiw/react-md-editor') ||
              id.includes('node_modules/@uiw/react-markdown-preview')) {
            return 'md-editor'
          }
          // Gemini AI SDK (only for AIChat)
          if (id.includes('node_modules/@google/genai')) {
            return 'gemini'
          }
          // Animation library
          if (id.includes('node_modules/framer-motion')) {
            return 'framer'
          }
        }
      }
    }
  }
})
