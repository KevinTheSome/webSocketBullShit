import { useState , useEffect } from 'react'

function ChatBord() {
  const [message, setMessage] = useState("")
  const [messages, setMessages] = useState([] as string[])
  const socket = new WebSocket("ws://localhost:8080/chat");

  useEffect(() => {
    socket.addEventListener("message", (event) => {
      setMessages([...messages, event.data])
    });
  }, [socket])

  function sendMessage(e: any) {
    if (message === "") {
      return console.log("message is empty dumbass")
    }
    setMessage(e.target.value)
    socket.send(message);
    setMessage("")
  }

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