<template>
  <div class="w-full flex justify-center items-center z-20 absolute top-10 inset-x-0">
    <div class="w-72 shadow-gray shadow-xl flex justify-center items-center rounded-xl h-12 bg-code-light dark:bg-code-dark shadow shadow-lg shadow-gray dark:shadow-white/5">
      {{ props.title }}
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string,
  subtitle?: string
}>()
</script>