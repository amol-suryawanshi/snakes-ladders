package entity

//Board of the Game
type Board struct {
	Dice    Dice
	Pawn    Pawn
	Snakes  Snakes
	Ladders Ladders
}

//Dice for game
type Dice struct {
	Type int
}

//Ladders set of ladders in the game
type Ladders struct {
	LaddersPos map[int]int
}

//Pawn of player
type Pawn struct {
	CurrentPos int8
}

//Snakes set of snakes
type Snakes struct {
	SnakesPos map[int]int
}
