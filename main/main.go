package main

import (
	"reverse-jam-2023/framework"
	"reverse-jam-2023/game"
	"reverse-jam-2023/loader"
)

func main() {
	res := loader.ResourceLoader{}
	ttf := res.LoadFont("default")
	gg := game.NewGame(ttf)
	g := framework.InitWindowGame(gg, 800, 600, gg.GetTitle(), ttf)

	if err := g.Run(); err != nil {
		panic(err)
	}
}
