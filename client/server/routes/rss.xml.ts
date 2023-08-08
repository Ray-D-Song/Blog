import { serverQueryContent } from '#content/server'
// @ts-ignore
import RSS from 'rss'

export default defineEventHandler(async (event) => {
  const docs = await serverQueryContent(event).find()
  const domain = 'https://ray-d-song.com'
  const feed = new RSS({
    title: "Ray-D-Song's Blog",
    description: '只是作为码农发牢骚的地方罢了',
    site_url: domain,
    feed_url: domain + '/rss.xml'
  })

  for (const doc of docs) {
    feed.item({
      title: doc.title as string,
      url: domain + `/post${doc._path}`,
      guid: domain + `/post${doc._path}`,
      date: Number(doc._path?.substr(1)),
      description: doc.description
    })
  }

  const feedString = feed.xml({ indent: true })
  event.res.setHeader('content-type', 'text/xml')
  event.res.end(feedString)
})
