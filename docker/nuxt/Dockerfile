FROM node:16.14.2-alpine

RUN apk update && \
    npm install -g npm && \
    npm install -g create-nuxt-app

WORKDIR /app
COPY ./ui .

ENV HOST 0.0.0.0
EXPOSE 3000
