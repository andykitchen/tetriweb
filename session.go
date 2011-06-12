package main

import (
  "fmt"
)

type Player struct {
  id int
  name string
}

type Game struct {
  sessions []GameSession
}

type GameSession struct {
  player *Player
  board *Board
}


func (g *Game) AddPlayer(player *Player) {
  s := new(GameSession)
  s.board = new(Board)
  s.player = player
  
  g.sessions = append(g.sessions, *s)
}

func (s *GameSession) HandleKey(key string) {
  switch key {
  case "h":
    fmt.Println(s.player.id, " LEFT!")
    s.board.MoveLeft()
  case "l":
    fmt.Println(s.player.id, " RIGHT!")
    s.board.MoveRight()
  case "j":
    fmt.Println(s.player.id, " DOWN!")
    s.board.MoveDown()
  case "u":
    fmt.Println(s.player.id, " A ROTATE!")
    s.board.RotateCounterClockwise()
  case "o":
    fmt.Println(s.player.id, " C ROTATE!")
    s.board.RotateClockwise()
  case " ":
    fmt.Println(s.player.id, " DROP!")
    s.board.DropShape()
  }
}

func (s *GameSession) Start() {
  s.board.AddShape()
}
