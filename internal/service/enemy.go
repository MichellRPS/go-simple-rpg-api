package service

import (
	"errors"
	"fmt"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	repository "github.com/MichellRPS/go-simple-rpg-api/internal/repository"
)

type EnemyService struct {
	EnemyRepository repository.EnemyRepository
}

func NewEnemyService(EnemyRepository repository.EnemyRepository) *EnemyService {
	return &EnemyService{EnemyRepository: EnemyRepository}
}

func (es *EnemyService) AddEnemy(nickname string) (*entity.Enemy, error) {
	if nickname == "" {
		return nil, errors.New("enemy nickname is required")
	}

	if len(nickname) > 255 {
		return nil, errors.New("enemy nickname cannot exceed 255 characters")
	}

	enemy, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy != nil {
		return nil, errors.New("enemy nickname already exits")
	}

	// select random weapon
	weapon, err := es.EnemyRepository.LoadEnemyWeapon()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if weapon == nil {
		return nil, errors.New("no weapon found")
	}

	enemy = entity.NewEnemy(nickname, weapon.ID, weapon.Durability)
	if _, err := es.EnemyRepository.AddEnemy(enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}

func (es *EnemyService) LoadEnemies() ([]*entity.Enemy, error) {
	enemies, err := es.EnemyRepository.LoadEnemies()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if enemies == nil {
		return []*entity.Enemy{}, nil
	}
	return enemies, nil
}

func (es *EnemyService) DeleteEnemy(id string) error {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if enemy == nil {
		return errors.New("enemy id not found")
	}
	if err := es.EnemyRepository.DeleteEnemyById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (es *EnemyService) LoadEnemy(id string) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}
	return enemy, nil
}

func (es *EnemyService) SaveEnemy(id, nickname string) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}

	if nickname != "" && nickname != enemy.Nickname {
		hasNickname, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasNickname != nil {
			return nil, errors.New("enemy nickname already exits")
		}
		if len(nickname) > 255 {
			return nil, errors.New("enemy nickname cannot exceed 255 characters")
		}
		enemy.Nickname = nickname
	}

	if err := es.EnemyRepository.SaveEnemy(id, enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}

func (es *EnemyService) RepairEnemyWeapon(id string) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}

	weapon, err := es.EnemyRepository.LoadWeaponById(enemy.WeaponID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if weapon == nil {
		return nil, errors.New("weapon not found")
	}

	if err := es.EnemyRepository.SaveEnemyWeaponDurability(id, weapon.Durability); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	return enemy, nil
}
