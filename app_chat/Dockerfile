FROM node:14.17-buster

WORKDIR /app

COPY package*.json ./
RUN npm install

# とりあえず起動したらサーバ動いているように
# buildしてもいいかも
CMD ["node", "serve.js"]