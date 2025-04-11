import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  define: {
    "process.env": process.env,
  },
  server: {
    port: 3000,
  },
  resolve: {
    alias: {
      "@assets": "/src/assets",
      "@components": "/src/components",
      "@constants": "/src/constants",
      "@interfaces": "/src/interfaces",
      "@layouts": "/src/layouts",
      "@libs": "/src/libs",
      "@pages": "/src/pages",
      "@redux": "/src/redux",
      "@utils": "/src/utils",
    },
  },
  optimizeDeps: {
    exclude: ["js-big-decimal"],
  },
  css: {
    preprocessorOptions: {
      scss: {
        api: "modern",
        quietDeps: true,
      },
    },
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ["react", "react-dom"],
        },
      },
    },
  },
});
