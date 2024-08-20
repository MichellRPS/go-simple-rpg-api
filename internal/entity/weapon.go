package entity

import "github.com/google/uuid"

type Weapon struct {
	ID      string
	Name    string
	Attack  int
	Defense float64
}

func NewWeapon(name string, attack int, defense float64) *Weapon {
	return &Weapon{
		ID:      uuid.New().String(),
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}
