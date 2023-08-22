<template>
  <div @click="handleRedirect" class="w-72 max-h-64 bg-curtain-light dark:bg-curtain-dark rounded-xl hover:cursor-pointer hover:opacity-80 opacity-100 mb-8 shadow shadow-lg shadow-gray dark:shadow-white/5">
    <div class="w-full h-36 flex-center rounded-xl">
      <img class="w-full h-36 rounded-t-xl object-cover" :src="$props.cover"/>
    </div>
    <div class="mt-2 mx-2">
      <div class="text-2xl mb-2 font-medium">
        {{$props.title}}
      </div>
<!--      <div class="opacity-75">-->
<!--        {{$props.description}}-->
<!--      </div>-->
    </div>
  </div>
</template>

<script setup lang="ts">
const router = useRouter()

const props = defineProps<{
  cover: string
  title: string
  description: string
  _path: string
}>()

const handleRedirect = () => {
  router.push({
    path: 'post' + props._path,
  })
}
</script>