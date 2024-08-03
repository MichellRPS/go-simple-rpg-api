package repository

import (
	"database/sql"
	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
)

func TestEnemyRepository(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	testCases := []EnemyRepository{
		{db},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			enemyRepository := NewEnemyRepository(tc.db)

			if enemyRepository == nil {
				t.Fatalf("enemyRepository is nil")
			}

			enemy := entity.NewEnemy("TestEnemy")

			_, err = enemyRepository.AddEnemy(enemy)

			if err != nil {
				t.Errorf("enemyRepository.AddEnemy returned an error; %s", err.Error())
			}

			_, err := enemyRepository.LoadEnemyByNickname(enemy.Nickname)

			if err != nil {
				t.Errorf("enemyRepository.LoadEnemyByNickname returned an error; %s", err.Error())
			}

			_, err = enemyRepository.LoadEnemies()

			if err != nil {
				t.Errorf("enemyRepository.LoadEnemies returned an error; %s", err.Error())
			}

			_, err = enemyRepository.LoadEnemyById(enemy.ID)

			if err != nil {
				t.Errorf("enemyRepository.LoadEnemyById returned an error; %s", err.Error())
			}

			err = enemyRepository.SaveEnemy(enemy.ID, enemy)

			if err != nil {
				t.Errorf("enemyRepository.SaveEnemy returned an error; %s", err.Error())
			}

			err = enemyRepository.DeleteEnemyById(enemy.ID)

			if err != nil {
				t.Errorf("enemyRepository.DeleteEnemyById returned an error; %s", err.Error())
			}
		})
	}
}
