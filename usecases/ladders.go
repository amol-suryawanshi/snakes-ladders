package usecases

import "snakes-ladders/entity"

func NewLadders(ladders map[int]int) entity.Ladders {
	return entity.Ladders{LaddersPos: ladders}
}

func GetLaddersPos(l entity.Ladders, startPos int) (int, bool) {
	v, ok := l.LaddersPos[startPos]
	return v, ok
}
