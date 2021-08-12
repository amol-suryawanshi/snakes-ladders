package usecases

import (
	"math/rand"
	"snakes-ladders/entity"
	"time"
)

func RollDice(dice entity.Dice) int {
	rand.Seed(time.Now().UnixNano())
	num := randomInt(1, 6)
	if dice.Type == entity.CrookedDice {
		return (num / 2) * 2
	}
	return num
}

func randomInt(min, max int) int {
	return min + rand.Intn(int(max-min))
}

func NewDice(dt int) entity.Dice {
	return entity.Dice{Type: dt}
}
