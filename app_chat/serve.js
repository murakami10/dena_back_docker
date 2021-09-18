const app = require('express')();
const http = require('http').createServer(app);
const io = require('socket.io')(http);

const mysql = require('mysql2/promise');

/**
 * [イベント] ユーザーが接続
 */
io.on('connection', async (socket) => {
  console.log('ユーザーが接続しました');

  let connection;
  try {
    connection = await mysql.createConnection({
      host: 'db',
      user: 'user',
      password: 'password',
      database: 'test_database',
    });
  } catch (e) {
    console.log('connection error');
    console.log(e);
    exit;
  }

  socket.on('join', async (req) => {
    // user認証!!!!!!!!!!!!!!

    // userがroom_idに所属しているか確認

    try {
      let [results, _] = await connection.execute(
        'SELECT 1 FROM `room_members` WHERE `room_id` = ? AND `user_id` = ? ',
        [req.room, req.user_id]
      );
      if (results.length != 1) {
        // io.emit('join-response', { error: "user desen't belong to room" });
        return;
      }
    } catch (e) {
      console.log(e);
      // io.emit('join-response', { error: 'Internal error' });
    }

    let chats = [];
    let user_ids = [];

    try {
      // 過去のチャットを取得
      let [results, _] = await connection.execute(
        'SELECT `text`, `sender_id`, `created_at` FROM `chats` WHERE `room_id` = ? ',
        [req.room]
      );
      chats = Object.values(JSON.parse(JSON.stringify(results)));

      // roomに入っているuser_idを返す
      [results, _] = await connection.execute(
        'SELECT `user_id` FROM `room_members` WHERE `room_id` = ? ',
        [req.room]
      );

      let set_user_ids = new Set();
      const users = Object.values(JSON.parse(JSON.stringify(results)));
      for (let user of users) {
        set_user_ids.add(user['user_id']);
      }
      user_ids = Array.from(set_user_ids);
    } catch (e) {
      console.log(e);
      // io.emit('join-response', { error: 'Internal error' });
    }

    res = { chats, user_ids };

    // 新しくroomに入る時他のroomをすべて抜ける
    my_rooms = Array.from(socket.rooms);
    for (let my_room of my_rooms) {
      socket.leave(my_room);
    }
    socket.join(req.room);
    io.to(req.room).emit('join-response', res);
  });

  socket.on('post', (req) => {
    io.to(req.room).emit('join-response', 'aaa');
  });
});

/**
 * 3000番でサーバを起動する
 */
http.listen(3000, () => {
  console.log('listening on *:3000');
});
