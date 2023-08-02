<template>
  <div class="flex flex-row w-full justify-center">
    <div class="flex flex-col items-center">
      <Card
        v-for="p in postList"
        :title="p.title"
        :subtitle="p.subtitle"
        :img-url="p.cover"
        :id="p.id"
      />
    </div>
    <div class="hidden md:block md:ml-16">
      This is a simple test.
      Creating structs and interfaces is an essential part of implementing dependency injection in Golang. Structs are used to define objects, while interfaces are used to define the behavior of those objects. When creating structs and interfaces for your code, it’s important to consider the dependencies that each object will have and how those dependencies will be injected.
    </div>
  </div>
</template>

<script setup lang="ts">
import Card from '~/components/card.vue'
import customFetch from '~/custom/customFetch'

type PostList = Array<{
  cover: string
  title: string
  subtitle: string
  id: number
}>

const getPostList = useFetch('/api/getPostList')

const {list: postList}: {list: PostList} = await customFetch(getPostList)
</script>
