import { ComputedRef, Ref } from 'vue'
export type LayoutKey = "default"
declare module "/Users/raysong/Code/CodeBase/MonoRepo/Blog/client/node_modules/.pnpm/nuxt@3.6.5_x2evywih7tvalpk7uoe7hnmwgq/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    layout?: false | LayoutKey | Ref<LayoutKey> | ComputedRef<LayoutKey>
  }
}