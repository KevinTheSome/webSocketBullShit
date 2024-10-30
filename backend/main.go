package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	// Allow cross-domain connections
	CheckOrigin: func(r *http.Request) bool { return true },
} // use default options

func echo(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	db := OpenDB()
	defer CloseDB(db)

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		AddMessage(db, string(message), "test user")

		messageJson, err := json.Marshal(GetMessages(db))
		if err != nil {
			log.Fatal(err)
		}
		err = c.WriteMessage(mt, messageJson)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.Println("server started")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/chat", chat)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
