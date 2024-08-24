package repository

import (
	"database/sql"
	"errors"
	"math/rand"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
)

type EnemyRepository struct {
	db *sql.DB
}

func NewEnemyRepository(db *sql.DB) *EnemyRepository {
	return &EnemyRepository{db: db}
}

func (er *EnemyRepository) LoadEnemyByNickname(nickname string) (*entity.Enemy, error) {
	var enemy entity.Enemy
	err := er.db.QueryRow("SELECT id, nickname, life, weaponid, weapondurability FROM enemy WHERE nickname LIKE $1", nickname).Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.WeaponID, &enemy.WeaponDurability)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &enemy, nil
}

func (er *EnemyRepository) AddEnemy(enemy *entity.Enemy) (string, error) {
	_, err := er.db.Exec("INSERT INTO enemy (id, nickname, life, weaponid, weapondurability) VALUES ($1, $2, $3, $4, $5)", enemy.ID, enemy.Nickname, enemy.Life, enemy.WeaponID, enemy.WeaponDurability)
	if err != nil {
		return "", err
	}
	return enemy.ID, nil
}

func (er *EnemyRepository) LoadEnemies() ([]*entity.Enemy, error) {
	rows, err := er.db.Query("SELECT id, nickname, life, weaponid, weapondurability FROM enemy")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enemies []*entity.Enemy
	for rows.Next() {
		var enemy entity.Enemy
		if err := rows.Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.WeaponID, &enemy.WeaponDurability); err != nil {
			return nil, err
		}
		enemies = append(enemies, &enemy)
	}
	return enemies, nil
}

func (er *EnemyRepository) LoadEnemyById(id string) (*entity.Enemy, error) {
	var enemy entity.Enemy
	err := er.db.QueryRow("SELECT id, nickname, life, weaponid, weapondurability FROM enemy WHERE id = $1", id).Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.WeaponID, &enemy.WeaponDurability)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &enemy, nil
}

func (er *EnemyRepository) DeleteEnemyById(id string) error {
	_, err := er.db.Exec("DELETE FROM enemy WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (er *EnemyRepository) SaveEnemy(id string, enemy *entity.Enemy) error {
	_, err := er.db.Exec("UPDATE enemy SET nickname = $1 WHERE id = $2", enemy.Nickname, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *EnemyRepository) SaveEnemyWeaponDurability(id string, weaponDurability int) error {
	_, err := er.db.Exec("UPDATE enemy SET weapondurability = $1 WHERE id = $2", weaponDurability, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *EnemyRepository) LoadWeapons() ([]*entity.Weapon, error) {
	rows, err := er.db.Query("SELECT id, name, attack, defense, durability FROM weapon")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weapons []*entity.Weapon
	for rows.Next() {
		var weapon entity.Weapon
		if err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense, &weapon.Durability); err != nil {
			return nil, err
		}
		weapons = append(weapons, &weapon)
	}
	return weapons, nil
}

func (er *EnemyRepository) LoadWeaponById(weaponId string) (*entity.Weapon, error) {
	var weapon entity.Weapon
	err := er.db.QueryRow("SELECT id, name, attack, defense, durability FROM weapon WHERE id = $1", weaponId).Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense, &weapon.Durability)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &weapon, nil
}

func (er *EnemyRepository) LoadEnemyWeapon() (*entity.Weapon, error) {
	weapons, err := er.LoadWeapons()
	if err != nil {
		return nil, errors.New("internal server error")
	}
	if len(weapons) == 0 {
		return nil, errors.New("no weapon found")
	}
	return weapons[rand.Intn(len(weapons))], nil
}
