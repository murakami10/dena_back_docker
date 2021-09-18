<script
  src="https://cdn.socket.io/4.1.2/socket.io.min.js"
  integrity="sha384-toS6mmwu70G0fw54EGlWWeA4z3dyJ+dlXBtSURSKN4vyRFOcxd3Bzjj/AoOwY+Rg"
  crossorigin="anonymous"
></script>;

const socket = io('http://localhost:3000', { transports: ['websocket'] });

// １つ目のブラウザ
socket.emit('join', { user_id: 1, room: 1 });
socket.on('join-response', (msg) => {
  console.log(msg);
});

// 2つ目のブラウザ
socket.emit('join', { user_id: 2, room: 1 });
socket.on('join-response', (msg) => {
  console.log(msg);
});

// 3つ目のブラウザ
socket.emit('join', { user_id: 3, room: 2 });
socket.on('join-response', (msg) => {
  console.log(msg);
});
