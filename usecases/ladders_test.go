package usecases

import (
	"testing"
)

func TestGetLaddersPos(t *testing.T) {
	ladder := make(map[int]int)
	ladder[2] = 15
	ladder[10] = 50
	ladders := NewLadders(ladder)
	pos, res := GetLaddersPos(ladders, 2)
	t.Run("Success", func(t *testing.T) {
		if !res {
			t.Errorf("Invalid ladder")
		}
		if pos != 15 {
			t.Errorf("Invalid ladder")
		}
	})
	pos, res = GetLaddersPos(ladders, 22)
	t.Run("Failure", func(t *testing.T) {
		if res {
			t.Errorf("Invalid ladder")
		}
		if pos != 0 {
			t.Errorf("Invalid ladder")
		}
	})
}
