/*
  news
    fetch the latest blog posts
*/

import RSSParser from 'rss-parser'

async function GetFeed () {
  var parser = new RSSParser()
  var feed = await parser.parseURL('https://blog.flattrack.io/rss')
  return feed
}

export default {
  GetFeed
}
