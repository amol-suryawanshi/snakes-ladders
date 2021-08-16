package usecases

import (
	"snakes-ladders/entity"
	"testing"
)

func TestRollDice(t *testing.T) {
	dice := NewDice(entity.NormalDice)
	pawn := RollDice(dice)
	t.Run("Success", func(t *testing.T) {
		if pawn > 6 && pawn <= 0 {
			t.Errorf("Invalid roll")
		}
	})
	dice1 := NewDice(entity.CrookedDice)
	pawn = RollDice(dice1)
	t.Run("Success", func(t *testing.T) {
		if !(pawn == 6 || pawn == 2 || pawn == 4) {
			t.Errorf("Invalid roll")
		}
	})
	dice2 := NewDice(0)
	pawn = RollDice(dice2)
	t.Run("Failure", func(t *testing.T) {
		if pawn != -1 {
			t.Errorf("Invalid roll")
		}
	})

}
