FROM node:12.10.0-alpine AS ui
WORKDIR /opt/flattrack.io
COPY . .
RUN npm i
RUN npm rebuild node-sass
RUN npm run build-ui

FROM golang:1.13.4-alpine3.10 AS api
WORKDIR /opt/flattrack.io
COPY . .
RUN adduser -D flattrack
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o flattrackio src/server/server.go

FROM scratch
WORKDIR /opt/flattrack.io
ENV PATH=/opt/flattrack.io
COPY --from=ui /opt/flattrack.io/dist /opt/flattrack.io/dist
COPY --from=ui /opt/flattrack.io/package.json .
COPY --from=api /opt/flattrack.io/flattrackio .
COPY --from=api /etc/passwd /etc/passwd
EXPOSE 8080
USER flattrack
ENTRYPOINT ["/opt/flattrack.io/flattrackio"]
