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

export default {
  DeviceIsMobile
}
