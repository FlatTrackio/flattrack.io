/*
  news
    fetch the latest blog posts
*/

import Request from '@/frontend/requests/requests'

// GetLatestRSSPost
// returns the latest post
function GetLatestRSSPost () {
  return Request({
    url: '/api/latestPost',
    method: 'GET'
  }, false, true)
}

export default {
  GetLatestRSSPost
}
