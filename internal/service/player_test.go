package service

import (
	"database/sql"
	"github.com/MichellRPS/go-simple-rpg-api/internal/repository"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
)

func TestPlayerService(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	playerRepository := repository.NewPlayerRepository(db)

	testCases := []PlayerService{
		{*playerRepository},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			playerService := NewPlayerService(tc.PlayerRepository)

			if playerService == nil {
				t.Fatalf("playerService is nil")
			}

			player, err := playerService.AddPlayer("TestPlayer", 10, 10)

			if err != nil {
				t.Errorf("playerService.AddPlayer returned an error; %s", err.Error())
			}

			_, err = playerService.LoadPlayer(player.ID)

			if err != nil {
				t.Errorf("playerService.LoadPlayer returned an error; %s", err.Error())
			}

			_, err = playerService.LoadPlayers()

			if err != nil {
				t.Errorf("playerService.LoadPlayers returned an error; %s", err.Error())
			}

			_, err = playerService.SavePlayer(player.ID, player.Nickname, player.Life, player.Attack)

			if err != nil {
				t.Errorf("playerService.SavePlayer returned an error; %s", err.Error())
			}

			err = playerService.DeletePlayer(player.ID)

			if err != nil {
				t.Errorf("playerService.DeletePlayer returned an error; %s", err.Error())
			}
		})
	}
}
