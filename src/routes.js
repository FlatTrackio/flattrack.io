#!/usr/bin/node

var express = require('express')
var router = express.Router()

router.route('/')
    .get((req, res, next) => {
        res.json({message: "Welcome to flattrack.io"})
        res.status(200)
        res.end()
    })

router.route('/interested')
    .post((req, res, next) => {
        const form = req.body.form
        if (typeof form !== 'undefined' &&
            typeof form.email === 'string' &&
            /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(form.email)) {
            
            res.json({ message: 'Added to notify list sucessfully.' })
            res.status(200).end()
            return
        } else {
            res.json({ message: 'An error occured' })
            res.status(400).end()
            return
        }
    })

module.exports = router