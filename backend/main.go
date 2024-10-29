package main

import (
	"flag"
	"log"
	"net/http"

	"database/sql"

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

func test(db *sql.DB) {
	AddMessage(db, "test message", "test user")
	log.Println(GetMessages(db))

}

func main() {
	db := OpenDB()
	log.Println("server started")
	defer CloseDB(db)

	test(db)
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
