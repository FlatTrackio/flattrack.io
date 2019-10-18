FROM node:12.10.0-alpine
RUN mkdir -p /opt/flattrack.io
ADD . /opt/flattrack.io/.
RUN chown -R node:node /opt/flattrack.io
USER node
WORKDIR /opt/flattrack.io
RUN npm i
RUN npm run build
ENV PUID=1000 \
    PGID=1000 \
    NODE_ENV=production \
    APP_PORT=8080 \
    NO_UPDATE_NOTIFIER=true
HEALTHCHECK CMD node src/healthcheck.js
EXPOSE 8080
CMD npm start
