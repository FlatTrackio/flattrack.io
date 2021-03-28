/*
  interested
    post an email to the interested list
*/

import Request from '@//requests/requests'

// PostInterested
// adds an email to the interested list
function PostInterested (email) {
  return Request({
    url: '/api/interested',
    method: 'POST',
    data: {
      email
    }
  }, false, true)
}

export default {
  PostInterested
}
