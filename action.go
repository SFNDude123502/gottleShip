package main

import "fmt"

func initGame() *Game {
	var game *Game = &Game{}
	game.turn = 1
	game.P1 = &Player{make([]int, 100), make([]int, 100), []shot{}, []Ship{}}
	game.P2 = &Player{make([]int, 100), make([]int, 100), []shot{}, []Ship{}}
	return game
}

func (game *Game) HumanMove(pos int) []int {
	var out = make([]int, 2) // 0: miss, 1: hit, 2: sink, 3: win, 4: error

	var result = game.makeMove(pos)
	switch result {
	case 0:
		out[0] = 0
		break
	case 1:
		out[0] = 1
		break
	case 3:
		out[0] = 4
		break
	}
	if game.P1.shots[len(game.P1.shots)-1].sink {
		out[0] = 2
	}
	if result == 2 {
		out[0] = 3
	}

	result = game.makeMove(game.P2.CalcTurn())
	switch result {
	case 0:
		out[1] = 0
		break
	case 1:
		out[1] = 1
		break
	case 3:
		out[1] = 4
		break
	}
	if game.P2.shots[len(game.P2.shots)-1].sink {
		out[1] = 2
	}
	if result == 2 {
		out[1] = 3
	}

	return out
}

func (game *Game) SetShip(pos, player, id int, verticle bool) string {
	if len(game.P1.shots) != 0 || len(game.P2.shots) != 0 {
		return "Cannot place ships once game has started"
	}
	var p *Player
	if player == 1 {
		p = game.P1
	} else {
		p = game.P2
	}

	for _, ival := range p.ships {
		if ival.id == id {
			return "Ship already exists"
		}
	}

	var scalar int
	if verticle {
		if pos+10*(idToLen[id]-1) > 99 {
			return "Ship placed over edge of map"
		}
		scalar = 10
	} else {
		if pos%10+(idToLen[id]-1) > 9 {
			return "Ship cannot placed over edge of map"
		}
		scalar = 1
	}
	for i := 0; i < idToLen[id]; i++ {
		if p.iBoard[pos+scalar*i] == 1 {
			return "Ships cannot overlap"
		}
	}
	for i := 0; i < idToLen[id]; i++ {
		p.iBoard[pos+scalar*i] = 1
	}

	p.ships = append(p.ships, Ship{id, pos, verticle, false})
	return ""
}

func (game *Game) makeMove(pos int) int { // 0: Miss, 1: Hit, 2: Win, 3: Error
	var p, o *Player
	if game.turn == 1 {
		p = game.P1
		o = game.P2
		game.turn++
	} else {
		p = game.P2
		o = game.P1
		game.turn--
	}

	var hit = o.iBoard[pos]

	if hit > 1 {
		return 3 // shooting a previously shot space
	}

	o.iBoard[pos] += 2            // marking the square as shot on opponent's board
	p.uBoard[pos] = o.iBoard[pos] // marking the shot on player's board

	var move = shot{} //Recording the recently taken shot
	move.i = pos

	if len(p.shots) == 0 {
		p.shots = []shot{move}
		return hit
	}

	var lastShot = p.shots[len(p.shots)-1]
	move.hunt = (p.uBoard[lastShot.i] == 3 && !lastShot.hunt) || (lastShot.hunt && !lastShot.sink) || (p.findHit() && lastShot.sink)
	move.sink = game.sinkDetector(game.turn)

	fmt.Println(move)
	p.shots = append(p.shots, move)

	for _, ival := range o.ships {
		if !ival.sunk {
			return hit
		}
	}

	return 2
}
