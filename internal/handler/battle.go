package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	"github.com/MichellRPS/go-simple-rpg-api/internal/service"
)

type BattleHandler struct {
	BattleService *service.BattleService
}

func NewBattleHandler(BattleService *service.BattleService) *BattleHandler {
	return &BattleHandler{BattleService: BattleService}
}

func (bh *BattleHandler) AddBattle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var battle entity.Battle
	if err := json.NewDecoder(r.Body).Decode(&battle); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Message: "internal server error"})
		return
	}

	result, err := bh.BattleService.AddBattle(battle.EnemyID, battle.PlayerID)
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

func (bh *BattleHandler) LoadBattles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	battles, err := bh.BattleService.LoadBattles()
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
	json.NewEncoder(w).Encode(battles)
}
