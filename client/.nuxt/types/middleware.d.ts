import type { NavigationGuard } from 'vue-router'
export type MiddlewareKey = string
declare module "/Users/raysong/Code/CodeBase/MonoRepo/Blog/client/node_modules/.pnpm/nuxt@3.6.5_x2evywih7tvalpk7uoe7hnmwgq/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    middleware?: MiddlewareKey | NavigationGuard | Array<MiddlewareKey | NavigationGuard>
  }
}