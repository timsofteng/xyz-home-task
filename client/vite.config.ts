/// <reference types="vitest" />
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react-swc";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  const port = +(process.env.VITE_SERVER_PORT || 3001);

  return {
    server: { port },
    plugins: [react()],
    test: {
      environment: "jsdom",
    },
    resolve: {
      alias: {
        "@assets": resolve(__dirname, "./src/assets"),
        "@components": resolve(__dirname, "./src/components"),
        "@layouts": resolve(__dirname, "./src/layouts"),
        "@lib": resolve(__dirname, "./src/lib"),
        "@pages": resolve(__dirname, "./src/pages"),
        "@store": resolve(__dirname, "./src/store"),
        "@styles": resolve(__dirname, "./src/styles"),
        "@config": resolve(__dirname, "./src/config.ts"),
        "@main": resolve(__dirname, "./src/main.tsx"),
      },
    },
  };
});
