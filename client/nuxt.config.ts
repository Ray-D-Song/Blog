// https://nuxt.com/docs/api/configuration/nuxt-config
import {queryContent} from '@nuxt/content/dist/runtime/composables/query'

const pathList = (await queryContent('/').only('_path').find()).map(item => ('post' + item))

export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxthq/ui',
    '@nuxtjs/color-mode',
    'nuxt-icon',
    '@nuxt/content',
  ],
  ssr: true,
  ui: {
    icons: ['eva', 'fa6-brands', 'simple-icons', 'bi', 'fluent']
  },
  nitro: {
    routeRules: {
      '/api/v1/**': {
        proxy: 'http://127.0.0.1:9000/api/v1/**',
      }
    },
    prerender: {
      routes: ['/sitemap.xml', '/rss.xml', ...pathList]
    },
    static: true
  },
  content: {
    highlight: {
      theme: {
        default: 'github-dark',
        dark: 'github-dark',
        light: 'github-light'
      },
      preload: [
        'javascript',
        'typescript',
        'vue',
        'vue-html',
        'java',
        'ruby',
        'go',
        'jsx',
        'tsx',
        'bash',
        'dockerfile',
        'c'
      ]
    }
  }
})
