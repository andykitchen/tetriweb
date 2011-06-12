package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
)
var counter int

func GameServer(ws *websocket.Conn) {
  counter += 1
  g := new(Game)
  p := new(Player)
  p.id = counter
  p.name = "Test"
  
  g.AddPlayer(p)
  session := g.sessions[0]
  session.Start()
  b := session.board

	// rand.Seed(1)
	
	buf := make([]uint8, 512)
	
	go func() {
		for {
			fmt.Println("Start read")

			n, err := ws.Read(buf)
			if err != nil {
				fmt.Println("Websocket read error: ", err.String())
				break
			}

			if(n > 0) {
        key := string(buf[:n])
				fmt.Println("Read key: ", key)
        session.HandleKey(key)
			}
		}
	}()
	
	for {
		fmt.Println("Start write")
    //session.HandleKey("l")
		_, err := ws.Write([]uint8(session.board.String()))
		if err != nil {
			fmt.Println("Websocket write error: ", err.String())
			break
		}

		syscall.Sleep(500000000)
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
