package entity

import (
	"github.com/google/uuid"
	"math/rand"
)

type Battle struct {
	ID         string
	EnemyID    string
	PlayerID   string
	DiceThrown int
}

func NewBattle(enemyId, playerId string) *Battle {
	return &Battle{
		ID:         uuid.New().String(),
		EnemyID:    enemyId,
		PlayerID:   playerId,
		DiceThrown: rand.Intn(6) + 1,
	}
}
