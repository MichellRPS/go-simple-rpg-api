package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

type PlayerRequest struct {
	Nickname string
	Life     int
	Attack   int
}

type PlayerResponse struct {
	Message string `json:"message"`
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var playerRequest PlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&playerRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Internal Server Error"})
		return
	}

	if playerRequest.Nickname == "" || playerRequest.Life == 0 || playerRequest.Attack == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname, life and attack is required"})
		return
	}

	if playerRequest.Attack > 10 || playerRequest.Attack <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player attack must be between 1 and 10"})
		return
	}

	if playerRequest.Life > 100 || playerRequest.Life <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player life must be between 1 and 100"})
		return
	}

	for _, player := range players {
		if player.Nickname == playerRequest.Nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname already exits"})
			return
		}
	}

	player := PlayerRequest{
		Nickname: playerRequest.Nickname,
		Life:     playerRequest.Life,
		Attack:   playerRequest.Attack}
	players = append(players, player)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func LoadPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := r.PathValue("nickname")

	for i, player := range players {
		if player.Nickname == nickname {
			players = append(players[:i], players[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(PlayerResponse{
		Message: "Player nickname not found",
	})
}

func LoadPlayerByNickname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := r.PathValue("nickname")

	for _, player := range players {
		if player.Nickname == nickname {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(PlayerResponse{
		Message: "Player nickname not found",
	})
}

func SavePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nickname := r.PathValue("nickname")

	var playerRequest PlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&playerRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Internal Server Error"})
		return
	}

	if playerRequest.Nickname == "" {
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname is required"})
		return
	}

	indexPlayer := -1
	for i, player := range players {
		if player.Nickname == nickname {
			indexPlayer = i
		}
		if player.Nickname == playerRequest.Nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname already exits"})
			return
		}
	}

	if indexPlayer != -1 {
		players[indexPlayer].Nickname = playerRequest.Nickname
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(players[indexPlayer])
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(PlayerResponse{
		Message: "Player nickname not found",
	})
}

var players []PlayerRequest

type EnemyRequest struct {
	Nickname string
	Life     int
	Attack   int
}

type EnemyResponse struct {
	Message string `json:"message"`
}

func AddEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var enemyRequest EnemyRequest
	if err := json.NewDecoder(r.Body).Decode(&enemyRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Internal Server Error"})
		return
	}

	if enemyRequest.Nickname == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname is required"})
		return
	}

	enemyRequest.Life = rand.Intn(10)
	enemyRequest.Attack = rand.Intn(10)

	if enemyRequest.Life == 0 {
		enemyRequest.Life += 1
	}

	if enemyRequest.Attack == 0 {
		enemyRequest.Attack += 1
	}

	for _, enemy := range enemies {
		if enemy.Nickname == enemyRequest.Nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname already exits"})
			return
		}
	}

	enemy := EnemyRequest{
		Nickname: enemyRequest.Nickname,
		Life:     enemyRequest.Life,
		Attack:   enemyRequest.Attack}
	enemies = append(enemies, enemy)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}

func LoadEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemies)
}

func DeleteEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := r.PathValue("nickname")

	for i, enemy := range enemies {
		if enemy.Nickname == nickname {
			enemies = append(enemies[:i], enemies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(EnemyResponse{
		Message: "Enemy nickname not found",
	})
}

func LoadEnemyByNickname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := r.PathValue("nickname")

	for _, enemy := range enemies {
		if enemy.Nickname == nickname {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(enemy)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(EnemyResponse{
		Message: "Enemy nickname not found",
	})
}

func SaveEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nickname := r.PathValue("nickname")

	var enemyRequest EnemyRequest
	if err := json.NewDecoder(r.Body).Decode(&enemyRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Internal Server Error"})
		return
	}

	if enemyRequest.Nickname == "" {
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname is required"})
		return
	}

	indexEnemy := -1
	for i, enemy := range enemies {
		if enemy.Nickname == nickname {
			indexEnemy = i
		}
		if enemy.Nickname == enemyRequest.Nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname already exits"})
			return
		}
	}

	if indexEnemy != -1 {
		enemies[indexEnemy].Nickname = enemyRequest.Nickname
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(enemies[indexEnemy])
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(EnemyResponse{
		Message: "Enemy nickname not found",
	})
}

var enemies []EnemyRequest

type BattleRequest struct {
	ID         string
	Enemy      string
	Player     string
	DiceThrown int
}

type BattleResponse struct {
	Message string `json:"message"`
}

func AddBattle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var battleRequest BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&battleRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(BattleResponse{Message: "Internal Server Error"})
		return
	}

	if battleRequest.Player == "" || battleRequest.Enemy == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BattleResponse{Message: "Player and enemy nicknames are required"})
		return
	}

	// check if player exists
	var playerRequest PlayerRequest
	playerIndex := -1

	for index, player := range players {
		if player.Nickname == battleRequest.Player {
			playerRequest = player
			playerIndex = index
		}
	}

	if playerRequest.Nickname == "" || playerIndex == -1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BattleResponse{Message: "Player not found"})
		return
	}

	// check if enemy exists
	var enemyRequest EnemyRequest
	enemyIndex := -1

	for index, enemy := range enemies {
		if enemy.Nickname == battleRequest.Enemy {
			enemyRequest = enemy
			enemyIndex = index
		}
	}

	if enemyRequest.Nickname == "" || enemyIndex == -1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BattleResponse{Message: "Enemy not found"})
		return
	}

	// check if player and enemy life is lesser than or equal to 0
	if playerRequest.Life <= 0 || enemyRequest.Life <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BattleResponse{Message: "Player and enemy life must be greater than 0"})
		return
	}

	// generate dicethrown
	battleRequest.DiceThrown = rand.Intn(6)

	if battleRequest.DiceThrown == 0 {
		battleRequest.DiceThrown += 1
	}

	if battleRequest.DiceThrown == 1 || battleRequest.DiceThrown == 2 || battleRequest.DiceThrown == 3 {
		// enemy wins and attacks player
		playerRequest.Life -= enemyRequest.Attack
		// save player's life
		players[playerIndex].Life = playerRequest.Life
	} else {
		// player wins and attacks enemy
		enemyRequest.Life -= playerRequest.Attack
		// save enemy's life
		enemies[enemyIndex].Life = enemyRequest.Life
	}

	battleRequest.ID = uuid.NewString()

	battle := BattleRequest{
		ID:         battleRequest.ID,
		Enemy:      battleRequest.Enemy,
		Player:     battleRequest.Player,
		DiceThrown: battleRequest.DiceThrown}
	battles = append(battles, battle)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(battle)
}

func LoadBattles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(battles)
}

var battles []BattleRequest

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /player", AddPlayer)
	mux.HandleFunc("GET /player", LoadPlayers)
	mux.HandleFunc("DELETE /player/{nickname}", DeletePlayer)
	mux.HandleFunc("GET /player/{nickname}", LoadPlayerByNickname)
	mux.HandleFunc("PUT /player/{nickname}", SavePlayer)

	mux.HandleFunc("POST /enemy", AddEnemy)
	mux.HandleFunc("GET /enemy", LoadEnemies)
	mux.HandleFunc("DELETE /enemy/{nickname}", DeleteEnemy)
	mux.HandleFunc("GET /enemy/{nickname}", LoadEnemyByNickname)
	mux.HandleFunc("PUT /enemy/{nickname}", SaveEnemy)

	mux.HandleFunc("POST /battle", AddBattle)
	mux.HandleFunc("GET /battle", LoadBattles)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}
