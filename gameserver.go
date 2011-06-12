package main

import (
	"fmt"
	"websocket"
	"json"
	"http"
	"syscall"
)

func GameServer(ws *websocket.Conn) {
	b := NewBoard()
	b.play_shape = NewShape(T_SHAPE)
	b.play_shape.RotateClockwise()
	b.sx = 15
	b.sy = 6

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
				fmt.Println("Read: ", string(buf))
			}
		}
	}()
	
	for {
		fmt.Println("Start write")

		_, err := ws.Write([]uint8(b.String()))
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
	http.Handle("/tetris", websocket.Handler(GameServer));
	err := http.ListenAndServe(":4000", nil);
	if err != nil {
		panic("ListenAndServe: " + err.String())
	}
}
