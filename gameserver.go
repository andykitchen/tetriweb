package main

import (
	"fmt"
	"net/http"
	"time"
	"code.google.com/p/go.net/websocket"
	// "syscall"
)

type Conn struct {
	*websocket.Conn
	done chan int
}

func (c *Conn) Close() {
	c.done <- 1
}

func newWebSocketHandler(connections chan Conn) websocket.Handler {
	return func(ws *websocket.Conn) {
		done := make(chan int)
		connections <- Conn{ws, done}
		<-done
	}
}

type Broadcaster struct {
	subscriptions chan chan GameSession
	writes        []chan GameSession
	send          chan GameSession
}

func newBroadcaster() (b *Broadcaster) {
	b = new(Broadcaster)
	b.subscriptions = make(chan chan GameSession)
	b.writes = make([]chan GameSession, 0)
	b.send = make(chan GameSession)

	go func() {
		for {
			select {
			case sub := <-b.subscriptions:
				b.writes = append(b.writes, sub)
			case data := <-b.send:
				for _, x := range b.writes {
					x <- data
				}
			}
		}
	}()

	return
}

func (b *Broadcaster) Subscribe(sub chan GameSession) {
	b.subscriptions <- sub
}

func gameserver(connections chan Conn) {
	ticker := time.NewTicker(1e8)
	sessions := make([]GameSession, 0)
	broadcast := newBroadcaster()
	player_count := 0

	for {
		select {
		case conn := <-connections:
			session := GameSession{&Player{player_count, "foo"}, &Board{}}
			session.Start() // required initialization
			sessions = append(sessions, session)
			player_count++

			go func() {
				buf := make([]byte, 512)

				for {
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Websocket read error: ", err.String())
						return // end goroutine
					}

					if n > 0 {
						key := string(buf[:n])
						fmt.Print(key)
						session.HandleKey(key)
					}
				}
			}()

			session_chan := make(chan GameSession)
			broadcast.Subscribe(session_chan)

			go func() {
				for {
					session := <-session_chan
					_, err := conn.Write(session.Encode())
					if err != nil {
						fmt.Println("Websocket write error: ", err.String())
						return // end goroutine
					}
				}
			}()

		case <-ticker.C:
			for _, session := range sessions {
				session.Tick()
				broadcast.send <- session
			}
		}
	}

	ticker.Stop()
}

func main() {
	connections := make(chan Conn)

	go gameserver(connections)

	http.Handle("/tetris", newWebSocketHandler(connections))
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.String())
	}
}
