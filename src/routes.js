#!/usr/bin/node

var express = require('express')
var router = express.Router()
const functions = require('./functions')

router.route('/')
    .get((req, res, next) => {
        res.json({message: "Welcome to FlatTrack.io"})
        res.status(200)
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
                res.json({ message: 'Added to notify list sucessfully.', counter: configJSON.emails.length })
                res.status(200).end()
            } else {
                res.json({ message: 'Already subscribed' })
                res.status(200).end()
            }
            return
        } else {
            res.json({ message: 'An error occured' })
            res.status(400).end()
            return
        }
    })

module.exports = router