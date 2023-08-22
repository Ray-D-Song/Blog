<template>
    <div class="toast">
      {{ props.title }}
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string,
  subtitle?: string
}>()
</script>