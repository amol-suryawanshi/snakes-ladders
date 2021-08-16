package usecases

import (
	"snakes-ladders/entity"
	"testing"
)

func TestNewBoard(t *testing.T) {
	dice := entity.Dice{
		Type: entity.NormalDice,
	}
	pawn := entity.Pawn{
		CurrentPos: 0,
	}
	snakes := entity.Snakes{}
	ladders := entity.Ladders{}
	board := NewBoard(dice, pawn, snakes, ladders)
	t.Run("Success", func(t *testing.T) {
		if board.Dice.Type != entity.NormalDice {
			t.Errorf("Invalid Board")
		}
	})
	t.Run("Success", func(t *testing.T) {
		if board.Pawn.CurrentPos != 0 {
			t.Errorf("Invalid Board")
		}
	})
	t.Run("Failure", func(t *testing.T) {
		if board.Dice.Type == entity.CrookedDice {
			t.Errorf("Invalid Board")
		}
	})
}
