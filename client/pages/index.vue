<template>
  <div class="flex flex-row w-full justify-center">
    <div class="flex flex-col items-center">
      <Card
        v-for="p in contentList"
        :title="p.title"
        :description="p.description"
        :cover="p.cover"
        :_path="p._path"
      />
    </div>
    <div class="hidden md:block md:ml-16 markdown w-96">
<!--      <h1>Hi, this is <span class="italic underline">Ray-D-Song</span>'s Blog.</h1>-->
<!--      <div class="text-xl my-5">My primary job is a web front-end development engineer (for now). <br/> But I also dabble in various other technologies.</div>-->
<!--      <div class="text-xl">You can click the icon button in the top right corner to subscribe to the RSS feed or check out my GitHub.</div>-->
    </div>
    <div class="hidden md:block markdown fixed sm:w-2/6 lg:w-2/5 top-1/4 right-20">
      <h1>Hi, this is <span class="italic underline">Ray-D-Song</span>'s Blog.</h1>
      <div class="text-xl my-5">My primary job is a web front-end development engineer (for now). <br/> But I also dabble in various other technologies.</div>
      <div class="text-xl">You can click the icon button in the top right corner to subscribe to the RSS feed or check out my GitHub.</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Card from '~/components/card.vue'
import type {ParsedContent} from '@nuxt/content/dist/runtime/types'

type Content = ParsedContent & {
  cover: string
  description: string
}

const contentList = <unknown>(await queryContent('/').find()) as Content
</script>
