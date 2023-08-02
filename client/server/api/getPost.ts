import customHandler from '~/custom/customHandler'

export default customHandler(async (event) => {
  try {
    const { id } = getQuery(event)
    const val = await $fetch('/api/v1/blog/content', {
      query: {
        id: id
      }
    })
    const res = await $fetch(val.data.list.content_url)
    return {
      data: {
        post: res,
        title: val.data.list.title,
        time: val.data.list.time
      }
    }
  } catch (e) {
    console.log(e)
  }
})