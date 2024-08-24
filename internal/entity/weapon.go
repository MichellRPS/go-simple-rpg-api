package entity

import "github.com/google/uuid"

type Weapon struct {
	ID         string
	Name       string
	Attack     int
	Defense    float64
	Durability int
}

func NewWeapon(name string, attack int, defense float64, durability int) *Weapon {
	return &Weapon{
		ID:         uuid.New().String(),
		Name:       name,
		Attack:     attack,
		Defense:    defense,
		Durability: durability,
	}
}
