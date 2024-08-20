package repository

import (
	"database/sql"
	"errors"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
)

type WeaponRepository struct {
	db *sql.DB
}

func NewWeaponRepository(db *sql.DB) *WeaponRepository {
	return &WeaponRepository{db: db}
}

func (wr *WeaponRepository) LoadWeapons() ([]*entity.Weapon, error) {
	rows, err := wr.db.Query("SELECT id, name, attack, defense FROM weapon")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weapons []*entity.Weapon
	for rows.Next() {
		var weapon entity.Weapon
		if err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense); err != nil {
			return nil, err
		}
		weapons = append(weapons, &weapon)
	}
	return weapons, nil
}

func (wr *WeaponRepository) LoadWeaponById(id string) (*entity.Weapon, error) {
	var weapon entity.Weapon
	err := wr.db.QueryRow("SELECT id, name, attack, defense FROM weapon WHERE id = $1", id).Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &weapon, nil
}

func (wr *WeaponRepository) LoadWeaponByName(name string) (*entity.Weapon, error) {
	var weapon entity.Weapon
	err := wr.db.QueryRow("SELECT id, name, attack, defense FROM weapon WHERE name LIKE $1", name).Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &weapon, nil
}

func (wr *WeaponRepository) AddWeapon(weapon *entity.Weapon) (string, error) {
	_, err := wr.db.Exec("INSERT INTO weapon (id, name, attack, defense) VALUES ($1, $2, $3, $4)", weapon.ID, weapon.Name, weapon.Attack, weapon.Defense)
	if err != nil {
		return "", err
	}
	return weapon.ID, nil
}

func (wr *WeaponRepository) DeleteWeaponById(id string) error {
	_, err := wr.db.Exec("DELETE FROM weapon WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WeaponRepository) SaveWeapon(id string, weapon *entity.Weapon) error {
	_, err := wr.db.Exec("UPDATE weapon SET name = $1, attack = $2, defense = $3 WHERE id = $4", weapon.Name, weapon.Attack, weapon.Defense, id)
	if err != nil {
		return err
	}
	return nil
}
