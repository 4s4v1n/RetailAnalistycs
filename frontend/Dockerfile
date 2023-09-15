FROM node:16.14-alpine as node
FROM node as frontend-local-build
ARG APP_HOME=/app
WORKDIR ${APP_HOME}
COPY ./package*.json .
COPY . .
RUN npm install
COPY . ${APP_HOME}
CMD ["npm","start"]