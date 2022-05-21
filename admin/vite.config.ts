import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { EsLinter, linterPlugin, TypeScriptLinter } from 'vite-plugin-linter';
import VitePluginHtmlEnv from 'vite-plugin-html-env';

// https://vitejs.dev/config/
export default defineConfig((configEnv) => ({
  plugins: [
    react(),
    VitePluginHtmlEnv(),
    linterPlugin({
      include: ['./src/**/*.ts', './src/**/*.tsx'],
      linters: [
        new EsLinter({
          configEnv: configEnv,
          serveOptions: { clearCacheOnStart: true },
        }),
        new TypeScriptLinter(),
      ],
      build: {
        includeMode: 'filesInFolder',
      },
    }),
  ],
  server: {
    port: 3002,
  },
}));
