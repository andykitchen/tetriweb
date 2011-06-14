package main

import (
	"flag"
	"os"
	"fmt"
	//"rand"
	"exec"
	// "syscall"
)

var terminalFlag bool

func init() {
	flag.BoolVar(&terminalFlag, "t", false, "Starts terminal game")
	flag.Parse()
}


func abs(i int) (j int) {
	if i < 0 {
		j = -i
	} else {
		j = i
	}

	return
}

func PlayTerminal() {
	//rand.Seed(1)
	g := new(Game)
	p := new(Player)
	p.id = 1
	p.name = "Test"

	g.AddPlayer(p)
	session := g.sessions[0]
	session.Start()
	ticks := 0
	//rand.Seed(1)
	exec.Run("/bin/stty", []string{"stty", "-icanon", "min", "1", "-echo"},
		nil, "", exec.PassThrough, exec.PassThrough, exec.PassThrough)
	for {

		fmt.Print(session.board.String())
		fmt.Print(ticks, "------------------\n")

		// syscall.Sleep(1000000000)
		buf := make([]byte, 1)
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println(err.String())
			return
		}

		s := string(buf)

		session.HandleKey(s)

		session.board.Tick()

		if session.board.state == FINISHED {
			fmt.Println("FINISHED")
		}
		ticks++
	}

}
