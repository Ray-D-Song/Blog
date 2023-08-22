<template>
  <div class="flex min-h-screen items-start flex-col justify-between">
    <div class="sticky top-0 flex items-center mb-8 w-full backdrop-blur-sm z-10 bg-gray-300 dark:bg-gray-900 justify-between h-16 shadow-sm">
      <div class="flex items-center" @click="handleRedirect(1)">
        <Icon
          name="bi:lightning-fill"
          class="w-9 h-9 hover:text-amber-400 mx-3 mt-3 cursor-pointer"/>
      </div>
      <div class="flex items-center">
        <UButton
          :padded="false"
          icon="i-simple-icons-github"
          size="sm"
          variant="ghost"
          class="icon-button"
          @click="handleRedirect(2)"
        />
        <div class="w-1"/>
        <UButton
          :padded="false"
          icon="i-bi-rss"
          size="sm"
          variant="ghost"
          class="icon-button"
          @click="handleRedirect(3)"
        />
        <UButton
          @click="toggleColorMode"
          :padded="false"
          icon="i-fluent-dark-theme-24-filled"
          size="xl"
          variant="ghost"
          class="icon-button"
        />
      </div>
    </div>
    <Toast title="RSS 订阅已复制到剪切板" class="transition-all duration-300" :class="showToast?'opacity-100':'opacity-0'" />
    <div class="flex-center w-full flex-grow">
      <div class="w-4/5 sm:w-7/10 lg:w-3/5 xl:w-3/5 2xl:w-2/5">
        <slot />
      </div>
    </div>
    <div class="mt-12 pb-6 text-sm text-center opacity-50 w-full self-end">
      <div>Built with Nuxt 3</div>
      Copyright 2023 © Ray Song
    </div>
  </div>
</template>

<script setup lang="ts">
import Toast from '~/components/toast.vue'
import * as clipboard from 'clipboard-polyfill'

const showToast = ref(false)
const colorMode = useColorMode()
const router = useRouter()
const routeTable = {
  1: '/',
  2: 'https://github.com/Ray-D-Song',
  3: 'http://ray-d-song/rss.xml'
}

const toggleColorMode = () => {
  if(colorMode.preference === 'light')
    colorMode.preference = 'dark'
  else colorMode.preference = 'light'
}

const handleRedirect = (key: number) => {
  switch (key) {
    case 1:
      router.push(routeTable[1])
      break
    case 2:
      window.location.replace(routeTable[2])
      break
    case 3:
      {
        clipboard.writeText(routeTable[3]).then(() => {
          showToast.value = true
          setTimeout(() => {
            showToast.value = false
          }, 2000)
        })
      }
      break
  }
}
</script>