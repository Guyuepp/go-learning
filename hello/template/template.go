package template

import "fmt"

type Game interface {
	initGame()
	playGame()
	endGame()
}

func Play(g Game) {
	g.initGame()
	g.playGame()
	g.endGame()
}

type Chess struct{}

func (c Chess) initGame() {
	fmt.Println("Initializing Chess game...")
}

func (c Chess) playGame() {
	fmt.Println("Playing Chess game...")
}

func (c Chess) endGame() {
	fmt.Println("Ending Chess game...")
}

func Test() {
	game := Chess{}
	Play(game)
}
