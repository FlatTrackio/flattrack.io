import axios from 'axios'

function Request (request, redirect = true, publicRoute = false) {
  return new Promise((resolve, reject) => {
    axios(request)
      .then(resp => resolve(resp))
      .catch(err => {
        if (err.response.status === 401) {
          console.log(request)
        }
        reject(err)
      })
  })
}

export default Request
