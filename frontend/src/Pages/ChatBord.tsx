import { useState } from 'react'

function ChatBord() {
  const [message, setMessage] = useState("")
  const [messages, setMessages] = useState([] as string[])

  function sendMessage(e: any) {
    setMessage(e.target.value)
    socket.send(message);
    setMessage("")
  }

  const socket = new WebSocket("ws://localhost:8080/chat");

  // Listen for messages
  socket.addEventListener("message", (event) => {
    console.log(event)
    console.log(JSON.parse(event.data))
    setMessages(JSON.parse(event.data))
  });

  const messagesJsx = messages.map((message, index) => <li key={index}>{message}</li>)

  return (
    <>
    <h1>Chat WebSocket server</h1>
    <ul>{messagesJsx}</ul>
    <input type="text" value={message} onChange={(e) => setMessage(e.target.value)} />
    <input type="button" value="send" onClick={sendMessage} />
    </>
  )
}

export default ChatBord