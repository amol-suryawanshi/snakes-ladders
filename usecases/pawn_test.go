package usecases

import (
	"testing"
)

func TestGetPos(t *testing.T) {
	pawn := NewPawn()
	SetPos(&pawn, 10)
	t.Run("Success", func(t *testing.T) {
		if pawn.CurrentPos != 10 {
			t.Errorf("Invalid ladder")
		}
	})
	pos := GetPos(&pawn)
	t.Run("Success", func(t *testing.T) {
		if pawn.CurrentPos != pos {
			t.Errorf("Invalid ladder")
		}
	})
}
