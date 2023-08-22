<template>
  <Toast :title="title" :class="`transition-all duration-300 ${showToast?'opacity-100':'opacity-0'}`" />
  <div class="dark:bg-code-dark bg-code-light rounded-md dark:shadow-white/5 shadow-inner shadow-black/5 shadow-sm">
    <div class="opacity-60 dark:border-b-indigo-400 border-b-black/20 border-b-2 pb-2 pl-4">{{language}}</div>
    <div class="relative">
      <UButton icon="i-eva-copy-fill" @click="copyCode" size="sm" variant="ghost" :padding="false" class="icon-button absolute top-2 right-4">copy</UButton>
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import Toast from '~/components/toast.vue'
import * as clipboard from 'clipboard-polyfill'

const showToast = ref(false)
const title = ref('')

const props = defineProps({
  code: {
    type: String,
    default: ''
  },
  language: {
    type: String,
    default: null
  },
  filename: {
    type: String,
    default: null
  },
  highlights: {
    type: Array as () => number[],
    default: () => []
  },
  meta: {
    type: String,
    default: null
  }
})

const copyCode = () => {
  clipboard.writeText(props.code).then(() => {
    title.value = '已复制到剪切板'
  }).catch(() => {
    title.value = '复制失败'
  }).finally(() => {
    showToast.value = true
    setTimeout(() => {
      showToast.value = false
    }, 2000)
  })
}
</script>

<style>
pre code .line {
  display: block;
  min-height: 1rem;
}
</style>