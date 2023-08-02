<template>
  <div>
    <div>
      <h1 class="text-4xl text-center font-bold my-2">{{title}}</h1>
      <h4 class="text-center font-mono mb-5">{{time}}</h4>
    </div>
    <div id="md" v-html="result" class="markdown"/>
  </div>
</template>

<script setup lang="ts">
import MarkdownIt from 'markdown-it'
import HighLightJs from 'highlight.js'
import customFetch from '~/custom/customFetch'

const route = useRoute()
const colorMode = useColorMode()
const result = ref('')
const codeTheme = computed(() => colorMode.preference === 'dark' ? 'hljs-tokyo-night-dark' : 'hljs-github-light')

const md: MarkdownIt = new MarkdownIt({
  highlight: (str, lang, attrs) => {
    if(lang && HighLightJs.getLanguage(lang)) {
      try {
        return HighLightJs.highlight(str, {language: lang, ignoreIllegals: true}).value
      } catch (_) {}
    }
    return ''
  }
})
const getPost = useFetch('/api/getPost', {
  query: {
    id: route.query.id
  }
})
const {post, title, time} = await customFetch(getPost)
result.value = md.render(String(post))

onMounted(()=>{
  document.getElementById('md')?.classList.add(codeTheme.value)
})
watch(colorMode, ()=>{
  if(process.client){
    document.getElementById('md')?.classList.replace(colorMode.preference==='dark'?'hljs-github-light':'hljs-tokyo-night-dark', codeTheme.value)
  }
}, {deep: true})
</script>