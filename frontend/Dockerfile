FROM node:16-buster-slim

WORKDIR /app
COPY package.json package.json
RUN yarn install

COPY . /app
