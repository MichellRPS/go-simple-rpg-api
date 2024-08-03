package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewEnemy(t *testing.T) {
	testCases := []Enemy{
		{"", "A", 0, 0},
		{"", "ABC", 0, 0},
		{"", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			enemy := NewEnemy(tc.Nickname)

			// Test type
			if enemy == nil {
				t.Fatalf("enemy is nil")
			}

			// Test ID value
			if _, err := uuid.Parse(enemy.ID); err != nil {
				t.Errorf("Expected enemy.ID to be a valid UUID; got %s; %s", enemy.ID, err.Error())
			}

			// Test Nickname value
			if enemy.Nickname != tc.Nickname {
				t.Errorf("Expected enemy.Nickname = %s; got %s", tc.Nickname, enemy.Nickname)
			}
			if len(enemy.Nickname) > 255 {
				t.Errorf("Expected len(enemy.Nickname) <= 255; got: %d", len(enemy.Nickname))
			}

			// Test Life value
			if enemy.Life < 1 || enemy.Life > 10 {
				t.Errorf("Expected enemy.Life to be in range [1, 10]; got %d", enemy.Life)
			}

			// Test Attack value
			if enemy.Attack < 1 || enemy.Attack > 10 {
				t.Errorf("Expected enemy.Attack to be in range [1, 10]; got %d", enemy.Attack)
			}
		})
	}
}
