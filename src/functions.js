const fs = require('fs')
const path = require('path')

const configJSONTemplate = {
  emails: []
}

function initConfigJSON () {
  if (! fs.existsSync(path.resolve(path.join('.', 'deployment', 'config.json')))) {
    fs.mkdir(path.resolve(path.join('.', 'deployment')), { recursive: true }, (err) => {
      console.log(err)
      return false
    })
    return fs.writeFileSync(path.resolve(path.join('.', 'deployment', 'config.json')), JSON.stringify(configJSONTemplate, null, 2))
  } else return true
}

function readConfigJSON () {
  return require(path.resolve(path.join('.', 'deployment', 'config.json')))
}

function writeConfigJSON (content) {
  return fs.writeFileSync(path.resolve(path.join('.', 'deployment', 'config.json')), JSON.stringify(content, null, 2))
}

module.exports = {
  config: {
    init: initConfigJSON,
    read: readConfigJSON,
    write: writeConfigJSON
  }
}