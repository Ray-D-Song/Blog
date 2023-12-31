// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxthq/ui',
    '@nuxtjs/i18n',
    '@nuxtjs/color-mode',
    'nuxt-icon',
    '@nuxt/content'
  ],
  ui: {
    icons: ['eva', 'fa6-brands', 'simple-icons', 'bi', 'fluent']
  },
  nitro: {
    routeRules: {
      '/api/v1/**': {
        proxy: 'http://127.0.0.1:9000/api/v1/**',
      }
    }
  }
})
