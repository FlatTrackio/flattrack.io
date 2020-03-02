# Building and development

> Getting started

## Dependencies
- nodejs (10+)
- npm
- golang (1.13+)

## Local

``` bash
# install dependencies
npm install
go get -v d ./...
go install ./...

# serve UI with hot reload at localhost:8080
npm run dev-ui

# build UI for production with minification
npm run build

# build UI for production and view the bundle analyzer report
npm run build --report

# run UI unit tests
npm run unit

# run UI e2e tests
npm run e2e

# run all UI tests
npm test

# run backend + frontend (requires npm run build)
npm start
```

## Local deployment and testing
Build with docker (or podman)
```bash
docker build -t registry.gitlab.com/flattrack/flattrack.io:latest .
```

Run with docker
```bash
# create volume
docker volume create flattrack.io_deployment

# bring up an instance
docker run -d -p 8080:8080 -v flattrack.io_deployment:/app/deployment:z --name flattrack.io --read-only registry.gitlab.com/flattrack/flattrack.io:latest
```
