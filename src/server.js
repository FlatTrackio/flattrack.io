#!/usr/bin/node

const express = require('express')
const bodyParser = require('body-parser')
const app = express()
const path = require('path')
const morgan = require('morgan')
const functions = require('./functions')

// development port is 8080
var port = process.env.APP_PORT || 8080
var serverObject

functions.config.init()

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({
    extended: true
}))

app.set('trust proxy', true)

app.use(express.json())
app.use(morgan('combined'))

var routes = require('./routes')
app.use('/api', routes)

// Sends static files  from the public path directory
app.use(express.static(path.join(__dirname, '..', 'dist')))

app.get('/#', (req, res) => {
    res.redirect('/')
})

app.get(/(.*)/, (req, res) => {
    res.redirect('/#/unknown-page')
})

function start () {
    return app.listen(port, () => {
        console.log(`Running on port ${port}`)
    })
}

if (require.main !== module) {
    module.exports = {
        start: () => {
            serverObject = start()
        },
        stop: () => {
            serverObject.close()
        }
    }
} else start()