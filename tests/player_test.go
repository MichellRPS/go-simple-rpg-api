package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MichellRPS/go-simple-rpg-api/internal/entity"
	"github.com/MichellRPS/go-simple-rpg-api/internal/handler"
	repository "github.com/MichellRPS/go-simple-rpg-api/internal/repository"
	"github.com/MichellRPS/go-simple-rpg-api/internal/service"
	_ "github.com/lib/pq"
)

func TestPlayerHandler(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	playerRepository := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(*playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /player", playerHandler.AddPlayer)
	mux.HandleFunc("GET /player", playerHandler.LoadPlayers)
	mux.HandleFunc("DELETE /player/{id}", playerHandler.DeletePlayer)
	mux.HandleFunc("GET /player/{id}", playerHandler.LoadPlayer)
	mux.HandleFunc("PUT /player/{id}", playerHandler.SavePlayer)
	mux.HandleFunc("GET /player/{id}/repair-weapon", playerHandler.RepairPlayerWeapon)

	// Test add player

	jsonBody := []byte(`{"nickname": "TestPlayer", "life": 100, "weaponid": "9b48bd04-d944-463b-be4e-a4eef73d7e15"}`)
	bodyReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest("POST", "/player", bodyReader)

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
			"playerHandler.AddPlayer returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var player entity.Player

	if err = json.NewDecoder(responseRecorder.Body).Decode(&player); err != nil {
		t.Fatalf(
            "json.Decoder.Decode returned an error: %v",
            err,
        )
	}

	// Test load player

	request, err = http.NewRequest("GET", "/player/"+player.ID, nil)

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
			"playerHandler.LoadPlayer returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test load players

	request, err = http.NewRequest("GET", "/player", nil)

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
			"playerHandler.LoadPlayers returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test save player

    jsonBody = []byte(`{"nickname": "TestPlayer2", "life": 50, "weaponid": "8c3910f5-7248-40c7-b0c2-38446f328ce9"}`)
	bodyReader = bytes.NewReader(jsonBody)
	request, err = http.NewRequest(http.MethodPut, "/player/"+player.ID, bodyReader)

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
			"playerHandler.SavePlayer returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test repair weapon

	request, err = http.NewRequest("GET", "/player/"+player.ID+"/repair-weapon", nil)

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
			"playerHandler.RepairPlayerWeapon returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test delete player

	request, err = http.NewRequest("DELETE", "/player/"+player.ID, nil)

	if err != nil {
		t.Fatalf(
            "http.NewRequest returned an error: %v",
            err,
        )
	}

	responseRecorder = httptest.NewRecorder()
	mux.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusNoContent {
		t.Errorf(
			"playerHandler.DeletePlayer returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}
}
