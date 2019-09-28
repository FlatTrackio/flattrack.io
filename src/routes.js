#!/usr/bin/node

var express = require('express')
var router = express.Router()

router.route('/')
    .get((req, res, next) => {
        res.json({message: "Welcome to vuejs-express-buefy-template."})
        res.end()
    })

module.exports = router