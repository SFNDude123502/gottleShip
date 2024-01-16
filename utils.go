package main

import (
	"fmt"
	"math/rand"
)

func sumBoards(b [][]int, double3 bool) []int {
	var out = make([]int, 100)
	for i := 0; i < 100; i++ {
		out[i] = b[0][i] + b[1][i] + b[2][i] + b[3][i]
		if double3 {
			out[i] += b[1][i]
		}
	}
	return out
}

func printBoard(b []int) {
	for i := range b {
		fmt.Print(b[i], " ")
		if i%10 == 9 {
			fmt.Println()
		}
	}
}

func findLargest(b []int) int {
	var biggest = 0
	for i := 0; i < 100; i++ {
		if b[i] > b[biggest] {
			biggest = i
		}
	}
	return biggest
}

func (g *Game) sinkDetector(player int) bool {
	var p *Player
	var o *Player
	if player == 1 {
		p = g.P1
		o = g.P2
	} else {
		p = g.P2
		o = g.P1
	}
	var j = 0
	var locs []int
	for i, ival := range p.ships {
		locs = ival.getLocs()
		for j = 0; j < len(locs); j++ {
			if p.iBoard[locs[j]] != 3 {
				break
			}
		}
		if j == len(locs) && !ival.sunk {
			p.ships[i].sunk = true
			for _, ii := range locs {
				p.iBoard[ii] = 4
			}
			for _, ii := range locs {
				o.uBoard[ii] = 4
			}
			return true
		}

	}

	return false
}

func (g *Game) SummonShips(player int) {
	var randBool bool
	var randInt int
	var err string

	for i := 4; i >= 0; i-- {
		randBool = (rand.Intn(2) == 1)
		err = "-1"
		for err != "" {
			randInt = rand.Intn(100)
			err = g.SetShip(randInt, player, i, randBool)
		}
	}

}

func (ship *Ship) getLocs() []int {
	var locs = make([]int, 0)
	var scalar = 1
	if ship.vertical {
		scalar = 10
	}

	for i := 0; i < idToLen[ship.id]; i++ {
		locs = append(locs, ship.pos+i*scalar)
	}

	return locs
}

func sortIndexList(arr []int) []int {
	var out = make([]int, len(arr))
	var lrgI int
	for i := range arr {
		lrgI = 0
		for j, jval := range arr {
			if jval > arr[lrgI] {
				lrgI = j
			}
		}
		out[i] = lrgI
		arr[lrgI] = 0
	}
	return out
}

func (p *Player) findHit() bool {
	for _, ival := range p.uBoard {
		if ival == 3 {
			return true
		}
	}
	return false
}
