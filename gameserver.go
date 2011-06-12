package main

import (
	"io"
	"websocket"
	"json"
)

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
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


//func main() {
//	http.Handle("/echo", websocket.Handler(EchoServer));
//	err := http.ListenAndServe(":12345", nil);
//	if err != nil {
//		panic("ListenAndServe: " + err.String())
//	}
//}
