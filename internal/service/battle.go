package service

import (
	"errors"
	"fmt"

	"github.com/Uemerson/go-simple-rpg-api/internal/entity"
	repository "github.com/Uemerson/go-simple-rpg-api/internal/repositoy"
)

type BattleService struct {
	BattleRepository repository.BattleRepository
}

func NewBattleService(BattleRepository repository.BattleRepository) *BattleService {
	return &BattleService{BattleRepository: BattleRepository}
}

func (bs *BattleService) AddBattle(enemyId, playerId string) (*entity.Battle, error) {
	if enemyId == "" || playerId == "" {
		return nil, errors.New("enemy id and player id are required")
	}

	// check if enemy exists
	enemy, err := bs.BattleRepository.LoadEnemyById(enemyId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy not found")
	}

	// check if player exists
	player, err := bs.BattleRepository.LoadPlayerById(playerId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player not found")
	}

	// check if player life and enemy life is lesser than or equal to 0
	if player.Life <= 0 || enemy.Life <= 0 {
		return nil, errors.New("enemy life and player life must be greater than 0")
	}

	battle := entity.NewBattle(enemyId, playerId)

	// check if enemy won the battle
	if battle.DiceThrown >= 1 && battle.DiceThrown <= 3 {
		// subtract player life by enemy attack
		player.Life -= enemy.Attack
		err := bs.BattleRepository.SavePlayerLife(player)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
	}

	// check if player won the battle
	if battle.DiceThrown >= 4 && battle.DiceThrown <= 6 {
		// subtract enemy life by player attack
		enemy.Life -= player.Attack
		err := bs.BattleRepository.SaveEnemyLife(enemy)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return battle, nil
}

func (bs *BattleService) LoadBattles() ([]*entity.Battle, error) {
	battles, err := bs.BattleRepository.LoadBattles()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if battles == nil {
		return []*entity.Battle{}, nil
	}
	return battles, nil
}
