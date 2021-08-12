package usecases

import "snakes-ladders/entity"

func NewBoard(dice entity.Dice, pawn entity.Pawn, snakes entity.Snakes, ladders entity.Ladders) entity.Board {
	return entity.Board{
		Dice:    dice,
		Pawn:    pawn,
		Snakes:  snakes,
		Ladders: ladders,
	}
}
