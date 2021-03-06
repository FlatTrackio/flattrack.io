FROM node:12.10.0-alpine AS ui
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

FROM golang:1.13.4-alpine3.10 AS api
WORKDIR /app
COPY src /app/src
COPY go.* /app/
RUN adduser -D user
# TODO build the version, git hash, and build date into the binary from build-args or env
ARG AppBuildVersion="0.0.0"
ARG AppBuildHash="???"
ARG AppBuildDate="???"
ARG AppBuildMode="development"
ARG GOARCH="amd64"
RUN CGO_ENABLED=0 GOOS=linux GOARCH="$GOARCH" go build \
  -a \
  -installsuffix cgo \
  -ldflags "-extldflags '-static' -s -w \
    -X gitlab.com/flattrack/flattrack.io/src/backend/common.AppBuildVersion=$AppBuildVersion \
    -X gitlab.com/flattrack/flattrack.io/src/backend/common.AppBuildHash=$AppBuildHash \
    -X gitlab.com/flattrack/flattrack.io/src/backend/common.AppBuildDate=$AppBuildDate \
    -X gitlab.com/flattrack/flattrack.io/src/backend/common.AppBuildMode=$AppBuildMode" \
  -o flattrackio \
  src/backend/main.go

FROM scratch
WORKDIR /app
ENV PATH=/app
COPY --from=ui /app/dist /app/dist
COPY --from=ui /app/package.json .
COPY --from=ui /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=api /app/flattrackio .
COPY --from=api /etc/passwd /etc/passwd
COPY --from=api /etc/group /etc/group
COPY --chown=user migrations /app/migrations
EXPOSE 8080
USER user
ENTRYPOINT ["/app/flattrackio"]
