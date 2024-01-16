package main

func (p *Player) CalcTurn() int {
	if len(p.shots) == 0 {
		return p.CalcGuess()
	}
	var lastGuess = p.shots[len(p.shots)-1]

	if p.uBoard[lastGuess.i] == 3 && !lastGuess.hunt {
		// Just got first hit, now guessing around
		if lastGuess.i > 9 { //up
			if p.uBoard[lastGuess.i-10] == 0 {
				return lastGuess.i - 10
			}
		}
		if lastGuess.i%10 > 0 { //left
			if p.uBoard[lastGuess.i-1] == 0 {
				return lastGuess.i - 1
			}
		}
		if lastGuess.i < 90 { //down
			if p.uBoard[lastGuess.i+10] == 0 {
				return lastGuess.i + 10
			}
		}
		return lastGuess.i + 1 //right

	}
	if (lastGuess.hunt && !lastGuess.sink) || (lastGuess.sink && p.findHit()) {
		// If you are searching but have yet to sink || you have sunk but there are still pending hits
		var g0 int = 0
		for i := 0; i < len(p.shots); i++ {
			if p.uBoard[p.shots[i].i] == 3 {
				g0 = p.shots[i].i
				break
			}
		}
		var c = []int{-10, 10, -1, 1}

		for _, ival := range c {
			for i := 0; (99 >= g0+i) && (g0+i >= 0) && (10 > i+(g0%10)) && (g0%10 >= i); i += ival {
				if p.uBoard[g0+i] == 0 {
					return g0 + i
				} else if p.uBoard[g0+i] == 3 {
					continue
				}
				break
			}
		}
	}
	// If there are no impending ship hits, calculate best guess
	return p.CalcGuess()
}

func (p Player) CalcGuess() int {
	var tBoard []int                                     // Contains total permutations for each square
	var twoBoard, threeBoard, fourBoard, fiveBoard []int // permutations for each ship size

	twoBoard = nBoard(p.uBoard, 2)
	threeBoard = nBoard(p.uBoard, 3)
	fourBoard = nBoard(p.uBoard, 4)
	fiveBoard = nBoard(p.uBoard, 5)

	tBoard = sumBoards([][]int{twoBoard, threeBoard, fourBoard, fiveBoard}, false)

	indexList := sortIndexList(tBoard)

	for _, ival := range indexList {
		if p.uBoard[ival] == 0 {
			return ival
		}
	}
	return -1
}

func nBoard(b []int, l int) []int {
	var out = make([]int, 100)
	var j = 0

	for i := 0; i < 100; i++ {
		if i%10 > 10-l {
			continue
		}
		for j = 0; j < l; j++ {
			if b[i+j] != 0 && b[i+j] != 3 {
				break
			}
		}
		if j == l {
			for k := 0; k < l; k++ {
				out[i+k]++
			}
		}
	}
	for i := 0; i < (10 * (11 - l)); i++ {
		for j = 0; j < l; j++ {
			if b[i+10*j] != 0 && b[i+10*j] != 3 {
				break
			}
		}
		if j == l {
			for k := 0; k < l; k++ {
				out[i+10*k]++
			}
		}
	}

	return out
}
