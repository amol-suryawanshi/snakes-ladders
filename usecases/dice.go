package usecases

import (
	"math/rand"
	"snakes-ladders/entity"
	"time"
)

func RollDice(dice entity.Dice) int {
	if dice.Type == entity.NormalDice || dice.Type == entity.CrookedDice {
		rand.Seed(time.Now().UnixNano())
		num := randomInt(1, 6)
		if dice.Type == entity.CrookedDice {
			return ((num + 1) / 2) * 2
		}
		return num
	}
	return -1
}

func randomInt(min, max int) int {
	return min + rand.Intn(int(max-min))
}

func NewDice(dt int) entity.Dice {
	if dt == entity.NormalDice || dt == entity.CrookedDice {
		return entity.Dice{Type: dt}
	}
	return entity.Dice{}
}
