FROM node:14.17-buster

WORKDIR /app/src

COPY ./package.json ./package-lock.json ./
RUN npm install

# とりあえず起動したらサーバ動いているように
# buildしてもいいかも
CMD ["node", "serve.js"]