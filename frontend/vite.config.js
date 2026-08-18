import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api/ping": "http://localhost:8080",
      "/api": "http://localhost:8080",
      "/api/tasks": "http://localhost:8080",
    },
  },
});
