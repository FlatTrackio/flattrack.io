FROM node:12.10.0-alpine
RUN mkdir -p /opt/flattrack.io
ADD . /opt/flattrack.io/.
RUN chown -R node:node /opt/flattrack.io
WORKDIR /opt/flattrack.io
RUN npm i
ENV NODE_ENV=production
ENV APP_PORT=80
ENV NO_UPDATE_NOTIFIER=true
USER node
CMD npm start
