const fs = require('fs')
const path = require('path')
const configPath = path.resolve(path.join('.', 'deployment', 'config.json'))

const configJSONTemplate = {
  emails: []
}

function doesExistConfigJSON () {
  return fs.existsSync(configPath)
}

function initConfigJSON () {
  if (! doesExistConfigJSON()) {
    if (fs.mkdirSync(path.resolve(path.join('.', 'deployment')), { recursive: true })) {
      return writeConfigJSON(configJSONTemplate)
    } else if (! doesExistConfigJSON()) {
      return writeConfigJSON(configJSONTemplate)
    } else  return false
  } else return true
}

function deinitConfigJSON () {
  if (fs.unlinkSync(configPath)) {
    return true
  } else {
    return false
  }
}

function readConfigJSON () {
  return require(configPath)
}

function writeConfigJSON (content) {
  return fs.writeFileSync(configPath, JSON.stringify(content, null, 2))
}

module.exports = {
  config: {
    exists: doesExistConfigJSON,
    init: initConfigJSON,
    deinit: deinitConfigJSON,
    read: readConfigJSON,
    write: writeConfigJSON,
    defaultMap: configJSONTemplate
  }
}