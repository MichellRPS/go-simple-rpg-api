package service

import (
	"errors"
	"fmt"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	repository "github.com/MichellRPS/go-simple-rpg-api/internal/repository"
)

type WeaponService struct {
	WeaponRepository repository.WeaponRepository
}

func NewWeaponService(WeaponRepository repository.WeaponRepository) *WeaponService {
	return &WeaponService{WeaponRepository: WeaponRepository}
}

func (ws *WeaponService) AddWeapon(name string, attack int, defense float64, durability int) (*entity.Weapon, error) {
	if name == "" || attack == 0 || defense == 0 || durability == 0 {
		return nil, errors.New("weapon name, attack, defense and durability is required")
	}

	if len(name) > 255 {
		return nil, errors.New("weapon name cannot exceed 255 characters")
	}

	if defense > 1.0 || defense <= 0.0 {
		return nil, errors.New("weapon defense must be between 0.1 and 1.0")
	}

	if attack > 100 || attack <= 0 {
		return nil, errors.New("weapon attack must be between 1 and 100")
	}

	if durability > 10 || durability <= 0 {
		return nil, errors.New("weapon durability must be between 1 and 10")
	}

	weapon, err := ws.WeaponRepository.LoadWeaponByName(name)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if weapon != nil {
		return nil, errors.New("weapon name already exits")
	}

	weapon = entity.NewWeapon(name, attack, defense, durability)
	if _, err := ws.WeaponRepository.AddWeapon(weapon); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return weapon, nil
}

func (ws *WeaponService) LoadWeapons() ([]*entity.Weapon, error) {
	weapons, err := ws.WeaponRepository.LoadWeapons()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if weapons == nil {
		return []*entity.Weapon{}, nil
	}
	return weapons, nil
}

func (ws *WeaponService) DeleteWeapon(id string) error {
	weapon, err := ws.WeaponRepository.LoadWeaponById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if weapon == nil {
		return errors.New("weapon id not found")
	}
	if err := ws.WeaponRepository.DeleteWeaponById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (ws *WeaponService) LoadWeapon(id string) (*entity.Weapon, error) {
	weapon, err := ws.WeaponRepository.LoadWeaponById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if weapon == nil {
		return nil, errors.New("weapon id not found")
	}
	return weapon, nil
}

func (ws *WeaponService) SaveWeapon(id, name string, attack int, defense float64, durability int) (*entity.Weapon, error) {
	weapon, err := ws.WeaponRepository.LoadWeaponById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if weapon == nil {
		return nil, errors.New("weapon id not found")
	}

	if name != "" && name != weapon.Name {
		hasName, err := ws.WeaponRepository.LoadWeaponByName(name)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasName != nil {
			return nil, errors.New("weapon name already exits")
		}
		if len(name) > 255 {
			return nil, errors.New("weapon name cannot exceed 255 characters")
		}
		weapon.Name = name
	}

	if defense != 0.0 && defense != weapon.Defense {
		if defense > 1.0 || defense <= 0.0 {
			return nil, errors.New("weapon defense must be between 0.1 and 1.0")
		}
		weapon.Defense = defense
	}

	if attack != 0 && attack != weapon.Attack {
		if attack > 100 || attack <= 0 {
			return nil, errors.New("weapon attack must be between 1 and 100")
		}
		weapon.Attack = attack
	}

	if durability != 0 && durability != weapon.Durability {
		if durability > 10 || durability <= 0 {
			return nil, errors.New("weapon durability must be between 1 and 10")
		}
		weapon.Durability = durability
	}

	if err := ws.WeaponRepository.SaveWeapon(id, weapon); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return weapon, nil
}
