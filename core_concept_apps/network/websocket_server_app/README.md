Пример простого websocket-сервера

Подключиться к серверу в js:
```javascript
// подключение к серверу
let socket = new WebSocket("ws://localhost:3000/ws")
// задать колбэк для вывода ответов от сервера
socket.onmessage = (event) => console.log("received from the server: ", event.data)
// отправка сообщения на сервер
socket.send("Hello")
```