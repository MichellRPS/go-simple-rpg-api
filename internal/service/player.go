package service

import (
	"errors"
	"fmt"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	repository "github.com/MichellRPS/go-simple-rpg-api/internal/repository"
)

type PlayerService struct {
	PlayerRepository repository.PlayerRepository
}

func NewPlayerService(PlayerRepository repository.PlayerRepository) *PlayerService {
	return &PlayerService{PlayerRepository: PlayerRepository}
}

func (ps *PlayerService) AddPlayer(nickname string, life int, weaponId string) (*entity.Player, error) {
	if nickname == "" || life == 0 || weaponId == "" {
		return nil, errors.New("player nickname, life and weapon id is required")
	}

	if len(nickname) > 255 {
		return nil, errors.New("player nickname cannot exceed 255 characters")
	}

	// check if weapon exists
	weapon, err := ps.PlayerRepository.LoadWeaponById(weaponId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if weapon == nil {
		return nil, errors.New("weapon not found")
	}

	if life > 100 || life <= 0 {
		return nil, errors.New("player life must be between 1 and 100")
	}

	player, err := ps.PlayerRepository.LoadPlayerByNickname(nickname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player != nil {
		return nil, errors.New("player nickname already exits")
	}

	player = entity.NewPlayer(nickname, life, weapon.ID, weapon.Durability)
	if _, err := ps.PlayerRepository.AddPlayer(player); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return player, nil
}

func (ps *PlayerService) LoadPlayers() ([]*entity.Player, error) {
	players, err := ps.PlayerRepository.LoadPlayers()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if players == nil {
		return []*entity.Player{}, nil
	}
	return players, nil
}

func (ps *PlayerService) DeletePlayer(id string) error {
	player, err := ps.PlayerRepository.LoadPlayerById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if player == nil {
		return errors.New("player id not found")
	}
	if err := ps.PlayerRepository.DeletePlayerById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (ps *PlayerService) LoadPlayer(id string) (*entity.Player, error) {
	player, err := ps.PlayerRepository.LoadPlayerById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player id not found")
	}
	return player, nil
}

func (ps *PlayerService) SavePlayer(id, nickname string, life int, weaponId string) (*entity.Player, error) {
	player, err := ps.PlayerRepository.LoadPlayerById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player id not found")
	}

	if nickname != "" && nickname != player.Nickname {
		hasNickname, err := ps.PlayerRepository.LoadPlayerByNickname(nickname)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasNickname != nil {
			return nil, errors.New("player nickname already exits")
		}
		if len(nickname) > 255 {
			return nil, errors.New("player nickname cannot exceed 255 characters")
		}
		player.Nickname = nickname
	}

	if weaponId != "" && weaponId != player.WeaponID {
		// check if weapon exists
		weapon, err := ps.PlayerRepository.LoadWeaponById(weaponId)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if weapon == nil {
			return nil, errors.New("weapon not found")
		}
		player.WeaponID = weaponId
		player.WeaponDurability = weapon.Durability
	}

	if life != 0 && life != player.Life {
		if life > 100 || life <= 0 {
			return nil, errors.New("player life must be between 1 and 100")
		}
		player.Life = life
	}

	if err := ps.PlayerRepository.SavePlayer(id, player); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return player, nil
}

func (ps *PlayerService) RepairPlayerWeapon(id string) (*entity.Player, error) {
	player, err := ps.PlayerRepository.LoadPlayerById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if player == nil {
		return nil, errors.New("player id not found")
	}

	weapon, err := ps.PlayerRepository.LoadWeaponById(player.WeaponID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if weapon == nil {
		return nil, errors.New("weapon not found")
	}

	if err := ps.PlayerRepository.SavePlayerWeaponDurability(id, weapon.Durability); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	return player, nil
}
