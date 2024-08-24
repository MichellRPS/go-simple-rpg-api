package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	"github.com/MichellRPS/go-simple-rpg-api/internal/service"
)

type WeaponHandler struct {
	WeaponService *service.WeaponService
}

func NewWeaponHandler(weaponService *service.WeaponService) *WeaponHandler {
	return &WeaponHandler{WeaponService: weaponService}
}

func (wh *WeaponHandler) AddWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var weapon entity.Weapon
	if err := json.NewDecoder(r.Body).Decode(&weapon); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := wh.WeaponService.AddWeapon(weapon.Name, weapon.Attack, weapon.Defense, weapon.Durability)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (wh *WeaponHandler) LoadWeapons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	weapons, err := wh.WeaponService.LoadWeapons()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weapons)
}

func (wh *WeaponHandler) DeleteWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if err := wh.WeaponService.DeleteWeapon(id); err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}

func (wh *WeaponHandler) LoadWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	weapon, err := wh.WeaponService.LoadWeapon(id)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weapon)
}

func (wh *WeaponHandler) SaveWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var weapon entity.Weapon
	if err := json.NewDecoder(r.Body).Decode(&weapon); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := wh.WeaponService.SaveWeapon(id, weapon.Name, weapon.Attack, weapon.Defense, weapon.Durability)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "internal server error"):
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
