package tests

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MichellRPS/go-simple-rpg-api/internal/handler"
	repository "github.com/MichellRPS/go-simple-rpg-api/internal/repository"
	"github.com/MichellRPS/go-simple-rpg-api/internal/service"
	_ "github.com/lib/pq"
)

func TestBattleHandler(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	battleRepository := repository.NewBattleRepository(db)
	battleService := service.NewBattleService(*battleRepository)
	battleHandler := handler.NewBattleHandler(battleService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /battle", battleHandler.AddBattle)
	mux.HandleFunc("GET /battle", battleHandler.LoadBattles)

	// Test add battle

	jsonBody := []byte(`{"enemyid": "94e48413-dddf-44e2-a4f3-8f6f557d1d8f", "playerid": "e963ebd9-e856-4c9d-8b94-89c3f1b36109"}`)
	bodyReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest("POST", "/battle", bodyReader)

	if err != nil {
		t.Fatalf(
            "http.NewRequest returned an error: %v",
            err,
        )
	}

	responseRecorder := httptest.NewRecorder()
	mux.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf(
			"battleHandler.AddBattle returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test load battles

	request, err = http.NewRequest("GET", "/battle", nil)

	if err != nil {
		t.Fatalf(
            "http.NewRequest returned an error: %v",
            err,
        )
	}

	responseRecorder = httptest.NewRecorder()
	mux.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf(
			"battleHandler.LoadBattles returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}
}
