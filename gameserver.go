package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
)
var counter int
var game *Game

func GameServer(ws *websocket.Conn) {
  if counter >= 10 {
    fmt.Println("Max number of players reached")
    return
  }

  counter += 1
  p := new(Player)
  p.id = counter
  p.name = "Test"
  
  game.AddPlayer(p)
  session := game.sessions[0]
  session.Start()
  b := session.board

	buf := make([]uint8, 512)

	fmt.Println("New player: ", p.id)

	go func() {
		for {
			n, err := ws.Read(buf)
			if err != nil {
				fmt.Println("Websocket read error: ", err.String())
				break
			}

			if(n > 0) {
        key := string(buf[:n])
				// fmt.Println("Read key: ", key)
        session.HandleKey(key)
			}
		}
	}()
	
	for {
		_, err := ws.Write(session.board.Encode())
		if err != nil {
			fmt.Println("Websocket write error: ", err.String())
			break
		}

		syscall.Sleep(100000000)
		b.Tick()
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
