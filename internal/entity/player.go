package entity

import "github.com/google/uuid"

type Player struct {
	ID               string
	Nickname         string
	Life             int
	WeaponID         string
	WeaponDurability int
}

func NewPlayer(nickname string, life int, weaponId string, weaponDurability int) *Player {
	return &Player{
		ID:               uuid.New().String(),
		Nickname:         nickname,
		Life:             life,
		WeaponID:         weaponId,
		WeaponDurability: weaponDurability,
	}
}
