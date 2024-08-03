package service

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/MichellRPS/go-simple-rpg-api/internal/repository"
	_ "github.com/lib/pq"
)

func TestEnemyService(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	enemyRepository := repository.NewEnemyRepository(db)

	testCases := []EnemyService{
		{*enemyRepository},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			enemyService := NewEnemyService(tc.EnemyRepository)

			if enemyService == nil {
				t.Fatalf("enemyService is nil")
			}

			enemy, err := enemyService.AddEnemy("TestEnemy")

			if err != nil {
				t.Errorf("enemyService.AddEnemy returned an error; %s", err.Error())
			}

			_, err = enemyService.LoadEnemy(enemy.ID)

			if err != nil {
				t.Errorf("enemyService.LoadEnemy returned an error; %s", err.Error())
			}

			_, err = enemyService.LoadEnemies()

			if err != nil {
				t.Errorf("enemyService.LoadEnemies returned an error; %s", err.Error())
			}

			_, err = enemyService.SaveEnemy(enemy.ID, enemy.Nickname)

			if err != nil {
				t.Errorf("enemyService.SaveEnemy returned an error; %s", err.Error())
			}

			err = enemyService.DeleteEnemy(enemy.ID)

			if err != nil {
				t.Errorf("enemyService.DeleteEnemy returned an error; %s", err.Error())
			}
		})
	}
}
