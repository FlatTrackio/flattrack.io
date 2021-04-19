FROM klakegg/hugo:0.81.0-ext-alpine-ci AS build
ARG HUGO_ENV=production
WORKDIR /site
COPY . .
RUN hugo

FROM registry.gitlab.com/safesurfer/go-http-server:1.2.0
ENV APP_SERVE_FOLDER=/app/dist \
    APP_TEMPLATE_MAP_PATH=/app/map.yaml \
    APP_HEADER_SET_ENABLE=true \
    APP_HEADER_MAP_PATH=/app/headers.yaml
LABEL maintainer="Caleb Woodbine <calebwoodbine.public@gmail.com>"
COPY --from=build /site/public /app/dist
COPY template-map.yaml /app/map.yaml
COPY template-headers.yaml /app/headers.yaml
