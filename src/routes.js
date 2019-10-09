#!/usr/bin/node

var express = require('express')
var router = express.Router()
const functions = require('./functions')
const packageJSON = require('../package.json')

router.route('/')
    .get((req, res, next) => {
        res.status(200)
        res.json({message: "Welcome to FlatTrack.io", version: packageJSON.version })
        res.end()
    })

router.route('/interested')
    .post((req, res, next) => {
        const form = req.body.form
        if (typeof form !== 'undefined' &&
            typeof form.email === 'string' &&
            form.email.length <= 70 &&
            /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(form.email)) {
            var configJSON = functions.config.read()
            if (! configJSON.emails.includes(form.email)) {
                configJSON.emails = [...configJSON.emails, form.email]
                functions.config.write(configJSON)
                res.status(200)
                res.json({ message: 'Added to notify list sucessfully.' }).end()
            } else {
                res.status(200)
                res.json({ message: 'Already subscribed' }).end()
            }
            return
        } else {
            res.status(400)
            res.json({ message: 'An error occured' }).end()
            return
        }
    })

module.exports = router