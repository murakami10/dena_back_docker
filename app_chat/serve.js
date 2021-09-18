const app = require('express')();
const http = require('http').createServer(app);
const io = require('socket.io')(http);
var mysql = require('mysql');

var connection = mysql.createConnection({
  host: 'localhost',
  user: 'user',
  password: 'password',
});

connection.connect(function (err) {
  if (err) {
    console.error('error connecting: ' + err.stack);
    return;
  }

  console.log('connected as id ' + connection.threadId);
});

connection.query(
  'SELECT * FROM `books` WHERE `author` = "David"',
  function (error, results, fields) {
    console.log(results);
  }
);

/**
 * [イベント] ユーザーが接続
 */
// io.on('connection', (socket) => {
// console.log('ユーザーが接続しました');
//
// socket.on('post', (msg) => {
// io.emit('member-post', msg);
// });
//
// socket.on('list-message', (msg) => {
// io.emit('list-message-response');
// });
// });

/**
 * 3000番でサーバを起動する
 */
// http.listen(3000, () => {
// console.log('listening on *:3000');
// });
