/*
  common
  commonly used JS functions
*/

import news from '@/frontend/requests/public/news'

// DeviceIsMobile
// returns bool if the device is mobile (from screen size)
function DeviceIsMobile () {
  return window.innerWidth <= 870
}

// GetLatestRSSPost
// returns a struct on the latest RSS post
async function GetLatestRSSPost () {
  var feed = await news.GetFeed()
  return feed.items[0]
}

export default {
  DeviceIsMobile,
  GetLatestRSSPost
}
