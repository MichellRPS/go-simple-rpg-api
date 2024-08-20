package entity

import "github.com/google/uuid"

type Player struct {
	ID       string
	Nickname string
	Life     int
	WeaponID string
}

func NewPlayer(nickname string, life int, weaponId string) *Player {
	return &Player{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		WeaponID: weaponId,
	}
}
