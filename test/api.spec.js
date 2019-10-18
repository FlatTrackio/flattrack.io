const axios = require('axios')
const assert = require('assert')
const flattrackio = require('../src/server.js')
const functions = require('../src/functions.js')

beforeEach('initialise site\'s config file', function() {
  if (functions.config.exists()) {
    functions.config.deinit()
    functions.config.init()
  }
  flattrackio.start()
})

afterEach('deinitialise site\'s config file', function () {
  if (functions.config.exists()) {
    functions.config.deinit() 
  }
  flattrackio.stop()
})

describe('interested', function () {
  /*
    Testname: Allow standard email address posting
    Description: When posting an email address, it MUST save and appear in deployment/config.json
	*/
  it('should allow email addresses to be written', function () {
    return new Promise((resolve, reject) => {
      var user1form = {email: 'user@example.com'}
      axios.post('http://localhost:8080/api/interested', {form: user1form}).then(resp => {
        var configRead = functions.config.read()
        assert.equal(configRead.emails.includes(user1form.email), true, 'Saved config file MUST contain the just posted email')
        assert.equal(resp.status, 200, 'Server should respond with a 200')
        resolve()
      }).catch(err => {
        console.log(err)
        assert.equal(err, null, 'API should not return error')
        reject(err)
      })
    })
  })

  /*
    Testname: Disallow non email address string posting
    Description: When posting a string which is not an email address, it MUST NOT save or appear in deployment/config.json
	*/
  it('should not allow non email address strings', function () {
    return new Promise((resolve, reject) => {
      var user1form = {email: 'userexamplecom'}
      axios.post('http://localhost:8080/api/interested', {form: user1form}).then(resp => {
        console.log(resp)
        reject(reps)
      }).catch(err => {
        assert.notEqual(err, null, 'Should return error')
        var configRead = functions.config.read()
        assert.notEqual(configRead.emails.includes(user1form.email), true, 'Saved config file MUST NOT contain the just posted string')
        assert.notEqual(err.status, 400, 'API should respond with a 400')
        resolve()
      })    
    })
  })

  /*
    Testname: Disallow too long email address strings
    Description: When posting an email address, it MUST NOT save or appear in deployment/config.json if it's length is above 70
	*/
  it('should not allow email address strings with length above 70', function () {
    return new Promise((resolve, reject) => {
      var user1form = {email: 'lqDhxqzxymVTmmsZxFUaIMEqYDfOQmkhY5D8TDG6qrwZpLhAKaVU4Wbb5GTLKSMt8nE4AuHU@example.com'}
      axios.post('http://localhost:8080/api/interested', {form: user1form}).then(resp => {
        console.log(resp)
        reject(reps)
      }).catch(err => {
        assert.notEqual(err, null, 'Should return error')
        var configRead = functions.config.read()
        assert.notEqual(configRead.emails.includes(user1form.email), true, 'Saved config file MUST NOT contain the just posted string')
        assert.notEqual(err.status, 400, 'API should respond with a 400')
        resolve()
      })    
    })
  })

  /*
    Testname: Disallow empty strings
    Description: When posting, it MUST NOT allow empty strings to be sent
	*/
  it('should not allow empty strings', function () {
    return new Promise((resolve, reject) => {
      var user1form = {email: ''}
      axios.post('http://localhost:8080/api/interested', {form: user1form}).then(resp => {
        console.log(resp)
        reject(reps)
      }).catch(err => {
        assert.notEqual(err, null, 'Should return error')
        assert.notEqual(err.status, 400, 'API should respond with a 400')
        resolve()
      })    
    })
  })

  /*
    Testname: Disallow email addresses posting multiple times
    Description: When posting an email address twice, it must not allow it the second time
	*/
  it('should not allow duplicate email address strings to be saved', function () {
    return new Promise((resolve, reject) => {
      var user1form = {email: 'user@example.com'}
      axios.post('http://localhost:8080/api/interested', {form: user1form}).then(resp => {
        var configRead = functions.config.read()
        assert.equal(configRead.emails.includes(user1form.email), true, 'Saved config file MUST contain the just posted email')
        assert.equal(resp.status, 200, 'Server should respond with a 200')
      }).then(() => {
        return axios.post('http://localhost:8080/api/interested', {form: user1form})
      }).then(resp => {
        assert.equal(resp.data.message, 'Already subscribed', 'Saved config file MUST contain the just posted email')
        assert.equal(resp.status, 200, 'Server should respond with a 200')
        resolve()
      }).catch(err => {
        console.log(err)
        assert.equal(err, null, 'API should not return error')
        reject(err)
      })
    })
  })
})

describe('general', function () {
  	/*
			Testname: API availability
		  Description: When launching, the API must be available
	*/
  it('should bring up the webserver with greeting at root of API', function () {
    return new Promise((resolve, reject) => {
      axios.get('http://localhost:8080/api').then(resp => {
        assert.equal(resp.data.message, 'Welcome to FlatTrack.io', 'Calling from /api MUST return a greeting message')
        assert.equal(resp.status, 200, 'Server should respond with a 200')
        resolve()
      }).catch(err => {
        console.log(err)
        assert.equal(err, null, 'API should not return error')
        reject(err)
      })
    })
  })
})
