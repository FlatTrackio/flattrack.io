const axios = require('axios')

axios.get(`http://localhost:${process.env.APP_PORT}/api/`).then(resp => {
    if (resp.data.message) {
        console.log('[Health] healthy')
        process.exit(0)
    } else {
        console.log('[Health] unhealthy')
        process.exit(1)
    }
}).catch(err => {
    console.log('[Health] unhealthy')
    process.exit(1)
})
