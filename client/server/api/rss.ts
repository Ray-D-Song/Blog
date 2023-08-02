import type { H3Event } from 'h3'

export default eventHandler(async (event: H3Event) => {
  try {
    return await $fetch('/api/v1/rss')
  } catch (e) {
    console.log(e)
  }
})
