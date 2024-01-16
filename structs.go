package main

type Player struct {
	iBoard []int // 0: Empty;    1: Hiding;  2: Missed;  3: Hit; 4: Sunk;
	uBoard []int // 0: No Guess; 1: N/A;     2: Miss;    3: Hit; 4: Sunk;

	shots []shot
	ships []Ship
}

type shot struct {
	i          int
	hunt, sink bool
}

type Game struct {
	P1 *Player
	P2 *Player

	turn int // 1: P1, 2: P2
}

type Ship struct {
	id  int
	pos int

	vertical bool
	sunk     bool
}

var idToLen = map[int]int{
	0: 2,
	1: 3,
	2: 3,
	3: 4,
	4: 5,
}
