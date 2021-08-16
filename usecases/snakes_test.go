package usecases

import (
	"testing"
)

func TestGetSnakesPos(t *testing.T) {
	snake := make(map[int]int)
	snake[15] = 2
	snake[50] = 10
	snakes := NewSnakes(snake)
	pos, res := GetSnakesPos(snakes, 15)
	t.Run("Success", func(t *testing.T) {
		if !res {
			t.Errorf("Invalid snake")
		}
		if pos != 2 {
			t.Errorf("Invalid snake")
		}
	})
	pos, res = GetSnakesPos(snakes, 22)
	t.Run("Failure", func(t *testing.T) {
		if res {
			t.Errorf("Invalid snake")
		}
		if pos != 0 {
			t.Errorf("Invalid snake")
		}
	})
}
