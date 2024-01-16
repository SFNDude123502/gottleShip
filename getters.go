package main

func (g *Game) GetIBoard() []int {
	return g.P1.iBoard
}

func (g *Game) GetUBoard() []int {
	return g.P1.uBoard
}
