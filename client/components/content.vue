<template>
  <ContentRenderer :value="data" class="markdown"/>
</template>

<script setup lang="ts">
const route = useRoute()
const {data} = await useAsyncData(String(route.params._path), () => queryContent('/' + route.params._path).findOne())
</script>