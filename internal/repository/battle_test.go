package repository

import (
	"database/sql"
	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
)

func TestBattleRepository(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	testCases := []BattleRepository{
		{db},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			battleRepository := NewBattleRepository(tc.db)

			if battleRepository == nil {
				t.Fatalf("battleRepository is nil")
			}

			enemyId := "9d27b44f-808c-4d6e-be48-3c834d27b824"
			playerId := "97b0c0fe-a311-44fd-ad87-0bcb98057843"
			battle := entity.NewBattle(enemyId, playerId)

			_, err := battleRepository.AddBattle(battle)

			if err != nil {
				t.Errorf("battleRepository.AddBattle returned an error; %s", err.Error())
			}

			_, err = battleRepository.LoadBattles()

			if err != nil {
				t.Errorf("battleRepository.LoadBattles returned an error; %s", err.Error())
			}

			enemy, err := battleRepository.LoadEnemyById(enemyId)

			if err != nil {
				t.Errorf("battleRepository.LoadEnemyById returned an error; %s", err.Error())
			}

			err = battleRepository.SaveEnemyLife(enemy)

			if err != nil {
				t.Errorf("battleRepository.SaveEnemyLife returned an error; %s", err.Error())
			}

			player, err := battleRepository.LoadPlayerById(playerId)

			if err != nil {
				t.Errorf("battleRepository.LoadPlayerById returned an error; %s", err.Error())
			}

			err = battleRepository.SavePlayerLife(player)

			if err != nil {
				t.Errorf("battleRepository.SavePlayerLife returned an error; %s", err.Error())
			}
		})
	}
}
