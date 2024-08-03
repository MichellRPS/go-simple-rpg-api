package service

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/MichellRPS/go-simple-rpg-api/internal/repository"
	_ "github.com/lib/pq"
)

func TestBattleService(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	battleRepository := repository.NewBattleRepository(db)

	testCases := []BattleService{
		{*battleRepository},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			battleService := NewBattleService(tc.BattleRepository)

			if battleService == nil {
				t.Fatalf("battleService is nil")
			}

			_, err := battleService.AddBattle("9d27b44f-808c-4d6e-be48-3c834d27b824", "97b0c0fe-a311-44fd-ad87-0bcb98057843")

			if err != nil {
				t.Errorf("battleService.AddBattle returned an error; %s", err.Error())
			}

			_, err = battleService.LoadBattles()

			if err != nil {
				t.Errorf("battleService.LoadBattles returned an error; %s", err.Error())
			}
		})
	}
}
