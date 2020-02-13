package app

import (
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

func setupSocketIOServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Error(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		s.Emit("hello", "yay")
		return nil
	})
	server.OnEvent("/", "hello", func(s socketio.Conn, msg string) {
		log.Println("I got! ", msg)
		s.Emit("hello", "das ca!")
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
		log.Println("Say... ", msg)
		s.Emit("msg", time.Now())
	})

	server.OnError("/", func(e error) {
		log.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
