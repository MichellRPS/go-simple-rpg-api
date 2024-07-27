package entity

import (
	"github.com/google/uuid"
	"math/rand"
)

type Enemy struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
}

func NewEnemy(nickname string) *Enemy {
	return &Enemy{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     rand.Intn(10) + 1,
		Attack:   rand.Intn(10) + 1,
	}
}
