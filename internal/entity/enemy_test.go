package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewEnemy(t *testing.T) {
	testCases := []Enemy{
		{"", "A", 0, "1ccaac4a-3be5-442f-bba0-e7d78dfdc4a3"},
		{"", "ABC", 0, "177b7b6f-ba70-4dc9-aae6-a92c07a02a2e"},
		{"", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, "0a96a984-2f14-43d6-9563-1ec193f6e932"},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			enemy := NewEnemy(tc.Nickname, tc.WeaponID)

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

			// Test WeaponID value
			if _, err := uuid.Parse(enemy.WeaponID); err != nil {
				t.Errorf("Expected enemy.WeaponID to be a valid UUID; got %s; %s", enemy.WeaponID, err.Error())
			}
		})
	}
}
