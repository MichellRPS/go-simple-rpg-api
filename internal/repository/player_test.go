package repository

import (
	"database/sql"
	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
)

func TestPlayerRepository(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	testCases := []PlayerRepository{
		{db},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			playerRepository := NewPlayerRepository(tc.db)

			if playerRepository == nil {
				t.Fatalf("playerRepository is nil")
			}

			p := entity.NewPlayer("TestPlayer", 10, 10)

			_, e := playerRepository.AddPlayer(p)

			if e != nil {
				t.Errorf("playerRepository.AddPlayer returned an error; %s", e.Error())
			}

			_, e = playerRepository.LoadPlayers()

			if e != nil {
				t.Errorf("playerRepository.LoadPlayers returned an error; %s", e.Error())
			}

			_, e = playerRepository.LoadPlayerById(p.ID)

			if e != nil {
				t.Errorf("playerRepository.LoadPlayerById returned an error; %s", e.Error())
			}

			_, e = playerRepository.LoadPlayerByNickname(p.Nickname)

			if e != nil {
				t.Errorf("playerRepository.LoadPlayerByNickname returned an error; %s", e.Error())
			}

			e = playerRepository.SavePlayer(p.ID, p)

			if e != nil {
				t.Errorf("playerRepository.SavePlayer returned an error; %s", e.Error())
			}

			e = playerRepository.DeletePlayerById(p.ID)

			if e != nil {
				t.Errorf("playerRepository.DeletePlayerById returned an error; %s", e.Error())
			}
		})
	}
}
