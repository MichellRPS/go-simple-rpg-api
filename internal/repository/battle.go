package repository

import (
	"database/sql"
	"errors"
	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
)

type BattleRepository struct {
	db *sql.DB
}

func NewBattleRepository(db *sql.DB) *BattleRepository {
	return &BattleRepository{db: db}
}

func (br *BattleRepository) LoadEnemyById(enemyId string) (*entity.Enemy, error) {
	var enemy entity.Enemy
	err := br.db.QueryRow("SELECT id, nickname, life, attack FROM enemy WHERE id = $1", enemyId).Scan(&enemy.ID, &enemy.Nickname, &enemy.Life, &enemy.Attack)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &enemy, nil
}

func (br *BattleRepository) LoadPlayerById(playerId string) (*entity.Player, error) {
	var player entity.Player
	err := br.db.QueryRow("SELECT id, nickname, life, attack FROM player WHERE id = $1", playerId).Scan(&player.ID, &player.Nickname, &player.Life, &player.Attack)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (br *BattleRepository) SaveEnemyLife(enemy *entity.Enemy) error {
	_, err := br.db.Exec("UPDATE enemy SET life = $1 WHERE id = $2", enemy.Life, enemy.ID)
	if err != nil {
		return err
	}
	return nil
}

func (br *BattleRepository) SavePlayerLife(player *entity.Player) error {
	_, err := br.db.Exec("UPDATE player SET life = $1 WHERE id = $2", player.Life, player.ID)
	if err != nil {
		return err
	}
	return nil
}

func (br *BattleRepository) AddBattle(battle *entity.Battle) (string, error) {
	_, err := br.db.Exec("INSERT INTO battle (id, enemyid, playerid, dicethrown) VALUES ($1, $2, $3, $4)", battle.ID, battle.EnemyID, battle.PlayerID, battle.DiceThrown)
	if err != nil {
		return "", err
	}
	return battle.ID, nil
}

func (br *BattleRepository) LoadBattles() ([]*entity.Battle, error) {
	rows, err := br.db.Query("SELECT id, enemyid, playerid, dicethrown FROM battle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var battles []*entity.Battle
	for rows.Next() {
		var battle entity.Battle
		if err := rows.Scan(&battle.ID, &battle.EnemyID, &battle.PlayerID, &battle.DiceThrown); err != nil {
			return nil, err
		}
		battles = append(battles, &battle)
	}
	return battles, nil
}
