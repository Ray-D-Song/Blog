import customHandler from '~/custom/customHandler'

export default customHandler(async () => {
  try {
    return await $fetch('/api/v1/blog/list')
  } catch (e) {
    console.log(e)
  }
})
