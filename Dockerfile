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
RUN npm run build

FROM registry.gitlab.com/safesurfer/go-http-server:1.2.0
ENV APP_SERVE_FOLDER=/app/dist \
    APP_TEMPLATE_MAP_PATH=/app/map.yaml \
    APP_HEADER_SET_ENABLE=true \
    APP_HEADER_MAP_PATH=/app/headers.yaml
LABEL maintainer="Caleb Woodbine <calebwoodbine.public@gmail.com>"
COPY --from=build /app/dist /app/dist
COPY template-map.yaml /app/map.yaml
COPY template-headers.yaml /app/headers.yaml
