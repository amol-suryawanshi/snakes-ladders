package usecases

import "snakes-ladders/entity"

func SetPos(p *entity.Pawn, pos int) {
	p.CurrentPos = pos
}

func GetPos(p *entity.Pawn) int {
	return p.CurrentPos
}

func NewPawn() entity.Pawn {
	return entity.Pawn{CurrentPos: 0}
}
