import type { H3Event } from 'h3'

const customHandler = (handler: (event: H3Event) => any) => {
  const _handler = async (event: H3Event) => {
    const response = await handler(event)

    const responseNew = response.data

    return responseNew
  }

  return defineEventHandler(_handler)
}

export default customHandler