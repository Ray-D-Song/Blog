<template>
  <div class="flex flex-col justify-center font-bold items-center w-full h-screen text-center text-2xl hover:cursor-pointer" @click="sendUHome">
    Oops! 你掉入了次元裂缝.
    <div>
      点击屏幕可以送你回家
    </div>
  </div>
</template>

<script setup lang="ts">
const sendUHome = () => {
  clearError({redirect: '/'})
}
</script>