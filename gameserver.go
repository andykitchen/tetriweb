package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
	"sync"
)

var (
	counter int
	game *Game
	game_mutex sync.Mutex
)

func GameServer(ws *websocket.Conn) {
  if counter > 10 {
    fmt.Println("Max number of players reached")
    return
  }

  p := new(Player)
  p.id = counter
  p.name = "Test"

	// TODO restructure code to remove mutex, use channels instead
	game_mutex.Lock()
	session := game.AddPlayer(p)
  session.Start()
	game_mutex.Unlock()

	fmt.Println("New player: ", p.id)
  counter += 1

	buf := make([]uint8, 512)

	go func() {
		for {
			n, err := ws.Read(buf)
			if err != nil {
				fmt.Println("Websocket read error: ", err.String())
				return // end goroutine
			}

			if(n > 0) {
        key := string(buf[:n])
				game_mutex.Lock()
        session.HandleKey(key)
				game_mutex.Unlock()
			}
		}
	}()
	
	for {
		game_mutex.Lock()
    for i := 0; i < len(game.sessions); i++ {
      _, err := ws.Write(game.sessions[i].Encode())
      if err != nil {
        fmt.Println("Websocket write error: ", err.String())
        return // end handler
      }
    }
		session.board.Tick()
		game_mutex.Unlock()

		syscall.Sleep(100000000)
	}
}

type PlayerBoard struct {
	CurrentState []int
	PlayShape    *Shape
	NextShape    *Shape
}


func (b *Board) ToJson() (result []byte) {
	player_board := new(PlayerBoard)
	player_board.PlayShape = b.play_shape
	player_board.NextShape = b.next_shape
	player_board.CurrentState = b.GetBoardState()

	result, _ = json.Marshal(player_board)
	return
}


func main() {
  game = new(Game)
  fmt.Println("starting game...")

  if terminalFlag {
    PlayTerminal()
  } else {
    http.Handle("/tetris", websocket.Handler(GameServer));
    err := http.ListenAndServe(":4000", nil);
    if err != nil {
      panic("ListenAndServe: " + err.String())
    }
  }
}
