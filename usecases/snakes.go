package usecases

import "snakes-ladders/entity"

func NewSnakes(snakes map[int]int) entity.Snakes {
	return entity.Snakes{SnakesPos: snakes}
}

func GetSnakesPos(s entity.Snakes, startPos int) (int, bool) {
	v, ok := s.SnakesPos[startPos]
	return v, ok
}
