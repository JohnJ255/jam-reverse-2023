package main

import (
	"reverse-jam-2023/framework"
	"reverse-jam-2023/game"
)

func main() {
	gg := game.NewGame()
	g := framework.InitWindowGame(gg, 800, 600, gg.GetTitle())

	if err := g.Run(); err != nil {
		panic(err)
	}
}
