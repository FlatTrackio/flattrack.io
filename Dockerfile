FROM node:12.10.0-alpine AS ui
WORKDIR /app
COPY . .
RUN npm i
RUN npm rebuild node-sass
RUN npm run build:frontend

FROM golang:1.13.4-alpine3.10 AS api
WORKDIR /app
COPY . .
RUN adduser -D flattrack
RUN rm -rf deployment && \
    mkdir -p deployment && \
    chown flattrack deployment
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static" -s -w' -o flattrackio src/server/server.go

FROM scratch
WORKDIR /app
ENV PATH=/app
COPY --from=ui /app/dist /app/dist
COPY --from=ui /app/package.json .
COPY --from=api /app/flattrackio .
COPY --from=api /etc/passwd /etc/passwd
COPY --chown=flattrack --from=api /app/deployment /app/deployment
EXPOSE 8080
USER flattrack
ENTRYPOINT ["/app/flattrackio"]
