import tailwindcss from '@tailwindcss/vite'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ['~/assets/css/tailwind.css'],
  modules: [
    '@vueuse/nuxt',
    'shadcn-nuxt'
  ],
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE
    },
    apiBase: process.env.API_BASE
  },
  shadcn: {
    /**
     * Prefix for all the imported component
     */
    prefix: '',
    /**
     * Directory that the component lives in.
     * @default "./components/ui/cn"
     */
    componentDir: './app/components/ui'
  },
  vite: {
    plugins: [
      tailwindcss(),
    ],
    server: {
      watch: {
        usePolling: true,
        interval: 100
      }
    }
  },

})