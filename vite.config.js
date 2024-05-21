import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
  root: 'scripts', // Adjust according to your project structure
  build: {
    outDir: 'dist', // The output directory for the build
    rollupOptions: {
      input: {
        index: './scripts/index.ts',
      },
      output: {
        // Keep original entry file names
        entryFileNames: '[name].js',
        // Keep original chunk file names
        chunkFileNames: '[name].js',
        // Optionally keep original asset file names
        assetFileNames: '[name].[ext]',
      },
    },
  },
  fs: {
    allow: [
      // Allow serving files from one level up to the project root
      path.resolve(__dirname),
      // If you need to access more directories, add them here
      // path.resolve(__dirname, 'some-other-directory'),
    ]
  }
})

