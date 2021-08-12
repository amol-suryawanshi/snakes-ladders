package usecases

//Ladders set of ladders in the game
type Ladders struct {
	LaddersPos map[int]int
}

func NewLadders(ladders map[int]int) Ladders {
	return Ladders{LaddersPos: ladders}
}

func GetLaddersPos(l Ladders, startPos int) (int, bool) {
	v, ok := l.LaddersPos[startPos]
	return v, ok
}
