package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
)

func GameServer(ws *websocket.Conn) {
	var i int
	b := NewBoard()
	b.play_shape = NewShape(T_SHAPE)
	b.play_shape.RotateClockwise()
	b.sx = 15
	b.sy = 6

	if err != nil {
		fmt.Println("Couldn't set timeout: ", err.String())
		return
	}

	// rand.Seed(1)
	
	buf := make([]uint8, 512)
	for {
		fmt.Println("loop")
		// s := fmt.Sprintf("Hello from server: %d!!!", i)
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println("Websocket read error: ", err.String())
		}
		
		if(n > 0) {
			fmt.Println("Read: ", string(buf))
		}
		
		ws.Write([]uint8(b.String()))
		syscall.Sleep(500000000)
		i++
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
	http.Handle("/echo", websocket.Handler(GameServer));
	err := http.ListenAndServe(":4000", nil);
	if err != nil {
		panic("ListenAndServe: " + err.String())
	}
}
