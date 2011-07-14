package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
)

var (
	counter    int
	game       *Game
)

func GameServer(ws *websocket.Conn) {
	if counter > 10 {
		fmt.Println("Max number of players reached")
		return
	}

	p := new(Player)
	p.id = counter
	p.name = "Test"

	session := game.AddPlayer(p)
	session.Start()

	fmt.Println("New player: ", p.id)
	counter += 1

	buf := make([]byte, 512)

	go func() {
		for {
			n, err := ws.Read(buf)
			if err != nil {
				fmt.Println("Websocket read error: ", err.String())
				return // end goroutine
			}

			if n > 0 {
				key := string(buf[:n])
				session.HandleKey(key)
			}
		}
	}()

	for {
		for i := 0; i < len(game.sessions); i++ {
			_, err := ws.Write(game.sessions[i].Encode())
			if err != nil {
				fmt.Println("Websocket write error: ", err.String())
				return // end handler
			}
		}
		session.Tick()

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

// type playerHandler struct {
// 	player Player
// 	keys chan string
// 	board 
// }
// 
// func gameserver(new_players chan Player) {
// 	
// }

func main() {
	game = new(Game)
	fmt.Println("starting game...")

	if terminalFlag {
		PlayTerminal()
	} else {
		http.Handle("/tetris", websocket.Handler(GameServer))
		err := http.ListenAndServe(":4000", nil)
		if err != nil {
			panic("ListenAndServe: " + err.String())
		}
	}
}
