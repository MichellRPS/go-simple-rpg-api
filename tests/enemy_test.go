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

func TestEnemyHandler(t *testing.T) {
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	enemyRepository := repository.NewEnemyRepository(db)
	enemyService := service.NewEnemyService(*enemyRepository)
	enemyHandler := handler.NewEnemyHandler(enemyService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /enemy", enemyHandler.AddEnemy)
	mux.HandleFunc("GET /enemy", enemyHandler.LoadEnemies)
	mux.HandleFunc("DELETE /enemy/{id}", enemyHandler.DeleteEnemy)
	mux.HandleFunc("GET /enemy/{id}", enemyHandler.LoadEnemy)
	mux.HandleFunc("PUT /enemy/{id}", enemyHandler.SaveEnemy)

	// Test add enemy

	jsonBody := []byte(`{"nickname": "TestEnemy"}`)
	bodyReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest("POST", "/enemy", bodyReader)

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
			"enemyHandler.AddEnemy returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var enemy entity.Enemy

	if err = json.NewDecoder(responseRecorder.Body).Decode(&enemy); err != nil {
		t.Fatalf(
			"json.Decoder.Decode returned an error: %v",
			err,
		)
	}

	// Test load enemy

	request, err = http.NewRequest("GET", "/enemy/"+enemy.ID, nil)

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
			"enemyHandler.LoadEnemy returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test load enemies

	request, err = http.NewRequest("GET", "/enemy", nil)

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
			"enemyHandler.LoadEnemies returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test save enemy

	jsonBody = []byte(`{"nickname": "TestEnemy2"}`)
	bodyReader = bytes.NewReader(jsonBody)
	request, err = http.NewRequest(http.MethodPut, "/enemy/"+enemy.ID, bodyReader)

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
			"enemyHandler.SaveEnemy returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	// Test delete enemy

	request, err = http.NewRequest("DELETE", "/enemy/"+enemy.ID, nil)

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
			"enemyHandler.DeleteEnemy returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}
}
