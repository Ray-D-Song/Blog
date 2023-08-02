export default defineEventHandler(async (event) => {
  try {
    return await $fetch(`api/v1/blog/hit/${event.context.params}`, {
      method: 'put'
    })
  } catch (e) {
    console.log(e)
  }
})
