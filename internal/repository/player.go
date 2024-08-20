package repository

import (
	"database/sql"
	"errors"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (pr *PlayerRepository) LoadPlayers() ([]*entity.Player, error) {
	rows, err := pr.db.Query("SELECT id, nickname, life, weaponid FROM player")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []*entity.Player
	for rows.Next() {
		var player entity.Player
		if err := rows.Scan(&player.ID, &player.Nickname, &player.Life, &player.WeaponID); err != nil {
			return nil, err
		}
		players = append(players, &player)
	}
	return players, nil
}

func (pr *PlayerRepository) LoadPlayerById(id string) (*entity.Player, error) {
	var player entity.Player
	err := pr.db.QueryRow("SELECT id, nickname, life, weaponid FROM player WHERE id = $1", id).Scan(&player.ID, &player.Nickname, &player.Life, &player.WeaponID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (pr *PlayerRepository) LoadPlayerByNickname(nickname string) (*entity.Player, error) {
	var player entity.Player
	err := pr.db.QueryRow("SELECT id, nickname, life, weaponid FROM player WHERE nickname LIKE $1", nickname).Scan(&player.ID, &player.Nickname, &player.Life, &player.WeaponID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (pr *PlayerRepository) AddPlayer(player *entity.Player) (string, error) {
	_, err := pr.db.Exec("INSERT INTO player (id, nickname, life, weaponid) VALUES ($1, $2, $3, $4)", player.ID, player.Nickname, player.Life, player.WeaponID)
	if err != nil {
		return "", err
	}
	return player.ID, nil
}

func (pr *PlayerRepository) DeletePlayerById(id string) error {
	_, err := pr.db.Exec("DELETE FROM player WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PlayerRepository) SavePlayer(id string, player *entity.Player) error {
	_, err := pr.db.Exec("UPDATE player SET nickname = $1, life = $2, weaponid = $3 WHERE id = $4", player.Nickname, player.Life, player.WeaponID, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PlayerRepository) LoadWeaponById(weaponId string) (*entity.Weapon, error) {
	var weapon entity.Weapon
	err := pr.db.QueryRow("SELECT id, name, attack, defense FROM weapon WHERE id = $1", weaponId).Scan(&weapon.ID, &weapon.Name, &weapon.Attack, &weapon.Defense)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &weapon, nil
}
