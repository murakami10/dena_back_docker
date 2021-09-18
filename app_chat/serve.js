const app = require('express')();
const http = require('http').createServer(app);
const io = require('socket.io')(http);

/**
 * [イベント] ユーザーが接続
 */
io.on('connection', (socket) => {
  console.log('ユーザーが接続しました');

  socket.on('post', (msg) => {
    io.emit('member-post', msg);
  });

  socket.on('list-message', (msg) => {
    io.emit('list-message-response');
  });
});

/**
 * 3000番でサーバを起動する
 */
http.listen(3000, () => {
  console.log('listening on *:3000');
});
