FROM node:12.10.0-alpine AS build
LABEL maintainer="Caleb Woodbine <calebwoodbine.public@gmail.com>"
ARG AppBuildVersion="0.0.0"
ARG AppBuildHash="???"
ARG AppBuildDate="???"
ARG AppBuildMode="development"
RUN apk add tzdata
WORKDIR /app
COPY src /app/src
COPY public /app/public
COPY *.js *.json /app/
RUN npm i
RUN npm run build:frontend

FROM registry.gitlab.com/safesurfer/go-http-server:1.1.0
ENV APP_SERVE_FOLDER=/app/dist
LABEL maintainer="Caleb Woodbine <calebwoodbine.public@gmail.com>"
COPY --from=build /site/public /app/dist
COPY template-map.yaml /app/map.yaml
